// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"context"
	"errors"

	"github.com/prontogui/golib/pgcomm"
)

// ProntoGUI is the API for streaming a GUI to one or more App clients over gRPC.
//
// There are two modes of operation:
//
// Single-connection mode: call StartServingSingle, then use SetGUI, Wait,
// WaitOrCancel, and Update to manage the GUI directly. Only one client
// is served at a time, but clients may disconnect and reconnect freely.
//
// Multi-connection mode: call StartServingMultiple, then call AcceptSession
// in a loop. Each returned Session has its own SetGUI, Wait, etc. methods
// for managing that client's GUI independently.
type ProntoGUI interface {
	// StartServing is deprecated. Use StartServingSingle instead, which has
	// the exact same semantics.
	StartServing(addr string, port int) error

	// StartServingSingle begins listening at addr:port for a single client
	// session. Only one client is handled at a time; additional session
	// attempts are rejected. When the active client disconnects, a new client
	// may establish a session. Use SetGUI, Wait, WaitOrCancel, and Update to interact
	// with the connected client.
	StartServingSingle(addr string, port int) error

	// StartServingMultiple begins listening at addr:port for multiple client
	// sessions. Up to maxSessions clients may be connected simultaneously;
	// session attempts beyond this limit are rejected. Use AcceptSession to
	// obtain a Session for each connected client.
	StartServingMultiple(addr string, port int, maxSessions int) error

	// StopServing shuts down the server and closes all active connections.
	StopServing()

	// AcceptSession blocks until a new client connects and returns a Session
	// for that client. Only valid in multi-connection mode (after calling
	// StartServingMultiple); returns an error if called in single-connection mode.
	// It returns ErrCanceled if the context was canceled or ErrInterrupted if
	// the caller interrupted the operation.
	AcceptSession(ctx context.Context, interrupt chan bool) (Session, error)

	// SetGUI sets the top-level primitives that define the GUI. Single-connection
	// mode only. May be called before a client connects; the GUI will be sent
	// once the first client connects.
	SetGUI(primitives ...Primitive) error

	// Wait sends the current GUI state to the client and blocks until the client
	// sends back an update. Returns the Primitive that was updated, or an error
	// if the client disconnects. Single-connection mode only.
	Wait() (Primitive, error)

	// WaitOrCancel is like Wait but also returns if the context is canceled
	// or the interrupt channel is selected.
	// Returns error of ErrCanceled if it was canceled or ErrInterrupted if
	// interrupted by the caller.
	// Single-connection mode only.
	WaitOrCancel(ctx context.Context, interrupt chan bool) (Primitive, error)

	// Update sends the current GUI state to the client and checks for an
	// inbound update without blocking. Returns nil if no update is available.
	// Single-connection mode only.
	Update() (Primitive, error)
}

// Internal data for handling the API of this library
type _ProntoGUI struct {
	pgcomm *pgcomm.PGComm

	// The default session used by single-connection mode.
	defaultSession Session

	// True if StartServingSingle was called and the user's intent
	// is to call SetGUI, Wait, etc., directly on this struct.
	singleSessionMode bool

	// Maximum number of sessions we'll accept. This is equal to 1
	// if we're in singleSessionMode is true.
	maxSessions int

	// The current GUI primitives when operating in single session mode.
	currentGUI []Primitive

	// Channel for delivering sessions to AcceptSession
	sessionDelivery chan Session

	// True if currently serving clients
	isServing bool
}

// Deprecated: use StartServingSingle or StartServingMultiple.
func (pg *_ProntoGUI) StartServing(addr string, port int) error {
	return pg.StartServingSingle(addr, port)
}

func (pg *_ProntoGUI) StartServingSingle(addr string, port int) error {
	pg.singleSessionMode = true
	return pg.startServing(addr, port, 1)
}

func (pg *_ProntoGUI) StartServingMultiple(addr string, port int, maxSessions int) error {
	pg.singleSessionMode = false
	return pg.startServing(addr, port, maxSessions)
}

func (pg *_ProntoGUI) startServing(addr string, port int, maxSessions int) error {
	if maxSessions < 1 {
		return errors.New("maxSession must be >= 1")
	}

	pg.maxSessions = maxSessions
	pg.sessionDelivery = make(chan Session, 2)

	err := pg.pgcomm.StartServing(addr, port, pg.maxSessions)
	if err != nil {
		return err
	}

	go func() {
		for {
			// Block until the streaming API is called
			apicall, err := pg.pgcomm.AcceptStreamingAPICall()
			if err != nil {
				return
			}

			pg.sessionDelivery <- NewSession(apicall)
		}
	}()

	pg.isServing = true
	return nil
}

func (pg *_ProntoGUI) StopServing() {
	pg.pgcomm.StopServing()
	if pg.defaultSession != nil {
		pg.defaultSession = nil
	}
	pg.currentGUI = []Primitive{}
	pg.isServing = false
}

func (pg *_ProntoGUI) AcceptSession(ctx context.Context, interrupt chan bool) (Session, error) {
	if !pg.isServing {
		return nil, errors.New("not currently serving clients")
	}

	if pg.singleSessionMode {
		return nil, errors.New("AcceptSession is only available when using StartServingMultiple")
	}

	select {
	case session, ok := <-pg.sessionDelivery:
		if !ok {
			return nil, errors.New("server stopped")
		}
		return session, nil
	case <-ctx.Done():
		return nil, ErrCanceled
	case <-interrupt:
		return nil, ErrInterrupted
	}
}

func (pg *_ProntoGUI) checkForDefaultSession(block bool, ctx context.Context, interrupt chan bool) error {

	var session Session
	var ok bool

	// Was a subsequent session was started?
	select {
	case session, ok = <-pg.sessionDelivery:
		if !ok {
			return errors.New("server was stopped")
		}

	default:
	}

	if pg.defaultSession == nil && session == nil && block {
		if ctx != nil {
			// Block until a session is started or context is cancelled
			select {
			case session, ok = <-pg.sessionDelivery:
				if !ok {
					return errors.New("server was stopped")
				}
			case <-ctx.Done():
				return ErrCanceled
			case <-interrupt:
				return ErrInterrupted
			}
		} else {
			select {
			case session, ok = <-pg.sessionDelivery:
				if !ok {
					return errors.New("server was stopped")
				}
			case <-interrupt:
				return ErrInterrupted
			}
		}
	}

	if session != nil {
		// Apply any buffered SetGUI call when operating in single session mode
		if pg.singleSessionMode {
			session.SetGUI(pg.currentGUI...)
		}
		pg.defaultSession = session
	}

	return nil
}

func (pg *_ProntoGUI) SetGUI(primitives ...Primitive) error {
	if !pg.isServing {
		return errors.New("not currently serving clients")
	}

	if !pg.singleSessionMode {
		return errors.New("SetGUI is only available when using StartServingSingle")
	}

	err := pg.checkForDefaultSession(false, nil, nil)
	if err != nil {
		return err
	}

	pg.currentGUI = primitives

	if pg.defaultSession != nil {
		pg.defaultSession.SetGUI(primitives...)
	}

	return nil
}

func (pg *_ProntoGUI) Wait() (Primitive, error) {
	if !pg.isServing {
		return nil, errors.New("not currently serving clients")
	}

	if !pg.singleSessionMode {
		return nil, errors.New("Wait is only available when using StartServingSingle")
	}

	err := pg.checkForDefaultSession(true, nil, nil)
	if err != nil {
		return nil, err
	}

	if pg.defaultSession == nil {
		return nil, errors.New("no session in progress")
	}

	// Exchange updates and wait until App has an update.
	p, err := pg.defaultSession.Wait()
	if err == ErrSessionEnded {
		pg.defaultSession = nil
		err = nil
	}
	return p, err
}

func (pg *_ProntoGUI) WaitOrCancel(ctx context.Context, interrupt chan bool) (Primitive, error) {
	if !pg.isServing {
		return nil, errors.New("not currently serving clients")
	}

	if !pg.singleSessionMode {
		return nil, errors.New("WaitOrCancel is only available when using StartServingSingle")
	}

	err := pg.checkForDefaultSession(true, ctx, interrupt)
	if err != nil {
		return nil, err
	}

	if pg.defaultSession == nil {
		return nil, errors.New("no session in progress")
	}

	// Exchange updates and wait until App has an update.
	p, err := pg.defaultSession.WaitOrCancel(ctx, interrupt)
	if err == ErrSessionEnded {
		pg.defaultSession = nil
		err = nil
	}
	return p, err
}

func (pg *_ProntoGUI) Update() (Primitive, error) {
	if !pg.isServing {
		return nil, errors.New("not currently serving clients")
	}

	if !pg.singleSessionMode {
		return nil, errors.New("Update is only available when using StartServingSingle")
	}

	err := pg.checkForDefaultSession(false, nil, nil)
	if err != nil {
		return nil, err
	}

	if pg.defaultSession == nil {
		return nil, nil
	}

	p, err := pg.defaultSession.Update()
	if err == ErrSessionEnded {
		pg.defaultSession = nil
		err = nil
	}

	return p, err
}

// NewProntoGUI creates a new ProntoGUI instance.
func NewProntoGUI() ProntoGUI {
	pgcomm := pgcomm.NewPGComm()
	pg := &_ProntoGUI{pgcomm: pgcomm}
	return pg
}
