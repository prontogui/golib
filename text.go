// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A text primitive displays text on the screen.
type TextWith struct {
	Content    string
	Embodiment string
	Status     int
	Tag        string
}

// Creates a new Text primitive using the supplied field assignments.
func (w TextWith) Make() *Text {
	text := &Text{}
	text.content.Set(w.Content)
	text.embodiment.Set(w.Embodiment)
	text.status.Set(w.Status)
	text.tag.Set(w.Tag)
	return text
}

// A text primitive displays text on the screen.
type Text struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	content    StringField
	embodiment StringField
	status     IntegerField
	tag        StringField
}

// Create a new Text and assign its content.
func NewText(content string) *Text {
	return TextWith{Content: content}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (txt *Text) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	txt.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Content, &txt.content},
			{key.FKey_Embodiment, &txt.embodiment},
			{key.FKey_Status, &txt.status},
			{key.FKey_Tag, &txt.tag},
		}
	})
}

// Returns a string representation of this primitive:  the content.
// Implements of fmt:Stringer interface.
func (txt *Text) String() string {
	return txt.content.Get()
}

// Returns the text content to display.
func (txt *Text) Content() string {
	return txt.content.Get()
}

// Sets the text content to display.
func (txt *Text) SetContent(s string) *Text {
	txt.content.Set(s)
	return txt
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (txt *Text) Embodiment() string {
	return txt.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (txt *Text) SetEmbodiment(s string) *Text {
	txt.embodiment.Set(s)
	return txt
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (txt *Text) Tag() string {
	return txt.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Texts as Table cells.
func (txt *Text) SetTag(s string) *Text {
	txt.tag.Set(s)
	return txt
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Text) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Text) SetStatus(i int) *Text {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Text) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Text) SetVisible(visible bool) *Text {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Text) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Text) SetEnabled(enabled bool) *Text {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Text) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Text) SetCollapsed(collapsed bool) *Text {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
