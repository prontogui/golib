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
	Showing    bool
	FrameItems []Primitive
}

// Creates a new Frame using the supplied field assignments.
func (w FrameWith) Make() *Frame {
	frame := &Frame{}
	frame.embodiment.Set(w.Embodiment)
	frame.showing.Set(w.Showing)
	frame.frameItems.Set(w.FrameItems)
	return frame
}

// A frame represents a complete user interface to show on the screen.  It could be
// the main user interface or a sub-screen in the app.  It includes the ability to
// layout controls in a specific manner.
type Frame struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	showing    BooleanField
	frameItems Any1DField
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
			{key.FKey_Showing, &frame.showing},
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

// Returns whether the Frame is being shown on the screen.
func (frame *Frame) Showing() bool {
	return frame.showing.Get()
}

// Sets whether the Frame is being shown on the screen.
func (frame *Frame) SetShowing(showing bool) *Frame {
	frame.showing.Set(showing)
	return frame
}
