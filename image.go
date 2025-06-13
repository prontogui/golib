// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// An image for displaying a graphic to display on the screen.
type ImageWith struct {
	// Embodiment specification
	Embodiment string

	// Binary data representing the image.  Supported formats are: JPEG, PNG, GIF, Animated GIF, WebP, Animated WebP, BMP, and WBMP
	Image []byte

	// The file path containing an image.  If specified, this takes precedence over Image field.
	// Supported file types are: JPEG, PNG, GIF, Animated GIF, WebP, Animated WebP, BMP, and WBMP
	FromFile string

	// Arbitraty tag string for the primitive
	Tag string

	// ID for this primitive. Used when referencing it from another primitive. This can be arbitrary but
	// it should ideally be a unique string.  Otherwise, when two primitives have the same ID, there is no
	// guarantee on which is referenced.
	ID string

	// Reference (ID) of another image primitive that contains the actual image data to display.  This allows
	// a way to use the same image in multiple places in a more efficient manner.
	Ref string
}

// Makes a new Image with specified field values.
func (w ImageWith) Make() (*Image, error) {
	image := &Image{}
	image.embodiment.Set(w.Embodiment)

	if len(w.FromFile) > 0 {
		if err := image.image.LoadFromFile(w.FromFile); err != nil {
			return nil, err
		}
	} else {
		image.image.Set(w.Image)
	}

	image.tag.Set(w.Tag)
	image.id.Set(w.ID)
	image.ref.Set(w.Ref)

	return image, nil
}

// An image for displaying a graphic to display on the screen.
type Image struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	id         StringField
	image      BlobField
	ref        StringField
	tag        StringField
}

// Creates a new Image from a file.
func NewImage(fromFile string) (*Image, error) {
	return ImageWith{FromFile: fromFile}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (image *Image) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	image.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &image.embodiment},
			{key.FKey_ID, &image.id},
			{key.FKey_Image, &image.image},
			{key.FKey_Ref, &image.ref},
			{key.FKey_Tag, &image.tag},
		}
	})
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (image *Image) Embodiment() string {
	return image.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (image *Image) SetEmbodiment(s string) *Image {
	image.embodiment.Set(s)
	return image
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (image *Image) ID() string {
	return image.id.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (image *Image) SetID(s string) *Image {
	image.id.Set(s)
	return image
}

// Returns the binary data for the image
func (image *Image) Image() []byte {
	return image.image.Get()
}

// Sets the binary data for the image
func (image *Image) SetImage(data []byte) *Image {
	image.image.Set(data)
	return image
}

func (image *Image) LoadFromFile(filename string) error {
	return image.image.LoadFromFile(filename)
}

func (image *Image) SaveToFile(filename string) error {
	return image.image.SaveToFile(filename)
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (image *Image) Ref() string {
	return image.ref.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (image *Image) SetRef(s string) *Image {
	image.ref.Set(s)
	return image
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (image *Image) Tag() string {
	return image.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (image *Image) SetTag(s string) *Image {
	image.tag.Set(s)
	return image
}
