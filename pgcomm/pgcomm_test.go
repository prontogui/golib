// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package pgcomm

import (
	"testing"

	"github.com/prontogui/golib/testhelp"
)

func Test_serve_badport(t *testing.T) {
	pgc := NewPGComm()
	err := pgc.StartServing("", -1, 1)
	testhelp.TestErrorMessage(t, err, "listen tcp: address -1: invalid port")
}

func Test_serve_good(t *testing.T) {
	pgc := NewPGComm()
	err := pgc.StartServing("", 0, 1)
	testhelp.TestNilError(t, err)
	pgc.StopServing()
}

func Test_accept_after_stop(t *testing.T) {
	pgc := NewPGComm()
	err := pgc.StartServing("", 0, 1)
	testhelp.TestNilError(t, err)
	pgc.StopServing()

	_, err = pgc.AcceptStreamingAPICall()
	if err == nil {
		t.Fatal("expected error from AcceptSession after StopServing")
	}
}

func Test_double_start(t *testing.T) {
	pgc := NewPGComm()
	err := pgc.StartServing("", 0, 1)
	testhelp.TestNilError(t, err)
	defer pgc.StopServing()

	err = pgc.StartServing("", 0, 1)
	testhelp.TestErrorMessage(t, err, "PGComm serving already started")
}
