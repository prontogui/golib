// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A group is a related set of primitives, such as a group of commands, that is
// static in type and quantity.  If a dynamic number of primitives is desired
// then consider using a List primitive instead.
type GroupWith struct {
	Embodiment string
	GroupItems []Primitive
	status     int
	Tag        string
}

// Creates a new Group using the supplied field assignments.
func (w GroupWith) Make() *Group {
	grp := &Group{}
	grp.embodiment.Set(w.Embodiment)
	grp.groupItems.Set(w.GroupItems)
	grp.status.Set(w.status)
	grp.tag.Set(w.Tag)
	return grp
}

// A group is a related set of primitives, such as a group of commands, that is
// static in type and quantity.  If a dynamic number of primitives is desired
// then consider using a List primitive instead.
type Group struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	groupItems Any1DField
	status     IntegerField
	tag        StringField
}

// Creates a new Group and assigns items.
func NewGroup(items ...Primitive) *Group {
	return GroupWith{GroupItems: items}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (grp *Group) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	grp.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &grp.embodiment},
			{key.FKey_GroupItems, &grp.groupItems},
			{key.FKey_Status, &grp.status},
			{key.FKey_Tag, &grp.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (grp *Group) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	if locator.NextIndex() == 0 {
		return grp.GroupItems()[locator.NextIndex()]
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (grp *Group) Embodiment() string {
	return grp.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (grp *Group) SetEmbodiment(s string) *Group {
	grp.embodiment.Set(s)
	return grp
}

// Returns the collection of primitives that make up the group.
func (grp *Group) GroupItems() []Primitive {
	return grp.groupItems.Get()
}

// Sets the collection of primitives that make up the group.
func (grp *Group) SetGroupItems(items []Primitive) *Group {
	grp.groupItems.Set(items)
	return grp
}

// Sets the collection of primitives (a variadic argument list) that make up the group.
func (grp *Group) SetGroupItemsVA(items ...Primitive) *Group {
	grp.groupItems.Set(items)
	return grp
}

// Returns the status of the command:  0 = Command Normal, 1 = Command Disabled, 2 = Command Hidden.
func (grp *Group) Status() int {
	return grp.status.Get()
}

// Sets the status of the command:  0 = Command Normal, 1 = Command Disabled, 2 = Command Hidden.
func (grp *Group) SetStatus(i int) *Group {
	grp.status.Set(i)
	return grp
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Groups inside other containers.
func (grp *Group) Tag() string {
	return grp.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Groups inside other containers.
func (grp *Group) SetTag(s string) *Group {
	grp.tag.Set(s)
	return grp
}

// Returns the visibility of the group.  This is derived from the Status field.
func (grp *Group) Visible() bool {
	return grp.status.Get() != 2
}

// Sets the visibility of the group.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (grp *Group) SetVisible(visible bool) *Group {
	if visible {
		grp.status.Set(0)
	} else {
		grp.status.Set(2)
	}
	return grp
}

// Returns the enabled status of the group.  This is derived from the Status field.
func (grp *Group) Enabled() bool {
	return grp.status.Get() == 0
}

// Sets the enabled status of the group.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (grp *Group) SetEnabled(enabled bool) *Group {
	if enabled {
		grp.status.Set(0)
	} else {
		grp.status.Set(1)
	}
	return grp
}
