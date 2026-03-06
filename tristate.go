// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A choice presented to the user that give three possible states:  Affirmative (Yes, On, 1, etc.),
// Negative (No, Off, 0, etc.), and Indeterminate.
type TristateWith struct {
	Embodiment string
	Label      string
	LabelItem  Primitive
	State      int
	Status     int
	Tag        string
}

// Creates a new TriState using the supplied field assignments.
func (w TristateWith) Make() *Tristate {
	tri := &Tristate{}
	tri.embodiment.Set(w.Embodiment)
	tri.label.Set(w.Label)
	tri.labelItem.Set(w.LabelItem)
	tri.state.Set(w.State)
	tri.status.Set(w.Status)
	tri.tag.Set(w.Tag)
	return tri
}

// A choice presented to the user that give three possible states:  Affirmative (Yes, On, 1, etc.),
// Negative (No, Off, 0, etc.), and Indeterminate.
type Tristate struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	label      StringField
	labelItem  AnyField
	state      IntegerField
	status     IntegerField
	tag        StringField
}

// Create a new TriState and assign a label.
func NewTristate(label string) *Tristate {
	return TristateWith{Label: label}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (tri *Tristate) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	tri.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tri.embodiment},
			{key.FKey_Label, &tri.label},
			{key.FKey_LabelItem, &tri.labelItem},
			{key.FKey_State, &tri.state},
			{key.FKey_Status, &tri.status},
			{key.FKey_Tag, &tri.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (tri *Tristate) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return tri.LabelItem()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a string representation of this primitive:  the label.
// Implements of fmt:Stringer interface.
func (tri *Tristate) String() string {
	return tri.label.Get()
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (tri *Tristate) Embodiment() string {
	return tri.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (tri *Tristate) SetEmbodiment(s string) *Tristate {
	tri.embodiment.Set(s)
	return tri
}

// Returns the label to display along with the tristate option.
func (tri *Tristate) Label() string {
	return tri.label.Get()
}

// Sets the label to display along with the tristate option.
func (tri *Tristate) SetLabel(s string) *Tristate {
	tri.label.Set(s)
	return tri
}

// Returns the label to display in the command.
func (tri *Tristate) LabelItem() Primitive {
	return tri.labelItem.Get()
}

// Sets the label to display in the command.
func (tri *Tristate) SetLabelItem(item Primitive) *Tristate {
	tri.labelItem.Set(item)
	return tri
}

// Returns the state of the option (0 = Negative, 1 = Affirmative, and -1 = Indeterminate).
func (tri *Tristate) State() int {
	return tri.state.Get()
}

// Sets the state of the option (0 = Negative, 1 = Affirmative, and -1 = Indeterminate).
func (tri *Tristate) SetState(i int) *Tristate {
	tri.state.Set(i)
	return tri
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tristates as Table cells.
func (tri *Tristate) Tag() string {
	return tri.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tristates as Table cells.
func (tri *Tristate) SetTag(s string) *Tristate {
	tri.tag.Set(s)
	return tri
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Tristate) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Tristate) SetStatus(i int) *Tristate {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Tristate) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Tristate) SetVisible(visible bool) *Tristate {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Tristate) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Tristate) SetEnabled(enabled bool) *Tristate {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Tristate) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Tristate) SetCollapsed(collapsed bool) *Tristate {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
