// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// An icon primitive displays a material icon on the screen.
type IconWith struct {
	IconID     string
	Embodiment string
	Tag        string
}

// Creates a new Icon primitive using the supplied field assignments.
func (w IconWith) Make() *Icon {
	icon := &Icon{}
	icon.iconID.Set(w.IconID)
	icon.embodiment.Set(w.Embodiment)
	icon.tag.Set(w.Tag)
	return icon
}

// An icon primitive displays a material icon on the screen.
type Icon struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	iconID     StringField
	embodiment StringField
	tag        StringField
}

// Create a new Icon and assign its content.
func NewIcon(iconID string) *Icon {
	return IconWith{IconID: iconID}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (icon *Icon) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	icon.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_IconID, &icon.iconID},
			{key.FKey_Embodiment, &icon.embodiment},
			{key.FKey_Tag, &icon.tag},
		}
	})
}

// Returns a string representation of this primitive:  the iconID.
// Implements of fmt:Stringer interface.
func (icon *Icon) String() string {
	return icon.iconID.Get()
}

// Returns the icon ID.
func (icon *Icon) IconID() string {
	return icon.iconID.Get()
}

// Sets the icon ID.
func (icon *Icon) SetIconID(s string) *Icon {
	icon.iconID.Set(s)
	return icon

}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (icon *Icon) Embodiment() string {
	return icon.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (icon *Icon) SetEmbodiment(s string) *Icon {
	icon.embodiment.Set(s)
	return icon
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (icon *Icon) Tag() string {
	return icon.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (icon *Icon) SetTag(s string) *Icon {
	icon.tag.Set(s)
	return icon
}
