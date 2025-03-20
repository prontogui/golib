// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A nothing primitive displays nothing on the screen. It can be used as a placeholder.
type NothingWith struct {
	Embodiment string
	Tag        string
}

// Creates a new Nothing primitive using the supplied field assignments.
func (w NothingWith) Make() *Nothing {
	nothing := &Nothing{}
	nothing.embodiment.Set(w.Embodiment)
	nothing.tag.Set(w.Tag)
	return nothing
}

// A nothing primitive displays nothing on the screen. It can be used as a placeholder.
type Nothing struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	tag        StringField
}

// Create a new Nothing.
func NewNothing() *Nothing {
	return NothingWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (nothing *Nothing) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	nothing.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &nothing.embodiment},
			{key.FKey_Tag, &nothing.tag},
		}
	})
}

// Returns a string representation of this primitive.
// Implements of fmt:Stringer interface.
func (nothing *Nothing) String() string {
	return "."
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (nothing *Nothing) Embodiment() string {
	return nothing.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (nothing *Nothing) SetEmbodiment(s string) *Nothing {
	nothing.embodiment.Set(s)
	return nothing
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (nothing *Nothing) Tag() string {
	return nothing.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (nothing *Nothing) SetTag(s string) *Nothing {
	nothing.tag.Set(s)
	return nothing
}
