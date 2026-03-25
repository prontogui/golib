// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"errors"
	"slices"

	"github.com/prontogui/golib/key"
)

// A table displays an array of primitives in a grid of rows and columns.
type TableWith struct {
	Embodiment    string
	HeaderRow     []Primitive
	ModelRow      []Primitive
	Rows          [][]Primitive
	SelectedRows  []int
	SelectionMode int
	Status        int
	Tag           string
}

// Creates a new Table using the supplied field assignments.
func (w TableWith) Make() *Table {

	table := &Table{}

	table.embodiment.Set(w.Embodiment)
	table.headerRow.Set(w.HeaderRow)
	table.modelRow.Set(w.ModelRow)
	table.rows.Set(w.Rows)
	table.selectedRows.Set(w.SelectedRows)
	table.selectionMode.Set(w.SelectionMode)
	table.status.Set(w.Status)
	table.tag.Set(w.Tag)
	return table
}

// A table displays an array of primitives in a grid of rows and columns.
type Table struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment       StringField
	headerRow        Any1DField
	modelRow         Any1DField
	rows             Any2DField
	selectedRows     Integer1DField
	selectionChanged EventField
	selectionMode    IntegerField
	status           IntegerField
	tag              StringField
}

// Creates a new Table primitive.
func NewTable() *Table {
	return TableWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (table *Table) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	table.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &table.embodiment},
			{key.FKey_HeaderRow, &table.headerRow},
			{key.FKey_ModelRow, &table.modelRow},
			{key.FKey_Rows, &table.rows},
			{key.FKey_SelectedRows, &table.selectedRows},
			{key.FKey_SelectionChanged, &table.selectionChanged},
			{key.FKey_SelectionMode, &table.selectionMode},
			{key.FKey_Status, &table.status},
			{key.FKey_Tag, &table.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (table *Table) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return table.HeaderRow()[locator.NextIndex()]
	case 1:
		return table.ModelRow()[locator.NextIndex()]
	case 2:
		// TODO:  Optimization - add a row/col accessor to Any2D field so we don't return all the contents just
		// to index a single item here.  Same could be done for Any1D.
		row := locator.NextIndex()
		col := locator.NextIndex()
		return table.Rows()[row][col]
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (table *Table) Embodiment() string {
	return table.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (table *Table) SetEmbodiment(s string) *Table {
	table.embodiment.Set(s)
	return table
}

// Returns the header row that describes each column in the table.
func (table *Table) HeaderRow() []Primitive {
	return table.headerRow.Get()
}

// Sets the headings to use for each column in the table.
func (table *Table) SetHeaderRow(items []Primitive) *Table {
	table.headerRow.Set(items)
	return table
}

// Sets the headings (as variadic arguments) to use for each column in the table.
func (table *Table) SetHeaderRowVA(items ...Primitive) *Table {
	table.headerRow.Set(items)
	return table
}

// Returns the model row.
func (table *Table) ModelRow() []Primitive {
	return table.modelRow.Get()
}

// Sets the model row.
func (table *Table) SetModelRow(items []Primitive) *Table {
	table.modelRow.Set(items)
	return table
}

// Returns the dynamically populated 2D (rows, cols) collection of primitives that appear in the table.
func (table *Table) Rows() [][]Primitive {
	return table.rows.Get()
}

// Sets the dynamically populated 2D (rows, cols) collection of primitives that appear in the table.
func (table *Table) SetRows(items [][]Primitive) *Table {
	table.rows.Set(items)
	return table
}

// Returns the selected rows.
func (table *Table) SelectedRows() []int {
	return table.selectedRows.Get()
}

// Sets the selected rows.
func (table *Table) SetSelectedRows(selected []int) *Table {
	table.selectedRows.Set(selected)
	return table
}

// Returns true if the selection was changed.
func (cmd *Table) SelectionChanged() bool {
	return cmd.selectionChanged.Issued()
}

// Returns the status of the table: 0 = None, 1 = One Always Selected, 2 = One or None, 3 = Multiple
// 4 = Range Atleast One, 5 = Range Any or None at all.
func (table *Table) SelectionMode() int {
	return table.selectionMode.Get()
}

// Sets the selection mode: 0 = None, 1 = One Always Selected, 2 = One or None, 3 = Multiple
// 4 = Range Atleast One, 5 = Range Any or None at all.
func (table *Table) SetSelectionMode(mode int) *Table {
	table.selectionMode.Set(mode)
	return table
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tables inside other containers.
func (table *Table) Tag() string {
	return table.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Tables inside other containers.
func (table *Table) SetTag(s string) *Table {
	table.tag.Set(s)
	return table
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *Table) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *Table) SetStatus(i int) *Table {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Table) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Table) SetVisible(visible bool) *Table {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Table) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Table) SetEnabled(enabled bool) *Table {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Table) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Table) SetCollapsed(collapsed bool) *Table {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}

// CONVENIENCE FUNCTIONS

// Inserts a new row in this table before the index specified.  If index is -1 or extends beyond the number
// of rows in the table then row is appended at the end of the table.
// The row must match the dimension and cell types of the template row
func (table *Table) InsertRow(index int, row []Primitive) {

	originalRows := table.rows.Get()

	if index < 0 || index > len(originalRows) {
		table.rows.Set(append(originalRows, row))
		return
	}

	table.rows.Set(slices.Insert(originalRows, index, row))
}

// Deletes a row in this table at the given index.  An error is returned if the index is out of range.
func (table *Table) DeleteRow(index int) error {

	originalRows := table.rows.Get()

	if index < 0 || index >= len(originalRows) {
		return errors.New("index out of range")
	}

	table.rows.Set(slices.Delete(originalRows, index, index+1))

	return nil
}

// Convenience function that returns the number of rows.
func (table *Table) RowCount() int {
	return table.rows.Length()
}

// Deletes all rows from the table.
func (table *Table) DeleteAllRows() {
	table.rows.Set([][]Primitive{})
}

// Generates a list of headings from each primitive in the HeaderRow. Although these are typically
// Text primitives, the scan relies on the primitive's string representation to allow for
// more general use.
func (table *Table) GetHeadings() []string {
	headerRow := table.HeaderRow()
	headings := make([]string, len(headerRow))
	for i, p := range headerRow {
		headings[i] = p.String()
	}
	return headings
}

// Constructs the HeaderRow from Text primitives, each holding content from the corresponding
// string provided in the headings argument.
func (table *Table) MakeHeadings(headings []string) *Table {
	primitives := make([]Primitive, len(headings))
	for i, h := range headings {
		primitives[i] = TextWith{Content: h}.Make()
	}
	table.SetHeaderRow(primitives)
	return table
}

// Constructs the HeaderRow from Text primitives, each holding content from the corresponding
// string provided in the variable argument list.
func (table *Table) MakeHeadingsVA(headings ...string) *Table {
	return table.MakeHeadings(headings)
}
