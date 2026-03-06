// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"context"
	"errors"
	"time"

	"github.com/prontogui/golib/pgcomm"
)

// Defined error indicating the session ended, typically when a client disconnects.
// This can be returned from Wait, WaitOrCancel, or Update functions.
var ErrSessionDisconnected = errors.New("session disconnected")

// Defined error indicating that serving has stopped.
// This can be returned from Wait, WaitOrCancel, or Update functions.
var ErrServingStopped = errors.New("serving stopped")

// Defined error indicating that the operation was canceled by way
// of the context provided in the call.  This can be returned from
// WaitOrCancel or AcceptSession functions.
var ErrCanceled = errors.New("operation canceled by provided context")

// Session represents a single client connection with its own GUI lifecycle.
type Session interface {
	// SetGUI sets the top-level primitives that define the GUI.
	SetGUI(primitives ...Primitive)

	// Wait sends the current GUI state to the client and blocks until the
	// client sends back an update. Returns the Primitive that was updated,
	// or an error if unsuccessful.
	Wait() (Primitive, error)

	// WaitOrCancel is like Wait but also returns ErrCanceled if the context
	// is canceled before an update arrives.
	WaitOrCancel(ctx context.Context) (Primitive, error)

	// Update sends the current GUI state to the client and checks for an
	// inbound update without blocking. Returns nil if no update is available.
	Update() (Primitive, error)
}

type _Session struct {
	synchro        *Synchro
	apicall        *pgcomm.StreamingAPICall
	isgui          bool
	fullupdate     bool
	eventTimestamp time.Time
}

// NewSession creates a new Session bound to the given streaming API call.
func NewSession(apicall *pgcomm.StreamingAPICall) Session {
	return &_Session{
		synchro:    NewSynchro(),
		apicall:    apicall,
		isgui:      false,
		fullupdate: true,
	}
}

// SetGUI sets the top-level primitives that define the GUI.
func (s *_Session) SetGUI(primitives ...Primitive) {
	s.fullupdate = true
	s.isgui = true
	s.synchro.SetTopPrimitives(s.getEventTimestamp, primitives...)
}

func (s *_Session) getEventTimestamp() time.Time {
	return s.eventTimestamp
}

func (s *_Session) updateEventTimestamp() {
	s.eventTimestamp = time.Now()
}

// Wait sends the current GUI state to the client and blocks until the client
// sends back an update. Returns the Primitive that was updated, or an error
// if unsuccessful such as ErrSessionDisconnected or ErrServingStopped.
func (s *_Session) Wait() (Primitive, error) {
	updateOut, err := s.getNextUpdate()
	if err != nil {
		return nil, err
	}

	// Send update to the client
	select {
	case s.apicall.Outbound <- updateOut:
		break
	case <-s.apicall.ServingStopped:
		return nil, ErrServingStopped
	}

	// Wait for inbound update from client.
	select {
	case updateIn, ok := <-s.apicall.Inbound:
		if !ok {
			s.fullupdate = true
			return nil, ErrSessionDisconnected
		}

		s.updateEventTimestamp()

		if len(updateIn) == 0 {
			return nil, nil
		}

		return s.synchro.IngestUpdate(updateIn)

	case <-s.apicall.ServingStopped:
		return nil, ErrServingStopped
	}
}

// WaitOrCancel is like Wait but also returns ErrCanceled if the provided
// context is canceled before an update arrives from the client.
func (s *_Session) WaitOrCancel(ctx context.Context) (Primitive, error) {
	updateOut, err := s.getNextUpdate()
	if err != nil {
		return nil, err
	}

	// Send update to client.
	select {
	case s.apicall.Outbound <- updateOut:
		break
	case <-s.apicall.ServingStopped:
		return nil, ErrServingStopped
	case <-ctx.Done():
		return nil, ErrCanceled
	}

	// Wait for inbound update or cancellation.
	select {
	case updateIn, ok := <-s.apicall.Inbound:
		if !ok {
			s.fullupdate = true
			return nil, ErrSessionDisconnected
		}

		s.updateEventTimestamp()

		if len(updateIn) == 0 {
			return nil, nil
		}

		p, err := s.synchro.IngestUpdate(updateIn)
		return p, err

	case <-ctx.Done():
		return nil, ErrCanceled

	case <-s.apicall.ServingStopped:
		return nil, ErrServingStopped
	}
}

// Update sends the current GUI state to the client and checks for an
// inbound update without blocking. Returns the primitive that was updated or
// nil if no update is available. Returns an error if unsuccessful such as
// ErrSessionDisconnected or ErrServingStopped.
func (s *_Session) Update() (Primitive, error) {
	updateOut, err := s.getNextUpdate()
	if err != nil {
		return nil, err
	}

	// Send update to client.
	select {
	case s.apicall.Outbound <- updateOut:
		break
	case <-s.apicall.ServingStopped:
		return nil, ErrServingStopped
	}

	// Non-blocking check for inbound update.
	select {
	case updateIn, ok := <-s.apicall.Inbound:
		if !ok {
			s.fullupdate = true
			return nil, ErrSessionDisconnected
		}

		s.updateEventTimestamp()

		if len(updateIn) == 0 {
			return nil, nil
		}
		return s.synchro.IngestUpdate(updateIn)
	default:
		return nil, nil
	}
}

func (s *_Session) getNextUpdate() ([]byte, error) {
	if !s.isgui {
		return nil, errors.New("no GUI has been set")
	}

	if s.fullupdate {
		s.fullupdate = false
		return s.synchro.GetFullUpdate()
	}
	return s.synchro.GetPartialUpdate()
}
