// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A field for entering numeric values.
type NumericFieldWith struct {
	Embodiment   string
	NumericEntry string
	Tag          string
}

// Creates a new NumericField primitive using the supplied field assignments.
func (w NumericFieldWith) Make() *NumericField {
	nf := &NumericField{}
	nf.embodiment.Set(w.Embodiment)
	nf.numericEntry.Set(w.NumericEntry)
	nf.tag.Set(w.Tag)
	return nf
}

// A field for entering numeric values.
type NumericField struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment   StringField
	numericEntry StringField
	tag          StringField
}

// Create a new NumericField and assign its numeric entry field.
func NewNumericField(numericEntry string) *NumericField {
	return NumericFieldWith{NumericEntry: numericEntry}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (nf *NumericField) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	nf.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &nf.embodiment},
			{key.FKey_NumericEntry, &nf.numericEntry},
			{key.FKey_Tag, &nf.tag},
		}
	})
}

// Returns a string representation of this primitive:  the numeric entry.
// Implements of fmt:Stringer interface.
func (nf *NumericField) String() string {
	return nf.numericEntry.Get()
}

// Returns the numeric entry value.
func (nf *NumericField) NumericEntry() string {
	return nf.numericEntry.Get()
}

// Sets the the numeric entry value.
func (nf *NumericField) SetNumericEntry(s string) *NumericField {
	nf.numericEntry.Set(s)
	return nf
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (nf *NumericField) Embodiment() string {
	return nf.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (nf *NumericField) SetEmbodiment(s string) *NumericField {
	nf.embodiment.Set(s)
	return nf
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (nf *NumericField) Tag() string {
	return nf.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (nf *NumericField) SetTag(s string) *NumericField {
	nf.tag.Set(s)
	return nf
}
