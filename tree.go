// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A tree displays nodes in a hierarchical structure.  It is often used to display an outline of
// a document, a directory structure, etc.
type TreeWith struct {
	Embodiment   string
	ModelItem    Primitive
	Root         Primitive
	SelectedPath []int
	Tag          string
}

// Makes a new Tree with specified field values.
func (w TreeWith) Make() *Tree {
	tree := &Tree{}
	tree.embodiment.Set(w.Embodiment)
	tree.modelItem.Set(w.ModelItem)
	tree.root.Set(w.Root)
	tree.selectedPath.Set(w.SelectedPath)
	tree.tag.Set(w.Tag)
	return tree
}

// A tree displays nodes in a hierarchical structure.  It is often used to display an outline of
// a document, a directory structure, etc.
type Tree struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment   StringField
	modelItem    AnyField
	root         AnyField
	selectedPath Integer1DField
	tag          StringField
}

// Creates a new Tree and assigns the root item.
func NewTree(root Primitive) *Tree {
	return TreeWith{Root: root}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (tree *Tree) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tree.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tree.embodiment},
			{key.FKey_ModelItem, &tree.modelItem},
			{key.FKey_Root, &tree.root},
			{key.FKey_SelectedPath, &tree.selectedPath},
			{key.FKey_Tag, &tree.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (tree *Tree) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	nextIndex := locator.NextIndex()
	switch nextIndex {
	case 0:
		return tree.ModelItem()
	case 1:
		return tree.Root()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a string representation of this primitive.
// Implements of fmt:Stringer interface.
func (tree *Tree) String() string {
	return "<tree>"
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (tree *Tree) Embodiment() string {
	return tree.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (tree *Tree) SetEmbodiment(s string) *Tree {
	tree.embodiment.Set(s)
	return tree
}

// Returns the model item for this card.
func (tree *Tree) ModelItem() Primitive {
	return tree.modelItem.Get()
}

// Sets the model item for this card.
func (tree *Tree) SetModelItem(p Primitive) *Tree {
	tree.modelItem.Set(p)
	return tree
}

// Returns the root node for this card.
func (tree *Tree) Root() Primitive {
	return tree.root.Get()
}

// Sets the root node for this card.
func (tree *Tree) SetRoot(p Primitive) *Tree {
	tree.root.Set(p)
	return tree
}

// Returns the selected path.
func (tree *Tree) SelectedPath() []int {
	return tree.selectedPath.Get()
}

// Sets the selected path.
func (tree *Tree) SetSelectedPath(path []int) *Tree {
	tree.selectedPath.Set(path)
	return tree
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (tree *Tree) Tag() string {
	return tree.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (tree *Tree) SetTag(s string) *Tree {
	tree.tag.Set(s)
	return tree
}
