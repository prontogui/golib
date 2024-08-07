// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TableAttachedFields(t *testing.T) {
	table := &Table{}
	table.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, table.PrimitiveBase, "Embodiment", "Headings", "Rows", "TemplateRow")
}

func Test_TableMake(t *testing.T) {
	table := TableWith{
		Embodiment:  "paginated",
		Headings:    []string{"H1", "H2"},
		Rows:        [][]Primitive{{&Command{}, &Command{}}, {&Command{}, &Command{}}},
		TemplateRow: []Primitive{&Command{}, &Command{}},
	}.Make()

	if table.Embodiment() != "paginated" {
		t.Error("'Embodiment' field was not initialized correctly")
	}

	if len(table.Headings()) != 2 {
		t.Error("'Headings' field was not initialized correctly")
	}

	if len(table.Rows()) != 2 {
		t.Error("'Rows' field was not initialized correctly")
	}

	if len(table.Rows()[0]) != 2 {
		t.Error("'Rows' field was not initialized correctly")
	}

	if len(table.Rows()[1]) != 2 {
		t.Error("'Rows' field was not initialized correctly")
	}

	if len(table.TemplateRow()) != 2 {
		t.Error("'Rows' field was not initialized correctly")
	}
}

func Test_TableFieldSettings(t *testing.T) {

	table := &Table{}

	// Embodiment field
	table.SetEmbodiment("paginated")
	if table.Embodiment() != "paginated" {
		t.Error("Unable to properly set the Embodiment field")
	}

	// Headings field

	table.SetHeadingsVA("H1", "H2")
	if len(table.Headings()) != 2 {
		t.Errorf("Headings() returned %d items.  Expecting 2 items.", len(table.Headings()))
	}

	if table.Headings()[0] != "H1" {
		t.Error("Headings()[0] not equal to 'H1'")
	}

	if table.Headings()[1] != "H2" {
		t.Error("Headings()[1] not equal to 'H2'")
	}

	// Rows field

	table.SetRows([][]Primitive{{&Command{}, &Command{}}})

	tableGet := table.Rows()

	if len(tableGet) != 1 {
		t.Errorf("ListItems() returned %d items.  Expecting 1 items.", len(tableGet))
	}

	_, ok := tableGet[0][0].(*Command)
	if !ok {
		t.Error("First group is not a Command primitive.")
	}
	_, ok = tableGet[0][1].(*Command)
	if !ok {
		t.Error("Second group is not a Command primitive.")
	}

	// TemplateRow field tests
	table.SetTemplateRow([]Primitive{&Text{}, &Command{}})

	_, ok = table.TemplateRow()[0].(*Text)

	if !ok {
		t.Error("Unable to set template item to a Text primitive")
	}

	_, ok = table.TemplateRow()[1].(*Command)

	if !ok {
		t.Error("Unable to set template item to a Command primitive")
	}
}

func Test_TableGetChildPrimitive(t *testing.T) {

	table := &Table{}

	cmdr0c0 := CommandWith{Label: "r0c0"}.Make()
	cmdr0c1 := CommandWith{Label: "r0c1"}.Make()
	cmdr1c0 := CommandWith{Label: "r1c0"}.Make()
	cmdr1c1 := CommandWith{Label: "r1c1"}.Make()
	cmdtr0 := CommandWith{Label: "a"}.Make()
	cmdtr1 := CommandWith{Label: "b"}.Make()

	table.SetRows([][]Primitive{{cmdr0c0, cmdr0c1}, {cmdr1c0, cmdr1c1}})
	table.SetTemplateRow([]Primitive{cmdtr0, cmdtr1})

	locate := func(pkey key.PKey) *Command {
		locator := key.NewPKeyLocator(pkey)
		return table.LocateNextDescendant(locator).(*Command)
	}

	if locate(key.NewPKey(0, 0, 0)).Label() != "r0c0" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 0, 0.")
	}

	if locate(key.NewPKey(0, 0, 1)).Label() != "r0c1" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 0, 1.")
	}

	if locate(key.NewPKey(0, 1, 0)).Label() != "r1c0" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 1, 0.")
	}

	if locate(key.NewPKey(0, 1, 1)).Label() != "r1c1" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 0, 1, 1.")
	}

	if locate(key.NewPKey(1, 0)).Label() != "a" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 1, 0.")
	}

	if locate(key.NewPKey(1, 1)).Label() != "b" {
		t.Fatal("LocateNextDescendant doesn't return a child for pkey 1, 1.")
	}
}
