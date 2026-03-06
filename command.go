// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A command is used to handle momentary requests by the user such that, when the command is issued,
// the service does something useful.  It is often rendered as a button with clear boundaries that
// suggest it can be clicked.
type CommandWith struct {
	Embodiment string
	Label      string
	LabelItem  Primitive
	Status     int
	Tag        string
}

// Makes a new Command with specified field values.
func (w CommandWith) Make() *Command {
	cmd := &Command{}

	// Set fields
	cmd.embodiment.Set(w.Embodiment)
	cmd.label.Set(w.Label)
	cmd.labelItem.Set(w.LabelItem)
	cmd.status.Set(w.Status)
	cmd.tag.Set(w.Tag)

	return cmd
}

// A command is used to handle momentary requests by the user such that, when the command is issued,
// the service does something useful.  It is often rendered as a button with clear boundaries that
// suggest it can be clicked.
type Command struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	commandIssued EventField
	embodiment    StringField
	label         StringField
	labelItem     AnyField
	status        IntegerField
	tag           StringField
}

// Creates a new command and assigns a label.
func NewCommand(label string) *Command {
	cmd := CommandWith{Label: label}.Make()
	return cmd
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	cmd.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_CommandIssued, &cmd.commandIssued},
			{key.FKey_Embodiment, &cmd.embodiment},
			{key.FKey_Label, &cmd.label},
			{key.FKey_LabelItem, &cmd.labelItem},
			{key.FKey_Status, &cmd.status},
			{key.FKey_Tag, &cmd.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (cmd *Command) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return cmd.LabelItem()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a string representation of this primitive:  the label.
// Implements of fmt:Stringer interface.
func (cmd *Command) String() string {
	return cmd.label.Get()
}

// Returns true if the command was issued during the current Wait cycle.
func (cmd *Command) Issued() bool {
	return cmd.commandIssued.Issued()
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (cmd *Command) Embodiment() string {
	return cmd.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (cmd *Command) SetEmbodiment(s string) *Command {
	cmd.embodiment.Set(s)
	return cmd
}

// Returns the label to display in the command.
func (cmd *Command) Label() string {
	return cmd.label.Get()
}

// Sets the label to display in the command.
func (cmd *Command) SetLabel(s string) *Command {
	cmd.label.Set(s)
	return cmd
}

// Returns the label to display in the command.
func (cmd *Command) LabelItem() Primitive {
	return cmd.labelItem.Get()
}

// Sets the label to display in the command.
func (cmd *Command) SetLabelItem(item Primitive) *Command {
	cmd.labelItem.Set(item)
	return cmd
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (cmd *Command) Tag() string {
	return cmd.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (cmd *Command) SetTag(s string) *Command {
	cmd.tag.Set(s)
	return cmd
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Command) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Command) SetStatus(i int) *Command {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Command) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Command) SetVisible(visible bool) *Command {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Command) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Command) SetEnabled(enabled bool) *Command {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Command) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Command) SetCollapsed(collapsed bool) *Command {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
