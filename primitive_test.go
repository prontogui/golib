// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"slices"
	"testing"

	"github.com/prontogui/golib/key"
)

func _areFieldsAttachedAlphabetically(res PrimitiveBase) bool {

	attachedOrder := []string{}

	for _, fr := range res.fields {
		fieldName := key.FieldnameFor(fr.fkey)
		attachedOrder = append(attachedOrder, fieldName)
	}

	return slices.IsSorted(attachedOrder)
}

func verifyAllFieldsAttached(t *testing.T, res PrimitiveBase, fields ...string) {

	verifyFieldAttached := func(fields ...string) {
		for _, field := range fields {
			if res.findField(key.FKeyFor(field)) == nil {
				t.Errorf("Field '%s' is not attached to primitive.", field)
			}
		}
	}

	verifyFieldAttached(fields...)

	if !_areFieldsAttachedAlphabetically(res) {
		t.Error("fields were not attached in alphabetical order")
	}
}
