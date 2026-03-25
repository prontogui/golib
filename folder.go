// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A folder represents a container in a hierarchical list structure. It contains a label item
// (any primitive) and a level indicating its depth in the hierarchy. The Expanded field controls
// whether items logically "inside" this folder are shown.
type FolderWith struct {
	Embodiment string
	Expanded   bool
	LabelItem  Primitive
	Level      int
	Status     int
	Tag        string
}

// Makes a new Folder with specified field values.
func (w FolderWith) Make() *Folder {
	folder := &Folder{}
	folder.embodiment.Set(w.Embodiment)
	folder.expanded.Set(w.Expanded)
	folder.labelItem.Set(w.LabelItem)
	folder.level.Set(w.Level)
	folder.status.Set(w.Status)
	folder.tag.Set(w.Tag)
	return folder
}

// A folder represents a container in a hierarchical list structure. It contains a label item
// (any primitive) and a level indicating its depth in the hierarchy. The Expanded field controls
// whether items logically "inside" this folder are shown.
type Folder struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	expanded   BooleanField
	labelItem  AnyField
	level      IntegerField
	status     IntegerField
	tag        StringField
}

// Creates a new Folder at a specified level and assigns the label item.
func NewFolder(labelItem Primitive, level int) *Folder {
	return FolderWith{LabelItem: labelItem, Level: level}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (folder *Folder) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	folder.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &folder.embodiment},
			{key.FKey_Expanded, &folder.expanded},
			{key.FKey_LabelItem, &folder.labelItem},
			{key.FKey_Level, &folder.level},
			{key.FKey_Status, &folder.status},
			{key.FKey_Tag, &folder.tag},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (folder *Folder) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	nextIndex := locator.NextIndex()

	// Fields are handled in alphabetical order
	switch nextIndex {
	case 0:
		return folder.LabelItem()
	}

	panic("cannot locate descendent using a pkey that we assumed was valid")
}

// Returns a string representation of this primitive.
// Implements of fmt:Stringer interface.
func (folder *Folder) String() string {
	return "<folder>"
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (folder *Folder) Embodiment() string {
	return folder.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (folder *Folder) SetEmbodiment(s string) *Folder {
	folder.embodiment.Set(s)
	return folder
}

// Returns true if the folder is expanded.
func (folder *Folder) Expanded() bool {
	return folder.expanded.Get()
}

// Sets whether the folder is expanded.
func (folder *Folder) SetExpanded(b bool) *Folder {
	folder.expanded.Set(b)
	return folder
}

// Returns the label item for this folder.
func (folder *Folder) LabelItem() Primitive {
	return folder.labelItem.Get()
}

// Sets the label item for this folder.
func (folder *Folder) SetLabelItem(p Primitive) *Folder {
	folder.labelItem.Set(p)
	return folder
}

// Returns the hierarchical level of this folder.
func (folder *Folder) Level() int {
	return folder.level.Get()
}

// Sets the hierarchical level of this folder.
func (folder *Folder) SetLevel(level int) *Folder {
	folder.level.Set(level)
	return folder
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Folders inside other containers.
func (folder *Folder) Tag() string {
	return folder.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Folders inside other containers.
func (folder *Folder) SetTag(s string) *Folder {
	folder.tag.Set(s)
	return folder
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *Folder) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *Folder) SetStatus(i int) *Folder {
	p.status.Set(i)
	return p
}

// Returns the visibility of the folder.  This is derived from the Status field.
func (p *Folder) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Folder) SetVisible(visible bool) *Folder {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Folder) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Folder) SetEnabled(enabled bool) *Folder {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Folder) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Folder) SetCollapsed(collapsed bool) *Folder {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
