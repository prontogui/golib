// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pgcomm

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"

	pb "github.com/prontogui/golib/pb"
	"google.golang.org/grpc"
)

// The complete package representing and update go to the app, or coming from the app.
type UpdatePackage struct {
	// The ID of the streaming session that this update applies to.
	sessionId int8

	// The update encoded in bytes
	cbor []byte
}

// Implementation of the PGServer
type PGComm struct {
	pb.UnimplementedPGServiceServer

	// The active server.
	// - a valid reference after calling Serve() and returns no error.
	// - null reference after calling StopServing().
	active_server *grpc.Server

	// TODO:  make a struct holding the session # and []byte (payload).  Build channels of this struct instead.
	inboundUpdates  chan UpdatePackage
	outboundUpdates chan UpdatePackage

	// Session ID currently used for streaming updates back and forth.  Updated inside StreamUpdate call.
	streamingSessionId int8

	// Session ID used within an ExchangeUpdate call.
	exchangeSessionId int8

	// Channel to signify a new session ID was created, due to previous session was ended.
	newSessionId chan int8

	// Channel to prevent simultaneous calls the StreamUpdates
	busyStreaming chan bool
}

func NewPGComm() *PGComm {
	pgc := &PGComm{}
	return pgc
}

// Called as a Go routine by StreamUpdates.
//
// This function pulls updates from a channel and streams them to app.  It exits upon receiving
// an value from the cancel channel provided to this function.
func (pgc *PGComm) streamOutboundUpdates(sessionId int8, cancel chan bool, stream pb.PGService_StreamUpdatesServer) {

	// Loop for every udpate streamed back to caller...
	for {
		select {
		// If API call StreamUpdates cancels
		case <-cancel:
			return

		// When a new update is queued up or channel is closed...
		case updatePkg, ok := <-pgc.outboundUpdates:
			if !ok {
				return
			}

			// Ignore if update has a different session # (the previous session ended)
			if updatePkg.sessionId != sessionId {
				continue
			}

			// Package it up into a protobuf structure and stream it out
			uxs := &pb.PGUpdate{Cbor: updatePkg.cbor}
			stream.SendMsg(uxs)
		}
	}
}

// Sends an update (as CBOR) to the app and optionally waits for an update to come back.
//
// If successful, this returns an incoming update (which could be empty) and nil for error.
// If the streaming session was re-established then nil is returned for the update and nil for error.
// If an error occurred then an error is returned along with nil for the update.
func (pgc *PGComm) ExchangeUpdates(updateCbor []byte, noWait bool) ([]byte, error) {

	// Build an new update using session # and bytes
	updateOut := UpdatePackage{sessionId: pgc.exchangeSessionId, cbor: updateCbor}

	// Queue the update to be streamed to app
	pgc.outboundUpdates <- updateOut

	var updateIn UpdatePackage
	var ok bool

	if noWait {
		select {
		case pgc.exchangeSessionId = <-pgc.newSessionId:
			goto case_new_session
		case updateIn, ok = <-pgc.inboundUpdates:
			goto case_inbound_update
		default:
			// No update available, return immediately
			return nil, nil
		}

	} else {
		select {
		case pgc.exchangeSessionId = <-pgc.newSessionId:
			goto case_new_session
		case updateIn, ok = <-pgc.inboundUpdates:
			goto case_inbound_update
		}
	}

case_new_session:
	return nil, nil

case_inbound_update:
	if !ok {
		return nil, errors.New("inboundUpdates channel is invalid")
	}

	// Ignore if sessionId of incoming update doesn't match
	// (This case may not happen in practice but, for logical reasons, we'll treat it just in case.)
	if updateIn.sessionId != pgc.exchangeSessionId {
		return nil, nil
	}

	return updateIn.cbor, nil
}

// Implementation of PGServer.StreamUpdates API call.
//
// This function is invoked from a Go routine and it must use thread safety.
func (pgc *PGComm) StreamUpdates(stream pb.PGService_StreamUpdatesServer) error {

	// Prevent simultaneous calls to this API, e.g. two different apps connected at same time.
	select {
	case pgc.busyStreaming <- true:

		notBusy := func() {
			<-pgc.busyStreaming
		}

		defer notBusy()

	default:
		return errors.New("busy streaming updates to another app")
	}

	// Launch a Go routine to stream updates back to caller
	cancel := make(chan bool)
	go pgc.streamOutboundUpdates(pgc.streamingSessionId, cancel, stream)

	// Receive updates and process them
	var err error

	// Loop for every inbound update received...
	for {
		uxs := pb.PGUpdate{}
		err = stream.RecvMsg(&uxs)

		if err != nil {
			break
		}

		// TODO:  This is mainly a debugging aid.  It should be removed for production.
		fmt.Printf("Update received with %d bytes.\n", len(uxs.Cbor))

		// Queue up a new update package
		pgc.inboundUpdates <- UpdatePackage{sessionId: pgc.streamingSessionId, cbor: uxs.Cbor}
	}

	// End the GO routine for streaming outbound updates
	cancel <- true
	close(cancel)

	// Increment the streaming session # and put into channel
	pgc.streamingSessionId = pgc.streamingSessionId + 1
	pgc.newSessionId <- pgc.streamingSessionId

	if err == io.EOF {
		return nil
	}

	return err
}

// Starts serving for gRPC calls at specified address and port.  Returns an error if it has
// problems opening a port for listening.
func (pgc *PGComm) StartServing(addr string, port int) error {

	if pgc.active_server != nil {
		return errors.New("PGComm serving already started")
	}

	// Create the channels to handle "synchronizing by communication" amoung various Go routines.
	// Note:  there should be no practical need to go beyond 2 entries for each update channel.  Also keep
	// in mind that the number of entries should not approach the dynamic range of an int8 type (which is 256).
	pgc.inboundUpdates = make(chan UpdatePackage, 2)
	pgc.outboundUpdates = make(chan UpdatePackage, 2)
	pgc.newSessionId = make(chan int8, 1)
	pgc.busyStreaming = make(chan bool, 1)

	address := fmt.Sprintf("%s:%d", addr, port)

	lis, err := net.Listen("tcp", address)

	if err != nil {
		slog.Error("could not listen for network connection", "address", address, "error", err)
		return err
	}

	pgc.active_server = grpc.NewServer()

	pb.RegisterPGServiceServer(pgc.active_server, pgc)

	slog.Info("server is now listening", "address", address)

	go func() {
		if err := pgc.active_server.Serve(lis); err != nil {
			slog.Error("error occurred while serving", "address", address, "error", err)
		}
	}()

	return nil
}

// Stops serving of gRPC calls.
func (pgc *PGComm) StopServing() {
	if pgc.active_server != nil {
		pgc.active_server.GracefulStop()

		// Close the channels
		close(pgc.inboundUpdates)
		close(pgc.outboundUpdates)
		close(pgc.newSessionId)
		close(pgc.busyStreaming)

		pgc.active_server = nil
	}
}
