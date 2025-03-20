// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_Integer1DField_GetSet(t *testing.T) {
	field := &Integer1DField{}

	// Test setting and getting []int
	expected := []int{1, 2, 3}
	field.Set(expected)
	result := field.Get()

	if len(result) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected value %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func Test_Integer1DField_IngestValue(t *testing.T) {
	field := &Integer1DField{}

	tests := []struct {
		input    any
		expected []int
	}{
		{[]uint64{1, 2, 3}, []int{1, 2, 3}},
		{[]int64{4, 5, 6}, []int{4, 5, 6}},
		{[]int{7, 8, 9}, []int{7, 8, 9}},
		{[]uint{10, 11, 12}, []int{10, 11, 12}},
		{[]uint32{13, 14, 15}, []int{13, 14, 15}},
		{[]int32{16, 17, 18}, []int{16, 17, 18}},
		{[]uint16{19, 20, 21}, []int{19, 20, 21}},
		{[]int16{22, 23, 24}, []int{22, 23, 24}},
		{[]uint8{25, 26, 27}, []int{25, 26, 27}},
		{[]int8{28, 29, 30}, []int{28, 29, 30}},
	}

	for testNo, test := range tests {
		err := field.IngestValue(test.input)
		if err != nil {
			t.Errorf("Test #: %d, Unexpected error: %v", testNo, err)
		}

		result := field.Get()
		if len(result) != len(test.expected) {
			t.Fatalf("Test #: %d, Expected length %d, got %d", testNo, len(test.expected), len(result))
		}

		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("Test #: %d, Expected value %d at index %d, got %d", testNo, test.expected[i], i, v)
			}
		}
	}
}

func Test_Integer1DField_IngestValue_Error(t *testing.T) {
	field := &Integer1DField{}

	err := field.IngestValue("invalid type")
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func Test_Integer1DField_EgestValue(t *testing.T) {
	field := &Integer1DField{}
	expected := []int{1, 2, 3}
	field.Set(expected)

	result := field.EgestValue()
	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	resultSlice, ok := result.([]int)
	if !ok {
		t.Fatalf("Expected type []int, got %T", result)
	}

	if len(resultSlice) != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), len(resultSlice))
	}

	for i, v := range resultSlice {
		if v != expected[i] {
			t.Errorf("Expected value %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func Test_Integer1DField_PrepareForUpdates(t *testing.T) {
	field := &Integer1DField{}
	fkey := key.FKey_SelectedIndex
	pkey := key.NewPKey()
	fieldPKeyIndex := 0
	onset := func(pkey key.PKey, fkey key.FKey, b bool) {}

	isContainer := field.PrepareForUpdates(fkey, pkey, fieldPKeyIndex, onset)
	if isContainer {
		t.Error("Expected isContainer to be false")
	}
}
