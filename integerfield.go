// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type IntegerField struct {
	FieldBase
	i int
}

func (f *IntegerField) Get() int {
	return f.i
}

func (f *IntegerField) Set(i int) {
	f.i = i
	f.OnSet(false)
}

func (f *IntegerField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction, etsprovider EventTimestampProvider) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset, etsprovider)
	return false
}

func (f *IntegerField) EgestValue() any {
	return f.i
}

func ConvertAnyToInt(value any) (int, error) {

	ui64, ok := value.(uint64)
	if ok {
		return int(ui64), nil
	}

	i64, ok := value.(int64)
	if ok {
		return int(i64), nil
	}

	i, ok := value.(int)
	if ok {
		return i, nil
	}

	ui, ok := value.(uint)
	if ok {
		return int(ui), nil
	}

	ui32, ok := value.(uint32)
	if ok {
		return int(ui32), nil
	}

	i32, ok := value.(int32)
	if ok {
		return int(i32), nil
	}

	ui16, ok := value.(uint16)
	if ok {
		return int(ui16), nil
	}

	i16, ok := value.(int16)
	if ok {
		return int(i16), nil
	}

	ui8, ok := value.(uint8)
	if ok {
		return int(ui8), nil
	}

	i8, ok := value.(int8)
	if ok {
		return int(i8), nil
	}

	return 0, errors.New("unable to convert value (any) to field value")
}

func (f *IntegerField) IngestValue(value any) error {

	// Unfortunately, CBOR encodes different sizes of integers based on optimum space usage.  It's not deterministic
	// what we are converting from.  So we have to test each case until a successful conversion happens.

	i, err := ConvertAnyToInt(value)
	if err == nil {
		f.i = i
	}

	return err
}
