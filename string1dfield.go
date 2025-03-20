// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type String1DField struct {
	FieldBase
	sa []string
}

func (f *String1DField) Get() []string {
	return f.sa
}

func (f *String1DField) Set(sa []string) {
	f.sa = sa
	f.OnSet(false)
}

func (f *String1DField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *String1DField) EgestValue() any {
	return f.sa
}

func (f *String1DField) IngestValue(value any) error {
	sa, ok := value.([]string)
	if !ok {
		return errors.New("cannot convert value to []string")
	}
	f.sa = sa
	return nil
}
