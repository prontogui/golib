// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A list is a collection of primitives that have a sequential-like relationship
// and might be dynamic in quantity or kind.
type ListWith struct {
	Embodiment    string
	ListItems     []Primitive
	ModelItem     Primitive
	SelectedIndex int
	Tag           string
}

// Creates a new List using the supplied field assignments.
func (w ListWith) Make() *List {
	list := &List{}
	list.embodiment.Set(w.Embodiment)
	list.listItems.Set(w.ListItems)
	list.modelItem.Set(w.ModelItem)
	list.selectedIndex.Set(w.SelectedIndex)
	list.tag.Set(w.Tag)
	return list
}

// A list is a collection of primitives that have a sequential-like relationship
// and might be dynamic in quantity or kind.
type List struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment    StringField
	listItems     Any1DField
	modelItem     AnyField
	selectedIndex IntegerField
	tag           StringField
}

// Creates a new List and assigns items.
func NewList(items ...Primitive) *List {
	return ListWith{ListItems: items}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (list *List) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	list.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &list.embodiment},
			{key.FKey_ListItems, &list.listItems},
			{key.FKey_SelectedIndex, &list.selectedIndex},
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

// Returns the model item.
func (list *List) ModelItem() Primitive {
	return list.modelItem.Get()
}

// Sets the ttems to show in the list.
func (list *List) SetModelItem(item Primitive) *List {
	list.modelItem.Set(item)
	return list
}

// Returns the currently selected item or -1 for none selected.
func (list *List) SelectedIndex() int {
	return list.selectedIndex.Get()
}

// Sets the currently selected item or -1 for none selected.
func (list *List) SetSelectedIndex(selectedIndex int) *List {
	list.selectedIndex.Set(selectedIndex)
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

// SelectedItem returns the currently selected item from the list.
// If the selected index is within the valid range of list items, it returns the item at the selected index.
// If the selected index is out of range, it returns nil.
func (list *List) SelectedItem() Primitive {
	selectedIndex := list.SelectedIndex()
	if selectedIndex >= 0 && selectedIndex < len(list.ListItems()) {
		return list.ListItems()[selectedIndex]
	}
	return nil
}
