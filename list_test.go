// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ListAttachedFields(t *testing.T) {
	list := &List{}
	list.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, list.PrimitiveBase, "Embodiment", "ListItems", "SelectedIndex", "Tag")
}

func Test_ListMake(t *testing.T) {
	list := ListWith{
		Embodiment:    "scrolling",
		ListItems:     []Primitive{&Command{}, &Command{}},
		SelectedIndex: 1,
		Tag:           "F",
	}.Make()

	if list.Embodiment() != "scrolling" {
		t.Error("'Embodiment' field was not initialized correctly")
	}

	if len(list.ListItems()) != 2 {
		t.Error("'ListItems' field was not initialized correctly")
	}

	if list.SelectedIndex() != 1 {
		t.Error("List selection not initialized properly")
	}

	if list.Tag() != "F" {
		t.Error("Tag field was not initialized correctly")
	}
}

func Test_ListFieldSettings(t *testing.T) {

	list := &List{}

	// Embodiment field
	list.SetEmbodiment("scrolling")
	if list.Embodiment() != "scrolling" {
		t.Error("Unable to properly set the Embodiment field")
	}

	// ListItems field (as array)

	list.SetListItems([]Primitive{&Command{}, &Command{}})

	listGet := list.ListItems()

	if len(listGet) != 2 {
		t.Errorf("ListItems() returned %d items.  Expecting 2 items.", len(listGet))
	}

	_, ok1 := listGet[0].(*Command)
	if !ok1 {
		t.Error("First group is not a Command primitive.")
	}
	_, ok2 := listGet[1].(*Command)
	if !ok2 {
		t.Error("Second group is not a Command primitive.")
	}

	// ListItems field (as variadic items)

	list.SetListItemsVA(&Text{}, &Text{})

	listGet = list.ListItems()

	if len(listGet) != 2 {
		t.Errorf("ListItems() returned %d items after calling variadic setter.  Expecting 2 items.", len(listGet))
	}

	_, ok1 = listGet[0].(*Text)
	if !ok1 {
		t.Error("First item is not a Text primitive.")
	}
	_, ok2 = listGet[1].(*Text)
	if !ok2 {
		t.Error("Second item is not a Text primitive.")
	}

	// Selected field tests

	list.SetSelectedIndex(-1)
	if list.SelectedIndex() != -1 {
		t.Error("Unable to set seletion to -1")
	}

	list.SetSelectedIndex(0)
	if list.SelectedIndex() != 0 {
		t.Error("Unable to set seletion to 0")
	}

	list.SetSelectedIndex(1)
	if list.SelectedIndex() != 1 {
		t.Error("Unable to set seletion to 1")
	}

	// Tag field test
	list.SetTag("ABC")
	if list.Tag() != "ABC" {
		t.Error("unable to set Tag field")
	}
}

func Test_ListGetChildPrimitive(t *testing.T) {

	list := &List{}

	cmd1 := CommandWith{Label: "a"}.Make()
	cmd2 := CommandWith{Label: "b"}.Make()

	list.SetListItemsVA(cmd1, cmd2)

	locate := func(pkey key.PKey) *Command {
		locator := key.NewPKeyLocator(pkey)
		return list.LocateNextDescendant(locator).(*Command)
	}

	if locate(key.NewPKey(0, 0)).Label() != "a" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 0.")
	}

	if locate(key.NewPKey(0, 1)).Label() != "b" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 1.")
	}
}
func Test_ListSelectedItem(t *testing.T) {
	list := &List{}

	// Test when no item is selected
	list.SetSelectedIndex(-1)
	if list.SelectedItem() != nil {
		t.Error("Expected nil when no item is selected")
	}

	// Test when selected index is out of range
	list.SetListItems([]Primitive{&Command{}, &Command{}})
	list.SetSelectedIndex(2)
	if list.SelectedItem() != nil {
		t.Error("Expected nil when selected index is out of range")
	}

	// Test when selected index is within range
	list.SetSelectedIndex(1)
	if list.SelectedItem() == nil {
		t.Error("Expected a valid item when selected index is within range")
	}

	// Verify the selected item is correct
	cmd := &Command{}
	list.SetListItems([]Primitive{&Command{}, cmd})
	list.SetSelectedIndex(1)
	if list.SelectedItem() != cmd {
		t.Error("SelectedItem did not return the correct item")
	}
}
