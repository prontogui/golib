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

type AnyField struct {
	FieldBase
	p Primitive
}

func (f *AnyField) prepareDescendantForUpdates() {
	if f.p != nil {
		if f.onset == nil {
			f.p.PrepareForUpdates(key.EmptyPKey(), nil, f.etsprovider)
		} else {
			f.p.PrepareForUpdates(f.pkey.AddLevel(f.fieldPKeyIndex), f.onset, f.etsprovider)
		}
	}
}

func (f *AnyField) unprepareDescendantForUpdates() {
	if f.p != nil {
		f.p.UnprepareForUpdates()
	}
}

func (f *AnyField) Get() Primitive {
	return f.p
}

func (f *AnyField) Set(p Primitive) {
	f.unprepareDescendantForUpdates()
	f.p = p
	f.prepareDescendantForUpdates()
	f.OnSet(true)
}

func (f *AnyField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction, etsprovider EventTimestampProvider) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset, etsprovider)
	f.prepareDescendantForUpdates()
	return true
}

func (f *AnyField) UnprepareForUpdates() {
	f.ClearUpdateInfo()
	f.unprepareDescendantForUpdates()
}

func (f *AnyField) EgestValue() any {
	if f.p != nil {
		return f.p.EgestUpdate(true, nil)
	} else {
		return nil
	}
}

func (f *AnyField) IngestValue(value any) error {

	m, ok := value.(map[any]any)
	if !ok {
		return errors.New("invalid update")
	}

	if f.p != nil {
		return f.p.IngestUpdate(m)
	}

	return nil
}
