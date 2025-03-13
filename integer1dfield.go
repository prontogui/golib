// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

type Integer1DField struct {
	FieldBase
	ia []int
}

func (f *Integer1DField) Get() []int {
	return f.ia
}

func (f *Integer1DField) Set(ia []int) {
	f.ia = ia
	f.OnSet(false)
}

func (f *Integer1DField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *Integer1DField) EgestValue() any {
	return f.ia
}

func (f *Integer1DField) IngestValue(value any) error {

	ui64_a, ok := value.([]uint64)
	if ok {
		f.ia = make([]int, len(ui64_a))
		for i, ui64 := range ui64_a {
			f.ia[i] = int(ui64)
		}
		return nil
	}

	i64_a, ok := value.([]int64)
	if ok {
		f.ia = make([]int, len(i64_a))
		for i, i64 := range i64_a {
			f.ia[i] = int(i64)
		}
		return nil
	}

	i_a, ok := value.([]int)
	if ok {
		f.ia = i_a
		return nil
	}

	ui_a, ok := value.([]uint)
	if ok {
		f.ia = make([]int, len(ui_a))
		for i, ui := range ui_a {
			f.ia[i] = int(ui)
		}
		return nil
	}

	ui32_a, ok := value.([]uint32)
	if ok {
		f.ia = make([]int, len(ui32_a))
		for i, ui32 := range ui32_a {
			f.ia[i] = int(ui32)
		}
		return nil
	}

	i32_a, ok := value.([]int32)
	if ok {
		f.ia = make([]int, len(i32_a))
		for i, i32 := range i32_a {
			f.ia[i] = int(i32)
		}
		return nil
	}

	ui16_a, ok := value.([]uint16)
	if ok {
		f.ia = make([]int, len(ui16_a))
		for i, ui16 := range ui16_a {
			f.ia[i] = int(ui16)
		}
		return nil
	}

	i16_a, ok := value.([]int16)
	if ok {
		f.ia = make([]int, len(i16_a))
		for i, i16 := range i16_a {
			f.ia[i] = int(i16)
		}
		return nil
	}

	ui8_a, ok := value.([]uint8)
	if ok {
		f.ia = make([]int, len(ui8_a))
		for i, ui8 := range ui8_a {
			f.ia[i] = int(ui8)
		}
		return nil
	}

	i8_a, ok := value.([]int8)
	if ok {
		f.ia = make([]int, len(i8_a))
		for i, i8 := range i8_a {
			f.ia[i] = int(i8)
		}
		return nil
	}

	return errors.New("cannot convert value to []int")
}
