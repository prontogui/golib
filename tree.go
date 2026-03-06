// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

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
	Status       int
	Tag          string
}

// Makes a new Tree with specified field values.
func (w TreeWith) Make() *Tree {
	tree := &Tree{}
	tree.embodiment.Set(w.Embodiment)
	tree.modelItem.Set(w.ModelItem)
	tree.root.Set(w.Root)
	tree.selectedPath.Set(w.SelectedPath)
	tree.status.Set(w.Status)
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
	status       IntegerField
	tag          StringField
}

// Creates a new Tree and assigns the root item.
func NewTree(root Primitive) *Tree {
	return TreeWith{Root: root}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (tree *Tree) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	tree.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tree.embodiment},
			{key.FKey_ModelItem, &tree.modelItem},
			{key.FKey_Root, &tree.root},
			{key.FKey_SelectedPath, &tree.selectedPath},
			{key.FKey_Status, &tree.status},
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

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Tree) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *Tree) SetStatus(i int) *Tree {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Tree) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Tree) SetVisible(visible bool) *Tree {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Tree) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Tree) SetEnabled(enabled bool) *Tree {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Tree) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Tree) SetCollapsed(collapsed bool) *Tree {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
