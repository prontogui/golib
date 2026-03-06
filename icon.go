// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// An icon primitive displays a material icon on the screen.
type IconWith struct {
	IconID     string
	Embodiment string
	Status     int
	Tag        string
}

// Creates a new Icon primitive using the supplied field assignments.
func (w IconWith) Make() *Icon {
	icon := &Icon{}
	icon.iconID.Set(w.IconID)
	icon.embodiment.Set(w.Embodiment)
	icon.status.Set(w.Status)
	icon.tag.Set(w.Tag)
	return icon
}

// An icon primitive displays a material icon on the screen.
type Icon struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	iconID     StringField
	embodiment StringField
	status     IntegerField
	tag        StringField
}

// Create a new Icon and assign its content.
func NewIcon(iconID string) *Icon {
	return IconWith{IconID: iconID}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (icon *Icon) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	icon.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_IconID, &icon.iconID},
			{key.FKey_Embodiment, &icon.embodiment},
			{key.FKey_Status, &icon.status},
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

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Icon) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Icon) SetStatus(i int) *Icon {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Icon) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Icon) SetVisible(visible bool) *Icon {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Icon) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Icon) SetEnabled(enabled bool) *Icon {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Icon) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Icon) SetCollapsed(collapsed bool) *Icon {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
