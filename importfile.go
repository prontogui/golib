// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A file that represents a blob of data that can be imported from the app side
// and consumed on the server side.
type ImportFileWith struct {
	Data            []byte
	Embodiment      string
	Name            string
	ValidExtensions []string
}

// Makes a new ImportFile with specified field values.
func (w ImportFileWith) Make() *ImportFile {
	ifile := &ImportFile{}
	ifile.data.Set(w.Data)
	ifile.embodiment.Set(w.Embodiment)
	ifile.name.Set(w.Name)
	ifile.validExtensions.Set(w.ValidExtensions)
	return ifile
}

// A file that represents a blob of data that can be imported from the app side
// and consumed on the server side.
type ImportFile struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	data            BlobField
	embodiment      StringField
	imported        BooleanField
	name            StringField
	validExtensions Strings1DField
}

// Creates a new ImportFile.
func NewImportFile() *ImportFile {
	return ImportFileWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (ifile *ImportFile) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	ifile.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Data, &ifile.data},
			{key.FKey_Embodiment, &ifile.embodiment},
			{key.FKey_Imported, &ifile.imported},
			{key.FKey_Name, &ifile.name},
			{key.FKey_ValidExtensions, &ifile.validExtensions},
		}
	})
}

// Returns the blob of data for the file.
func (ifile *ImportFile) Data() []byte {
	return ifile.data.Get()
}

// Sets the blob of data for the file.
func (ifile *ImportFile) SetData(d []byte) *ImportFile {
	ifile.data.Set(d)
	return ifile
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (ifile *ImportFile) Embodiment() string {
	return ifile.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (ifile *ImportFile) SetEmbodiment(s string) *ImportFile {
	ifile.embodiment.Set(s)
	return ifile
}

// Returns true when the file has been imported by the app side and signals to the server
// side that file is ready to processs.  This field is normally only updated by the app.
func (ifile *ImportFile) Imported() bool {
	return ifile.imported.Get()
}

// Sets whether the file has been imported by the app side and signals to the server
// side that file is ready to processs.  This field is normally only updated by the app.
func (ifile *ImportFile) SetImported(b bool) *ImportFile {
	ifile.imported.Set(b)
	return ifile
}

// Returns the imported file name including its extension separated by a period.
func (ifile *ImportFile) Name() string {
	return ifile.name.Get()
}

// Sets the imported file name including its extension separated by a period.
func (ifile *ImportFile) SetName(s string) *ImportFile {
	ifile.name.Set(s)
	return ifile
}

// Returns the valid extensions for importing (non-case sensitive and period separator is omitted).
func (ifile *ImportFile) ValidExtensions() []string {
	return ifile.validExtensions.Get()
}

// Sets the valid extensions for importing (non-case sensitive and period separator is omitted).
func (ifile *ImportFile) SetValidExtensions(sa []string) *ImportFile {
	ifile.validExtensions.Set(sa)
	return ifile
}
