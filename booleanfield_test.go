// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_BooleanSetAndGetFalse(t *testing.T) {
	f := BooleanField{}

	f.Set(false)

	if f.Get() != false {
		t.Fatal("cannot set boolean to false and get the same value back.")
	}
}

func Test_BooleanSetAndGetTrue(t *testing.T) {
	f := BooleanField{}

	f.Set(true)

	if f.Get() != true {
		t.Fatal("cannot set boolean to true and get the same value back.")
	}
}

func Test_BooleanPrepareForUpdates(t *testing.T) {
	f := BooleanField{}

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.FieldBase)

	f.Set(true)

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_BooleanEgestValue(t *testing.T) {
	f := BooleanField{}
	f.Set(true)
	v := f.EgestValue()
	b, ok := v.(bool)
	if !ok {
		t.Fatal("cannot convert return value to bool")
	}
	if b != true {
		t.Fatal("incorrect value returned")
	}
}

func Test_BooleanIngestUpdateTrue(t *testing.T) {

	f := BooleanField{}
	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	err := f.IngestValue(true)

	testfunc := func() bool {
		return f.Get() == true
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_BooleanIngestUpdateFalse(t *testing.T) {

	f := BooleanField{}
	f.Set(true)
	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	err := f.IngestValue(false)

	testfunc := func() bool {
		return f.Get() == false
	}

	verifyIngestUpdateSuccessful(t, err, testfunc)
}

func Test_BooleanIngestUpdateInvalid(t *testing.T) {

	f := BooleanField{}
	err := f.IngestValue(10)
	verifyIngestUpdateInvalid(t, err)
}
