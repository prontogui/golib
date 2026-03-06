// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// An entry field that allows the user to enter text.
type TextFieldWith struct {
	Embodiment string
	Status     int
	Tag        string
	TextEntry  string
}

// Creates a new TextField using the supplied field assignments.
func (w TextFieldWith) Make() *TextField {
	textField := &TextField{}
	textField.embodiment.Set(w.Embodiment)
	textField.status.Set(w.Status)
	textField.textEntry.Set(w.TextEntry)
	textField.tag.Set(w.Tag)
	return textField
}

// An entry field that allows the user to enter text.
type TextField struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	status     IntegerField
	tag        StringField
	textEntry  StringField
}

// Create a new TextField with initial text.
func NewTextField(textEntry string) *TextField {
	return TextFieldWith{TextEntry: textEntry}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (txt *TextField) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	txt.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &txt.embodiment},
			{key.FKey_Status, &txt.status},
			{key.FKey_Tag, &txt.tag},
			{key.FKey_TextEntry, &txt.textEntry},
		}
	})
}

// Returns a string representation of this primitive:  the text entry.
// Implements of fmt:Stringer interface.
func (txt *TextField) String() string {
	return txt.textEntry.Get()
}

// Returns the text entered by the user.
func (txt *TextField) TextEntry() string {
	return txt.textEntry.Get()
}

// Sets the text entered by the user.
func (txt *TextField) SetTextEntry(s string) *TextField {
	txt.textEntry.Set(s)
	return txt
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (txt *TextField) Embodiment() string {
	return txt.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (txt *TextField) SetEmbodiment(s string) *TextField {
	txt.embodiment.Set(s)
	return txt
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using TextFields as Table cells.
func (txt *TextField) Tag() string {
	return txt.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using TextFields as Table cells.
func (txt *TextField) SetTag(s string) *TextField {
	txt.tag.Set(s)
	return txt
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *TextField) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *TextField) SetStatus(i int) *TextField {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *TextField) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *TextField) SetVisible(visible bool) *TextField {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *TextField) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *TextField) SetEnabled(enabled bool) *TextField {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *TextField) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *TextField) SetCollapsed(collapsed bool) *TextField {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
