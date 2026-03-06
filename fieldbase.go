// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

type FieldBase struct {
	// PKey of this field's container primitive.
	pkey key.PKey // `cbor:"omitempty"`

	// FKey of this field.
	fkey key.FKey

	// The function to call to notify the field was updated.
	onset func(key.PKey, key.FKey, bool)

	// This field's pkey index relative to its container primitive (if this field contains primitives).
	// It is used when assigning new primitives to this field.
	fieldPKeyIndex int

	// Provider for event timestamps
	etsprovider EventTimestampProvider
}

func (f *FieldBase) StashUpdateInfo(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction, etsprovider EventTimestampProvider) {
	f.fkey = fkey
	f.pkey = pkey
	f.onset = onset
	f.fieldPKeyIndex = fieldPKeyIndex
	f.etsprovider = etsprovider
}

func (f *FieldBase) UnprepareForUpdates() {
	f.ClearUpdateInfo()
}

func (f *FieldBase) ClearUpdateInfo() {
	f.pkey = key.EmptyPKey()
	f.onset = nil
	f.fieldPKeyIndex = -1
}

func (f *FieldBase) OnSet(structural bool) {
	if f.onset != nil {
		f.onset(f.pkey, f.fkey, structural)
	}
}
