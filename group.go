// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

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
	Status     int
	Tag        string
}

// Creates a new Group using the supplied field assignments.
func (w GroupWith) Make() *Group {
	grp := &Group{}
	grp.embodiment.Set(w.Embodiment)
	grp.groupItems.Set(w.GroupItems)
	grp.status.Set(w.Status)
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
func (grp *Group) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	grp.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
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

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Group) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Group) SetStatus(i int) *Group {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Group) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Group) SetVisible(visible bool) *Group {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Group) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Group) SetEnabled(enabled bool) *Group {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Group) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Group) SetCollapsed(collapsed bool) *Group {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
