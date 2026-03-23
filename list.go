// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A list is a collection of primitives that have a sequential-like relationship
// and might be dynamic in quantity or kind.
type ListWith struct {
	Embodiment      string
	ListItems       []Primitive
	ModelFolder Primitive
	ModelItem       Primitive
	SelectedItems   []int
	SelectionMode   int
	Status          int
	Tag             string
}

// Creates a new List using the supplied field assignments.
func (w ListWith) Make() *List {
	list := &List{}

	list.embodiment.Set(w.Embodiment)
	list.listItems.Set(w.ListItems)
	list.modelFolder.Set(w.ModelFolder)
	list.modelItem.Set(w.ModelItem)
	list.selectedItems.Set(w.SelectedItems)
	list.selectionMode.Set(w.SelectionMode)
	list.status.Set(w.Status)
	list.tag.Set(w.Tag)
	return list
}

// A list is a collection of primitives that have a sequential-like relationship
// and might be dynamic in quantity or kind.
type List struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment       StringField
	listItems        Any1DField
	modelFolder  AnyField
	modelItem        AnyField
	selectedItems    Integer1DField
	selectionMode    IntegerField
	selectionChanged EventField
	status           IntegerField
	tag              StringField
}

// Creates a new List and assigns items.
func NewList(items ...Primitive) *List {
	return ListWith{ListItems: items}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (list *List) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	list.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &list.embodiment},
			{key.FKey_ListItems, &list.listItems},
			{key.FKey_ModelFolder, &list.modelFolder},
			{key.FKey_ModelItem, &list.modelItem},
			{key.FKey_SelectedItems, &list.selectedItems},
			{key.FKey_SelectionChanged, &list.selectionChanged},
			{key.FKey_SelectionMode, &list.selectionMode},
			{key.FKey_Status, &list.status},
			{key.FKey_Tag, &list.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (list *List) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return list.ListItems()[locator.NextIndex()]
	case 1:
		return list.ModelFolder()
	case 2:
		return list.ModelItem()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (list *List) Embodiment() string {
	return list.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (list *List) SetEmbodiment(s string) *List {
	list.embodiment.Set(s)
	return list
}

// Returns the ttems to show in the list.
func (list *List) ListItems() []Primitive {
	return list.listItems.Get()
}

// Sets the items to show in the list.
func (list *List) SetListItems(items []Primitive) *List {
	list.listItems.Set(items)
	return list
}

// Returns the items to show in the list (as a variadic argument list).
func (list *List) SetListItemsVA(items ...Primitive) *List {
	list.listItems.Set(items)
	return list
}

// Returns the model folder item.
func (list *List) ModelFolder() Primitive {
	return list.modelFolder.Get()
}

// Sets the model folder item.
func (list *List) SetModelFolder(item Primitive) *List {
	list.modelFolder.Set(item)
	return list
}

// Returns the model item.
func (list *List) ModelItem() Primitive {
	return list.modelItem.Get()
}

// Sets the model item.
func (list *List) SetModelItem(item Primitive) *List {
	list.modelItem.Set(item)
	return list
}

// Returns the currently selected items or empty list for none selected.
func (list *List) SelectedItems() []int {
	return list.selectedItems.Get()
}

// Sets the currently selected items.
func (list *List) SetSelectedItems(selectedItems []int) *List {
	list.selectedItems.Set(selectedItems)
	return list
}

// Returns the status of the table: 0 = None, 1 = One Always Selected, 2 = One or None, 3 = Multiple
// 4 = Range Atleast One, 5 = Range Any or None at all.
func (list *List) SelectionMode() int {
	return list.selectionMode.Get()
}

// Sets the selection mode: 0 = None, 1 = One Always Selected, 2 = One or None, 3 = Multiple
// 4 = Range Atleast One, 5 = Range Any or None at all.
func (list *List) SetSelectionMode(mode int) *List {
	list.selectionMode.Set(mode)
	return list
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Lists inside other containers.
func (list *List) Tag() string {
	return list.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Lists inside other containers.
func (list *List) SetTag(s string) *List {
	list.tag.Set(s)
	return list
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *List) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *List) SetStatus(i int) *List {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *List) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *List) SetVisible(visible bool) *List {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *List) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *List) SetEnabled(enabled bool) *List {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *List) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *List) SetCollapsed(collapsed bool) *List {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
