// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"fmt"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TableAttachedFields(t *testing.T) {
	table := &Table{}
	table.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, table.PrimitiveBase, "Embodiment", "Headings", "Rows", "Status", "Tag")
}

func Test_TableMake(t *testing.T) {
	table := TableWith{
		Embodiment: "paginated",
		Headings:   []string{"H1", "H2"},
		Rows:       [][]Primitive{{&Command{}, &Command{}}, {&Command{}, &Command{}}},
		Status:     2,
		Tag:        "F",
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

	if table.Status() != 2 {
		t.Error("'Status' field was not initialized correctly")
	}

	if table.Tag() != "F" {
		t.Error("'Tag' field was not initialized correctly")
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

	// Status field
	table.SetStatus(2)
	if table.Status() != 2 {
		t.Error("Unable to properly set the Status field")
	}

	// Tag field
	table.SetTag("ABC")
	if table.Tag() != "ABC" {
		t.Error("unable to set the Tag field.")
	}
}

func Test_TableGetChildPrimitive(t *testing.T) {

	table := &Table{}

	cmdr0c0 := CommandWith{Label: "r0c0"}.Make()
	cmdr0c1 := CommandWith{Label: "r0c1"}.Make()
	cmdr1c0 := CommandWith{Label: "r1c0"}.Make()
	cmdr1c1 := CommandWith{Label: "r1c1"}.Make()

	table.SetRows([][]Primitive{{cmdr0c0, cmdr0c1}, {cmdr1c0, cmdr1c1}})

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
}

func _prepareTableForInsert() *Table {
	table := &Table{}

	r0c0 := CommandWith{Label: "r0c0"}.Make()
	r0c1 := TextWith{Content: "r0c1"}.Make()
	r1c0 := CommandWith{Label: "r1c0"}.Make()
	r1c1 := TextWith{Content: "r1c1"}.Make()
	r2c0 := CommandWith{Label: "r2c0"}.Make()
	r2c1 := TextWith{Content: "r2c1"}.Make()

	table.SetRows([][]Primitive{{r0c0, r0c1}, {r1c0, r1c1}, {r2c0, r2c1}})

	return table
}

func _prepareNewRowForInsert() []Primitive {
	newc0 := CommandWith{Label: "newc0"}.Make()
	newc1 := TextWith{Content: "newc1"}.Make()
	return []Primitive{newc0, newc1}
}

func _verifyRowsAfterInsertion(t *testing.T, table *Table, originalRows [3]int, newRow int) {

	numRows := len(table.Rows())
	if numRows != 4 {
		t.Fatalf("number of rows after insertion is:  %d. Expecting 4.", numRows)
	}

	testfunc := func(row int, prefix string) {
		cmd := table.Rows()[row][0].(*Command)
		text := table.Rows()[row][1].(*Text)

		cmdLabel := fmt.Sprintf("%sc0", prefix)
		textContent := fmt.Sprintf("%sc1", prefix)

		if cmd.Label() != cmdLabel || text.Content() != textContent {
			t.Errorf("row %d does not have the correct information after insertion", row)
		}
	}

	testfunc(originalRows[0], "r0")
	testfunc(originalRows[1], "r1")
	testfunc(originalRows[2], "r2")
	testfunc(newRow, "new")
}

func Test_TableInsertRow0(t *testing.T) {

	table := _prepareTableForInsert()
	table.InsertRow(0, _prepareNewRowForInsert())

	_verifyRowsAfterInsertion(t, table, [3]int{1, 2, 3}, 0)
}

func Test_TableInsertRow1(t *testing.T) {

	table := _prepareTableForInsert()
	table.InsertRow(1, _prepareNewRowForInsert())

	_verifyRowsAfterInsertion(t, table, [3]int{0, 2, 3}, 1)
}

func Test_TableInsertRow2(t *testing.T) {

	table := _prepareTableForInsert()
	table.InsertRow(-1, _prepareNewRowForInsert())

	_verifyRowsAfterInsertion(t, table, [3]int{0, 1, 2}, 3)
}

func Test_TableInsertRow3(t *testing.T) {

	table := _prepareTableForInsert()
	table.InsertRow(3, _prepareNewRowForInsert())

	_verifyRowsAfterInsertion(t, table, [3]int{0, 1, 2}, 3)
}

func _verifyRowsAfterDeletion(t *testing.T, table *Table, rowsLeft [2]int) {

	numRows := len(table.Rows())
	if numRows != 2 {
		t.Fatalf("number of rows after deletion is:  %d. Expecting 2.", numRows)
	}

	testfunc := func(row int, oldRow int) {
		cmd := table.Rows()[row][0].(*Command)
		text := table.Rows()[row][1].(*Text)

		cmdLabel := fmt.Sprintf("r%dc0", oldRow)
		textContent := fmt.Sprintf("r%dc1", oldRow)

		if cmd.Label() != cmdLabel || text.Content() != textContent {
			t.Errorf("row %d does not have the correct information after insertion", row)
		}
	}

	testfunc(0, rowsLeft[0])
	testfunc(1, rowsLeft[1])
}

func Test_TableDeleteRow0(t *testing.T) {

	table := _prepareTableForInsert()
	table.DeleteRow(0)

	_verifyRowsAfterDeletion(t, table, [2]int{1, 2})
}

func Test_TableDeleteRow1(t *testing.T) {

	table := _prepareTableForInsert()
	table.DeleteRow(1)

	_verifyRowsAfterDeletion(t, table, [2]int{0, 2})
}

func Test_TableDeleteRow2(t *testing.T) {

	table := _prepareTableForInsert()
	table.DeleteRow(2)

	_verifyRowsAfterDeletion(t, table, [2]int{0, 1})
}

func Test_TableDeleteRow3(t *testing.T) {

	table := _prepareTableForInsert()
	err := table.DeleteRow(3)

	if err == nil {
		t.Fatal("no error returned")
	}
	if err.Error() != "index out of range" {
		t.Fatal("unexpected error returned")
	}
}

func Test_TableDeleteRow4(t *testing.T) {

	table := _prepareTableForInsert()
	err := table.DeleteRow(-1)

	if err == nil {
		t.Fatal("no error returned")
	}
	if err.Error() != "index out of range" {
		t.Fatal("unexpected error returned")
	}
}

func Test_TableDeleteAllRows(t *testing.T) {
	table := _prepareTableForInsert()

	table.DeleteAllRows()

	if len(table.Rows()) != 0 {
		t.Fatalf("number of rows after deletion is: %d. Expecting 0.", len(table.Rows()))
	}
}
