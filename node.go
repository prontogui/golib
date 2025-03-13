// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A node is an item in a tree that represent a primitive 'node item' and contains 'sub nodes' or other nodes.
type NodeWith struct {
	Embodiment string
	NodeItem   Primitive
	SubNodes   []Primitive
	Tag        string
}

// Creates a new Group using the supplied field assignments.
func (w NodeWith) Make() *Node {
	node := &Node{}
	node.embodiment.Set(w.Embodiment)
	node.nodeItem.Set(w.NodeItem)
	node.subNodes.Set(w.SubNodes)
	node.tag.Set(w.Tag)
	return node
}

// A node is an item in a tree that represent a primitive 'node item' and contains 'sub nodes' or other nodes.
type Node struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	nodeItem   AnyField
	subNodes   Any1DField
	tag        StringField
}

// Creates a new Node and assigns items.
func NewNode(nodeItem Primitive, subNodes ...Primitive) *Node {
	return NodeWith{NodeItem: nodeItem, SubNodes: subNodes}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (node *Node) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	node.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &node.embodiment},
			{key.FKey_NodeItem, &node.nodeItem},
			{key.FKey_SubNodes, &node.subNodes},
			{key.FKey_Tag, &node.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (node *Node) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	nextIndex := locator.NextIndex()

	switch nextIndex {
	case 0:
		return node.NodeItem()
	case 1:
		return node.SubNodes()[locator.NextIndex()]
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (node *Node) Embodiment() string {
	return node.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (node *Node) SetEmbodiment(s string) *Node {
	node.embodiment.Set(s)
	return node
}

// Returns the node item.
func (node *Node) NodeItem() Primitive {
	return node.nodeItem.Get()
}

// Sets the node item.
func (node *Node) SetNodeItem(item Primitive) *Node {
	node.nodeItem.Set(item)
	return node
}

// Returns the sub nodes.
func (node *Node) SubNodes() []Primitive {
	return node.subNodes.Get()
}

// Sets the sub nodes.
func (node *Node) SetSubNodes(nodes []Primitive) *Node {
	node.subNodes.Set(nodes)
	return node
}

// Sets the sub nodes (a variadic argument list).
func (node *Node) SetSubNodesVA(nodes ...Primitive) *Node {
	node.subNodes.Set(nodes)
	return node
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Groups inside other containers.
func (node *Node) Tag() string {
	return node.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Groups inside other containers.
func (node *Node) SetTag(s string) *Node {
	node.tag.Set(s)
	return node
}
