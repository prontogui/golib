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
}

// Creates a new Group using the supplied field assignments.
func (w GroupWith) Make() *Group {
	grp := &Group{}
	grp.embodiment.Set(w.Embodiment)
	grp.groupItems.Set(w.GroupItems)
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
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (grp *Group) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.
	if locator.NextIndex() != 0 {
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
	return grp.GroupItems()[locator.NextIndex()]
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
