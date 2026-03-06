// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A field for entering numeric values.
type NumericFieldWith struct {
	Embodiment   string
	NumericEntry string
	Status       int
	Tag          string
}

// Creates a new NumericField primitive using the supplied field assignments.
func (w NumericFieldWith) Make() *NumericField {
	nf := &NumericField{}
	nf.embodiment.Set(w.Embodiment)
	nf.numericEntry.Set(w.NumericEntry)
	nf.status.Set(w.Status)
	nf.tag.Set(w.Tag)
	return nf
}

// A field for entering numeric values.
type NumericField struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment   StringField
	numericEntry StringField
	status       IntegerField
	tag          StringField
}

// Create a new NumericField and assign its numeric entry field.
func NewNumericField(numericEntry string) *NumericField {
	return NumericFieldWith{NumericEntry: numericEntry}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (nf *NumericField) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	nf.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &nf.embodiment},
			{key.FKey_NumericEntry, &nf.numericEntry},
			{key.FKey_Status, &nf.status},
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

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *NumericField) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *NumericField) SetStatus(i int) *NumericField {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *NumericField) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *NumericField) SetVisible(visible bool) *NumericField {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *NumericField) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *NumericField) SetEnabled(enabled bool) *NumericField {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *NumericField) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *NumericField) SetCollapsed(collapsed bool) *NumericField {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
