// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"
	"time"

	"github.com/prontogui/golib/key"
)

func Test_EventNotIssued1(t *testing.T) {

	ts := time.Now()

	f := EventField{}
	f.TimestampProvider = func() time.Time {
		return ts
	}

	if f.Issued() {
		t.Fatal("EventField shows the event as issued.  Expecting it to be not issued.")
	}
}

func Test_EventNotIssued2(t *testing.T) {

	ts := time.Now()

	f := EventField{}
	f.TimestampProvider = func() time.Time {
		return ts
	}

	f.IngestValue(true)

	ts = time.Now().Add(time.Second)

	if f.Issued() {
		t.Fatal("EventField shows the event as issued.  Expecting it to be not issued.")
	}
}

func Test_EventIssued(t *testing.T) {

	ts := time.Now()

	f := EventField{}
	f.TimestampProvider = func() time.Time {
		return ts
	}

	f.IngestValue(true)

	if !f.Issued() {
		t.Fatal("EventField doesn't show the event as issued.  Expecting it to be issued.")
	}
}

func Test_EventPrepareForUpdates(t *testing.T) {
	ts := time.Now()

	f := EventField{}
	f.TimestampProvider = func() time.Time {
		return ts
	}

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.FieldBase)
}

func Test_EventEgestValue(t *testing.T) {

	ts := time.Now()

	f := EventField{}
	f.TimestampProvider = func() time.Time {
		return ts
	}

	v := f.EgestValue()
	b, ok := v.(bool)
	if !ok {
		t.Fatal("cannot convert return value to bool")
	}
	if b != false {
		t.Fatal("incorrect value returned")
	}
}
