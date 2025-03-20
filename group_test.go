// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_GroupAttachedFields(t *testing.T) {
	grp := &Group{}
	grp.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, grp.PrimitiveBase, "Embodiment", "GroupItems", "Tag")
}

func Test_GroupMake(t *testing.T) {
	grp := GroupWith{
		Embodiment: "row",
		GroupItems: []Primitive{&Command{}, &Command{}},
		Tag:        "F",
	}.Make()

	if grp.Embodiment() != "row" {
		t.Error("Could not initialize Embodiment field.")
	}

	if len(grp.GroupItems()) != 2 {
		t.Error("'GroupItems' field was not initialized correctly")
	}

	if grp.Tag() != "F" {
		t.Error("'Tag' field was not initialized correctly")
	}
}

func Test_GroupFieldSettings(t *testing.T) {

	grp := &Group{}

	grp.SetEmbodiment("column")
	if grp.Embodiment() != "column" {
		t.Error("Could not set Embodiment field.")
	}

	grp.SetGroupItems([]Primitive{&Command{}, &Command{}})

	grpGet := grp.GroupItems()

	if len(grpGet) != 2 {
		t.Errorf("GroupItems() returned %d items.  Expecting 2 items.", len(grpGet))
	}

	_, ok1 := grpGet[0].(*Command)
	if !ok1 {
		t.Error("First group is not a Command primitive.")
	}
	_, ok2 := grpGet[1].(*Command)
	if !ok2 {
		t.Error("Second group is not a Command primitive.")
	}

	grp.SetGroupItemsVA(&Text{}, &Text{})

	grpGet = grp.GroupItems()

	if len(grpGet) != 2 {
		t.Errorf("GroupItems() returned %d items after calling variadic setter.  Expecting 2 items.", len(grpGet))
	}

	_, ok1 = grpGet[0].(*Text)
	if !ok1 {
		t.Error("First group is not a Text primitive.")
	}
	_, ok2 = grpGet[1].(*Text)
	if !ok2 {
		t.Error("Second group is not a Text primitive.")
	}

	grp.SetTag("ABC")
	if grp.Tag() != "ABC" {
		t.Error("Could not set Tag field.")
	}
}

func Test_GroupLocateChildPrimitive(t *testing.T) {

	cmd1 := CommandWith{Label: "a"}.Make()
	cmd2 := CommandWith{Label: "b"}.Make()

	grp := GroupWith{GroupItems: []Primitive{cmd1, cmd2}}.Make()

	locate := func(pkey key.PKey) *Command {
		locator := key.NewPKeyLocator(pkey)
		return grp.LocateNextDescendant(locator).(*Command)
	}

	if locate(key.NewPKey(0, 0)).Label() != "a" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0.")
	}

	if locate(key.NewPKey(0, 1)).Label() != "b" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 1.")
	}
}
