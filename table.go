// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type TableWith struct {
	Embodiment  string
	Headings    []string
	Rows        [][]Primitive
	TemplateRow []Primitive
}

func (w TableWith) Make() *Table {
	table := &Table{}
	table.SetEmbodiment(w.Embodiment)
	table.SetHeadings(w.Headings)
	table.SetRows(w.Rows)
	table.SetTemplateRow(w.TemplateRow)
	return table
}

type Table struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment  StringField
	headings    Strings1DField
	rows        Any2DField
	templateRow Any1DField
}

func (table *Table) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	table.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &table.embodiment},
			{key.FKey_Headings, &table.headings},
			{key.FKey_Rows, &table.rows},
			{key.FKey_TemplateRow, &table.templateRow},
		}
	})
}

// TODO:  generalize this code by handling inside primitive Reserved area.
func (table *Table) LocateNextDescendant(locator *key.PKeyLocator) Primitive {

	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		// TODO:  Optimization - add a row/col accessor to Any2D field so we don't return all the contents just
		// to index a single item here.  Same could be done for Any1D.
		row := locator.NextIndex()
		col := locator.NextIndex()
		return table.Rows()[row][col]
	case 1:
		return table.TemplateRow()[locator.NextIndex()]
	default:
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
}

func (table *Table) Embodiment() string {
	return table.embodiment.Get()
}

func (table *Table) SetEmbodiment(s string) {
	table.embodiment.Set(s)
}

func (table *Table) Headings() []string {
	return table.headings.Get()
}

func (table *Table) SetHeadings(s []string) {
	table.headings.Set(s)
}

func (table *Table) SetHeadingsVA(items ...string) {
	table.headings.Set(items)
}

func (table *Table) TemplateRow() []Primitive {
	return table.templateRow.Get()
}

func (table *Table) SetTemplateRow(items []Primitive) {
	table.templateRow.Set(items)
}

func (table *Table) Rows() [][]Primitive {
	return table.rows.Get()
}

func (table *Table) SetRows(items [][]Primitive) {
	table.rows.Set(items)
}
