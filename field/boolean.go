// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package field

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Boolean struct {
	Reserved
	b bool
}

func (f *Boolean) GetAsAny() any {
	return f.b
}

func (f *Boolean) Get() bool {
	return f.b
}

func (f *Boolean) Set(b bool) {
	f.b = b
	f.OnSet(false)
}

func (f *Boolean) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *Boolean) EgestValue() any {
	return f.b
}

func (f *Boolean) IngestValue(value any) error {

	b, ok := value.(bool)

	if !ok {
		return errors.New("unable to convert value (any) to field value")
	}

	f.b = b
	return nil
}
