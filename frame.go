// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A frame represents a complete user interface to show on the screen.  It could be
// the main user interface or a sub-screen in the app.  It includes the ability to
// layout controls in a specific manner.
type FrameWith struct {
	Embodiment string
	FrameItems []Primitive
	Icon       Primitive
	Showing    bool
	Tag        string
	Title      string
}

// Creates a new Frame using the supplied field assignments.
func (w FrameWith) Make() *Frame {
	frame := &Frame{}
	frame.embodiment.Set(w.Embodiment)
	frame.frameItems.Set(w.FrameItems)
	frame.icon.Set(w.Icon)
	frame.showing.Set(w.Showing)
	frame.tag.Set(w.Tag)
	frame.title.Set(w.Title)
	return frame
}

// A frame represents a complete user interface to show on the screen.  It could be
// the main user interface or a sub-screen in the app.  It includes the ability to
// layout controls in a specific manner.
type Frame struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	frameItems Any1DField
	showing    BooleanField
	tag        StringField
	title      StringField
	icon       AnyField
}

// Creates a new Frame and assigns a set of items.
func NewFrame(items ...Primitive) *Frame {
	return FrameWith{FrameItems: items}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (frame *Frame) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	frame.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &frame.embodiment},
			{key.FKey_FrameItems, &frame.frameItems},
			{key.FKey_Icon, &frame.icon},
			{key.FKey_Showing, &frame.showing},
			{key.FKey_Tag, &frame.tag},
			{key.FKey_Title, &frame.title},
		}
	})
}

// A non-recursive method to locate descendants by PKey.  This is used internally by this library
// and normally should not be called by users of the library.
func (frame *Frame) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	// TODO:  generalize this code by handling inside primitive Reserved area.
	if locator.NextIndex() != 0 {
		panic("cannot locate descendent using a pkey that we assumed was valid")
	}
	return frame.FrameItems()[locator.NextIndex()]
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (frame *Frame) Embodiment() string {
	return frame.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (frame *Frame) SetEmbodiment(s string) *Frame {
	frame.embodiment.Set(s)
	return frame
}

// Returns the collection of primitives that comprise the GUI frame.
func (frame *Frame) FrameItems() []Primitive {
	return frame.frameItems.Get()
}

// Sets the collection of primitives that comprise the GUI frame.
func (frame *Frame) SetFrameItems(items []Primitive) *Frame {
	frame.frameItems.Set(items)
	return frame
}

// Sets the collection of primitives (variadic argument list) that comprise the GUI frame.
func (frame *Frame) SetFrameItemsVA(items ...Primitive) *Frame {
	frame.frameItems.Set(items)
	return frame
}

// Returns the optional icon item for this frame.
func (frame *Frame) Icon() Primitive {
	return frame.icon.Get()
}

// Sets the optional icon for this frame.
func (frame *Frame) SetIcon(p Primitive) *Frame {
	frame.icon.Set(p)
	return frame
}

// Returns whether the Frame is being shown on the screen.
func (frame *Frame) Showing() bool {
	return frame.showing.Get()
}

// Sets whether the Frame is being shown on the screen.
func (frame *Frame) SetShowing(showing bool) *Frame {
	frame.showing.Set(showing)
	return frame
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Frames inside other containers.
func (frame *Frame) Tag() string {
	return frame.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, uch as using Frames inside other containers.
func (frame *Frame) SetTag(s string) *Frame {
	frame.tag.Set(s)
	return frame
}

// Returns the title of the frame.
func (frame *Frame) Title() string {
	return frame.title.Get()
}

// Sets the title of the frame.
func (frame *Frame) SetTitle(s string) *Frame {
	frame.title.Set(s)
	return frame
}
