// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

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
	Status          int
	Tag             string
	ValidExtensions []string
}

// Makes a new ImportFile with specified field values.
func (w ImportFileWith) Make() *ImportFile {
	ifile := &ImportFile{}
	ifile.data.Set(w.Data)
	ifile.embodiment.Set(w.Embodiment)
	ifile.name.Set(w.Name)
	ifile.status.Set(w.Status)
	ifile.tag.Set(w.Tag)
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
	status          IntegerField
	tag             StringField
	validExtensions String1DField
}

// Creates a new ImportFile.
func NewImportFile() *ImportFile {
	return ImportFileWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (ifile *ImportFile) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	ifile.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Data, &ifile.data},
			{key.FKey_Embodiment, &ifile.embodiment},
			{key.FKey_Imported, &ifile.imported},
			{key.FKey_Name, &ifile.name},
			{key.FKey_Status, &ifile.status},
			{key.FKey_Tag, &ifile.tag},
			{key.FKey_ValidExtensions, &ifile.validExtensions},
		}
	})
}

// Returns the blob of data for the file.  Note:  this data could be empty and
// yet represent a valid imported, albeit empty, file.  Therefore, it is important to
// look at Imported() field to know whether data has been imported.  Conversely,
// if the Imported() function returns false then this will return an empty array.
func (ifile *ImportFile) Data() []byte {
	return ifile.data.Get()
}

// Sets the blob of data for the file and sets imported flag to true.
func (ifile *ImportFile) ImportData(d []byte) *ImportFile {
	ifile.data.Set(d)
	ifile.imported.Set(true)
	return ifile
}

// Clears the imported data and the imported flag.
func (ifile *ImportFile) Reset() {
	ifile.data.Set([]byte{})
	ifile.imported.Set(false)
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

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (ifile *ImportFile) Tag() string {
	return ifile.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (ifile *ImportFile) SetTag(s string) *ImportFile {
	ifile.tag.Set(s)
	return ifile
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *ImportFile) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *ImportFile) SetStatus(i int) *ImportFile {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *ImportFile) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *ImportFile) SetVisible(visible bool) *ImportFile {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *ImportFile) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *ImportFile) SetEnabled(enabled bool) *ImportFile {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *ImportFile) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *ImportFile) SetCollapsed(collapsed bool) *ImportFile {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
