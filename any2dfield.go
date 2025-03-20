// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

// TODO:  swap any type with primitive.Interface and update the test accordingly.
type Any2DField struct {
	FieldBase
	ary [][]Primitive
}

func (f *Any2DField) prepareDescendantsForUpdate() {

	fieldPkey := f.pkey.AddLevel(f.fieldPKeyIndex)

	if f.onset == nil {
		for _, p1 := range f.ary {
			for _, p2 := range p1 {
				p2.PrepareForUpdates(key.EmptyPKey(), nil)
			}
		}
	} else {
		for i, p1 := range f.ary {
			pkeyi := fieldPkey.AddLevel(i)

			for j, p2 := range p1 {
				pkeyj := pkeyi.AddLevel(j)
				p2.PrepareForUpdates(pkeyj, f.onset)
			}
		}
	}
}

func (f *Any2DField) Get() [][]Primitive {
	return f.ary
}

func (f *Any2DField) Set(ary [][]Primitive) {
	f.ary = ary
	f.prepareDescendantsForUpdate()
	f.OnSet(true)
}

// Convenience function that returns the length of the array in rows.
func (f *Any2DField) Length() int {
	return len(f.ary)
}

func (f *Any2DField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	f.prepareDescendantsForUpdate()
	return true
}

func (f *Any2DField) EgestValue() any {
	ary := [][]any{}

	for _, row := range f.ary {
		ary2 := []any{}

		for _, cell := range row {
			ary2 = append(ary2, cell.EgestUpdate(true, nil))
		}

		ary = append(ary, ary2)
	}

	return ary
}

func (f *Any2DField) IngestValue(value any) error {

	ary, ok := value.([][]any)
	if !ok {
		return errors.New("invalid update")
	}

	if len(ary) != len(f.ary) {
		return errors.New("number of primitives in update does not equal existing primitives")
	}

	for i, row := range ary {

		if len(row) != len(f.ary) {
			return errors.New("number of primitives in update does not equal existing primitives")
		}

		for j, v := range row {

			m, ok := v.(map[any]any)
			if !ok {
				return errors.New("invalid update")
			}

			err := f.ary[i][j].IngestUpdate(m)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
