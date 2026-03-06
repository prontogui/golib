// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A nothing primitive displays nothing on the screen. It can be used as a placeholder.
type NothingWith struct {
}

// Creates a new Nothing primitive using the supplied field assignments.
func (w NothingWith) Make() *Nothing {
	nothing := &Nothing{}
	return nothing
}

// A nothing primitive displays nothing on the screen. It can be used as a placeholder.
type Nothing struct {
	// Mix-in the common guts for primitives
	PrimitiveBase
}

// Create a new Nothing.
func NewNothing() *Nothing {
	return NothingWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (nothing *Nothing) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	nothing.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{}
	})
}

// Returns a string representation of this primitive.
// Implements of fmt:Stringer interface.
func (nothing *Nothing) String() string {
	return "."
}
