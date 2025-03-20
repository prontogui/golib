// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package pgcomm

import (
	"testing"

	"github.com/prontogui/golib/testhelp"
)

func Test_serve_badport(t *testing.T) {
	pgc := NewPGComm()
	err := pgc.StartServing("", -1)
	testhelp.TestErrorMessage(t, err, "listen tcp: address -1: invalid port")
}

func Test_serve_good(t *testing.T) {
	pgc := NewPGComm()
	err := pgc.StartServing("", 0)
	testhelp.TestNilError(t, err)
	pgc.StopServing()
}

/*
BROKEN TEST - pgc.outboundUpdates is nil until pgc.StartServing() is called.

// Test the normal exchange of updates between server and the app.
func Test_ExchangeUpdates1(t *testing.T) {
	pgc := NewPGComm()

	go func() {
		update := <-pgc.outboundUpdates
		pgc.inboundUpdates <- update
	}()

	updateIn, err := pgc.ExchangeUpdates([]byte{1, 2}, false)

	if err != nil {
		t.Fatal("error was returned.  Expected no error")
	}
	if !reflect.DeepEqual(updateIn, []byte{1, 2}) {
		t.Fatal("wrong update was returned")
	}
}
*/

/*
BROKEN TEST - pgc.outboundUpdates is nil until pgc.StartServing() is called.

// Test proper handling of the inboundUpdates channel being closed during an exchange.
func Test_ExchangeUpdates2(t *testing.T) {
	pgc := NewPGComm()

	go func() {
		<-pgc.outboundUpdates
		close(pgc.inboundUpdates)
	}()

	_, err := pgc.ExchangeUpdates([]byte{1, 2}, false)

	if err == nil {
		t.Fatal("no error was returned.  Expected an error")
	}

	if err.Error() != "inboundUpdates channel is invalid" {
		t.Fatal("wrong error was returned")
	}
}
*/
