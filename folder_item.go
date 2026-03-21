// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A folder item represents an item in a hierarchical list structure. It contains an item
// (any primitive) and a level indicating its depth in the hierarchy.
type FolderItemWith struct {
	Embodiment string
	Item       Primitive
	Level      int
	Status     int
	Tag        string
}

// Makes a new FolderItem with specified field values.
func (w FolderItemWith) Make() *FolderItem {
	folderItem := &FolderItem{}
	folderItem.embodiment.Set(w.Embodiment)
	folderItem.item.Set(w.Item)
	folderItem.level.Set(w.Level)
	folderItem.status.Set(w.Status)
	folderItem.tag.Set(w.Tag)
	return folderItem
}

// A folder item represents an item in a hierarchical list structure. It contains an item
// (any primitive) and a level indicating its depth in the hierarchy.
type FolderItem struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	item       AnyField
	level      IntegerField
	status     IntegerField
	tag        StringField
}

// Creates a new FolderItem at a specified level and assigns the item.
func NewFolderItem(item Primitive, level int) *FolderItem {
	return FolderItemWith{Item: item, Level: level}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (folderItem *FolderItem) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	folderItem.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &folderItem.embodiment},
			{key.FKey_Item, &folderItem.item},
			{key.FKey_Level, &folderItem.level},
			{key.FKey_Status, &folderItem.status},
			{key.FKey_Tag, &folderItem.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (folderItem *FolderItem) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return folderItem.Item()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a string representation of this primitive.
// Implements of fmt:Stringer interface.
func (folderItem *FolderItem) String() string {
	return "<folderitem>"
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (folderItem *FolderItem) Embodiment() string {
	return folderItem.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (folderItem *FolderItem) SetEmbodiment(s string) *FolderItem {
	folderItem.embodiment.Set(s)
	return folderItem
}

// Returns the item for this folder item.
func (folderItem *FolderItem) Item() Primitive {
	return folderItem.item.Get()
}

// Sets the item for this folder item.
func (folderItem *FolderItem) SetItem(p Primitive) *FolderItem {
	folderItem.item.Set(p)
	return folderItem
}

// Returns the hierarchical level of this folder item.
func (folderItem *FolderItem) Level() int {
	return folderItem.level.Get()
}

// Sets the hierarchical level of this folder item.
func (folderItem *FolderItem) SetLevel(level int) *FolderItem {
	folderItem.level.Set(level)
	return folderItem
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using FolderItems inside other containers.
func (folderItem *FolderItem) Tag() string {
	return folderItem.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using FolderItems inside other containers.
func (folderItem *FolderItem) SetTag(s string) *FolderItem {
	folderItem.tag.Set(s)
	return folderItem
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *FolderItem) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hiddend and disabled, 3 = collapsed and disabled.
func (p *FolderItem) SetStatus(i int) *FolderItem {
	p.status.Set(i)
	return p
}

// Returns the visibility of the folder item.  This is derived from the Status field.
func (p *FolderItem) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *FolderItem) SetVisible(visible bool) *FolderItem {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *FolderItem) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *FolderItem) SetEnabled(enabled bool) *FolderItem {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *FolderItem) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *FolderItem) SetCollapsed(collapsed bool) *FolderItem {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
