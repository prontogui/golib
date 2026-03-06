// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"context"
	"testing"

	"github.com/prontogui/golib/pgcomm"
)

func newTestSession() (Session, *pgcomm.StreamingAPICall) {
	apicall := &pgcomm.StreamingAPICall{
		Inbound:  make(chan []byte, 2),
		Outbound: make(chan []byte, 2),
	}
	s := NewSession(apicall)
	return s, apicall
}

func Test_Session_Wait_NoGUI(t *testing.T) {
	s, _ := newTestSession()

	_, err := s.Wait()
	if err == nil {
		t.Fatal("expected error when no GUI set")
	}
}

func Test_Session_Wait_Disconnected(t *testing.T) {
	s, _ := newTestSession()

	txt := TextWith{Content: "hello"}.Make()
	s.SetGUI(txt)

	_, err := s.Wait()
	if err == nil {
		t.Fatal("expected error on disconnected session")
	}
}

func Test_Session_Wait_RoundTrip(t *testing.T) {
	s, conn := newTestSession()

	cmd := CommandWith{Label: "OK"}.Make()
	s.SetGUI(cmd)

	// Simulate client: read the outbound update, send back an empty update
	go func() {
		<-conn.Outbound
		// Send empty update (no event)
		conn.Inbound <- []byte{}
	}()

	p, err := s.Wait()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != nil {
		t.Fatal("expected nil primitive for empty update")
	}
}

func Test_Session_Update_NoEvent(t *testing.T) {
	s, conn := newTestSession()

	cmd := CommandWith{Label: "OK"}.Make()
	s.SetGUI(cmd)

	// Drain the outbound in background
	go func() {
		<-conn.Outbound
	}()

	p, err := s.Update()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p != nil {
		t.Fatal("expected nil primitive when no inbound update")
	}
}

func Test_Session_WaitOrCancel_Cancelled(t *testing.T) {
	s, conn := newTestSession()

	cmd := CommandWith{Label: "OK"}.Make()
	s.SetGUI(cmd)

	ctx, cancel := context.WithCancel(context.Background())

	// Drain outbound then cancel context
	go func() {
		<-conn.Outbound
		cancel()
	}()

	_, err := s.WaitOrCancel(ctx)
	if err != ErrCanceled {
		t.Fatalf("expecting ErrCanceled to be returned; got unexpected error: %v", err)
	}
}
