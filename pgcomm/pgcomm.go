// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package pgcomm

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync/atomic"

	pb "github.com/prontogui/golib/pb"
	"google.golang.org/grpc"
)

// StreamingAPICall represents a single client streaming API call with channels for communication.
type StreamingAPICall struct {
	// Streaming data coming from the client App
	Inbound chan []byte

	// Streaming data going to the client App
	Outbound chan []byte

	// Signals that API call has exited
	CallHasExited chan byte
}

// Implementation of the PGServer
type PGComm struct {
	pb.UnimplementedPGServiceServer

	// The active server.
	activeServer *grpc.Server

	// Channel for delivering new session connections to AcceptSession callers.
	acceptChan chan *StreamingAPICall

	// Signal that all streaming API calls should stop
	StopAllStreaming chan bool

	// Maximum allowed streaming API calls
	maxAPICalls int

	// Number of active streaming API calls
	activeCalls int64
}

func NewPGComm() *PGComm {
	return &PGComm{}
}

// Blocks until a client connects or the server stops.
// Returns a StreamingAPICall for the new client, or an error if the server has stopped.
func (pgc *PGComm) AcceptStreamingAPICall() (*StreamingAPICall, error) {
	select {
	case apicall, ok := <-pgc.acceptChan:
		if !ok {
			return nil, errors.New("server stopped")
		}
		return apicall, nil
	case <-pgc.StopAllStreaming:
		return nil, errors.New("server stopped")
	}
}

// Implementation of PGServer.StreamUpdates API call.
//
// This function is invoked from a Go routine for each client that connects.
func (pgc *PGComm) StreamUpdates(stream grpc.BidiStreamingServer[pb.PGUpdate, pb.PGUpdate]) error {

	if atomic.LoadInt64(&pgc.activeCalls) >= int64(pgc.maxAPICalls) {
		// At limit, reject the connection.
		return errors.New("too many connections already - try again later")
	}

	atomic.AddInt64(&pgc.activeCalls, 1)

	defer func() {
		atomic.AddInt64(&pgc.activeCalls, -1)
	}()

	apicall := &StreamingAPICall{
		Inbound:       make(chan []byte, 2),
		Outbound:      make(chan []byte, 2),
		CallHasExited: make(chan byte),
	}

	// Deliver this session to AcceptSession.
	pgc.acceptChan <- apicall

	cancelOutbound := make(chan bool)

	// Launch goroutine to send outbound updates to the client.
	go func() {
		for {
			select {
			case <-cancelOutbound:
				return
			case update, ok := <-apicall.Outbound:
				if !ok {
					return
				}
				uxs := &pb.PGUpdate{Cbor: update}
				if err := stream.SendMsg(uxs); err != nil {
					return
				}
			}
		}
	}()

	// Receive updates from the client.
	var err error
	for {
		uxs := pb.PGUpdate{}
		err = stream.RecvMsg(&uxs)
		if err != nil {
			break
		}

		select {
		case apicall.Inbound <- uxs.Cbor:
			continue
		case <-pgc.StopAllStreaming:
			goto cleanup
		}
	}

cleanup:
	close(apicall.CallHasExited)

	close(apicall.Inbound)
	close(cancelOutbound)

	if err == io.EOF {
		return nil
	}
	return err
}

// Starts serving for gRPC calls at specified address and port.
func (pgc *PGComm) StartServing(addr string, port int, maxAPICalls int) error {

	if pgc.activeServer != nil {
		return errors.New("PGComm serving already started")
	}

	pgc.maxAPICalls = maxAPICalls
	pgc.acceptChan = make(chan *StreamingAPICall, pgc.maxAPICalls)
	pgc.StopAllStreaming = make(chan bool)

	address := fmt.Sprintf("%s:%d", addr, port)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		slog.Error("could not listen for network connection", "address", address, "error", err)
		return err
	}

	pgc.activeServer = grpc.NewServer()
	pb.RegisterPGServiceServer(pgc.activeServer, pgc)

	slog.Info("server is now listening", "address", address)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				slog.Info("server stopped", "address", address)
			}
		}()
		if err := pgc.activeServer.Serve(lis); err != nil {
			slog.Error("error occurred while serving", "address", address, "error", err)
		}
	}()

	return nil
}

// Stops serving of gRPC calls.
func (pgc *PGComm) StopServing() {
	if pgc.activeServer != nil {

		// Signal to all active API calls to quit what their doing
		close(pgc.StopAllStreaming)

		// Wait until all active API calls have returned
		// TODO

		pgc.activeServer.Stop()
		close(pgc.acceptChan)
		pgc.activeServer = nil
	}
}

/*
func (pgc *PGComm) IsServing() bool {
	return pgc.active_server != nil
}
*/
