// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A check provides a yes/no, on/off, 1/0, kind of choice to the user.
// It is often represented with a check box like you would see on a form.
type CheckWith struct {
	Checked    bool
	Embodiment string
	Label      string
	LabelItem  Primitive
	Status     int
	Tag        string
}

// Makes a new Check with specified field values.
func (w CheckWith) Make() *Check {
	check := &Check{}
	check.checked.Set(w.Checked)
	check.embodiment.Set(w.Embodiment)
	check.label.Set(w.Label)
	check.labelItem.Set(w.LabelItem)
	check.status.Set(w.Status)
	check.tag.Set(w.Tag)
	return check
}

// A check provides a yes/no, on/off, 1/0, kind of choice to the user.
// It is often represented with a check box like you would see on a form.
type Check struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	checked    BooleanField
	embodiment StringField
	label      StringField
	labelItem  AnyField
	status     IntegerField
	tag        StringField
}

// Creates a new Check and assigns a label.
func NewCheck(label string) *Check {
	return CheckWith{Label: label}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (check *Check) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	check.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Checked, &check.checked},
			{key.FKey_Embodiment, &check.embodiment},
			{key.FKey_Label, &check.label},
			{key.FKey_LabelItem, &check.labelItem},
			{key.FKey_Status, &check.status},
			{key.FKey_Tag, &check.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (check *Check) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return check.LabelItem()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a string representation of this primitive:  the label.
// Implements of fmt:Stringer interface.
func (check *Check) String() string {
	return check.label.Get()
}

// Returns true if the check state is Yes, On, 1, etc., and false if the check state is No, Off, 0, etc.
func (check *Check) Checked() bool {
	return check.checked.Get()
}

// Sets the check state.
func (check *Check) SetChecked(b bool) *Check {
	check.checked.Set(b)
	return check
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (check *Check) Embodiment() string {
	return check.embodiment.Get()
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (check *Check) SetEmbodiment(s string) *Check {
	check.embodiment.Set(s)
	return check
}

// Returns the label to display in the check.
func (check *Check) Label() string {
	return check.label.Get()
}

// Sets the label to display in the check.
func (check *Check) SetLabel(s string) *Check {
	check.label.Set(s)
	return check
}

// Returns the label to display in the command.
func (check *Check) LabelItem() Primitive {
	return check.labelItem.Get()
}

// Sets the label to display in the command.
func (check *Check) SetLabelItem(item Primitive) *Check {
	check.labelItem.Set(item)
	return check
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (check *Check) Tag() string {
	return check.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (check *Check) SetTag(s string) *Check {
	check.tag.Set(s)
	return check
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Check) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Check) SetStatus(i int) *Check {
	p.status.Set(i)
	return p
}

// Returns the visibility of the Check.  This is derived from the Status field.
func (p *Check) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Check) SetVisible(visible bool) *Check {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Check) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Check) SetEnabled(enabled bool) *Check {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Check) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Check) SetCollapsed(collapsed bool) *Check {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
