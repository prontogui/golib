// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A file represents a blob of data that can be exported from the server side and
// stored to a file on the app side.  The perspective of "export" is centered around
// the server software.  This seems to be a little clearer than using Download/Upload
// terminology.
type ExportFileWith struct {
	Data       []byte
	Embodiment string
	Name       string
	Status     int
	Tag        string
}

// Makes a new Command with specified field values.
func (w ExportFileWith) Make() *ExportFile {
	ef := &ExportFile{}
	ef.data.Set(w.Data)
	ef.embodiment.Set(w.Embodiment)
	ef.name.Set(w.Name)
	ef.status.Set(w.Status)
	ef.tag.Set(w.Tag)
	return ef
}

// A file represents a blob of data that can be exported from the server side and
// stored to a file on the app side.  The perspective of "export" is centered around
// the server software.  This seems to be a little clearer than using Download/Upload
// terminology.
type ExportFile struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	data       BlobField
	embodiment StringField
	exported   BooleanField
	name       StringField
	status     IntegerField
	tag        StringField
}

// Creates a new ExportFile.
func NewExportFile() *ExportFile {
	return ExportFileWith{}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (ef *ExportFile) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	ef.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Data, &ef.data},
			{key.FKey_Embodiment, &ef.embodiment},
			{key.FKey_Exported, &ef.exported},
			{key.FKey_Name, &ef.name},
			{key.FKey_Status, &ef.status},
			{key.FKey_Tag, &ef.tag},
		}
	})
}

// Returns the blob of data representing the binary contents of the file.  Note:  this
// data could be empty and yet represent a valid, albeit empty, file for export.
func (ef *ExportFile) Data() []byte {
	return ef.data.Get()
}

// Sets the blob of data representing the binary contents of the file to export.
func (ef *ExportFile) SetData(d []byte) *ExportFile {
	ef.data.Set(d)
	return ef
}

// Clears the exported data and the exported flag.
func (ef *ExportFile) Reset() {
	ef.data.Set([]byte{})
	ef.exported.Set(false)
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (ef *ExportFile) Embodiment() string {
	return ef.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (ef *ExportFile) SetEmbodiment(s string) *ExportFile {
	ef.embodiment.Set(s)
	return ef
}

// Returns true when the file has been exported (stored to a file) by the app.
// This field is normally only updated by the app.
func (ef *ExportFile) Exported() bool {
	return ef.exported.Get()
}

// Sets whether or not the file has been exported (stored to a file) by the app.
// This field is normally only updated by the app.
func (ef *ExportFile) SetExported(b bool) *ExportFile {
	ef.exported.Set(b)
	return ef
}

// Returns the suggested file name (including its extension separated by a period) to save the file as.
func (ef *ExportFile) Name() string {
	return ef.name.Get()
}

// Sets the suggested file name (including its extension separated by a period) to save the file as.
func (ef *ExportFile) SetName(s string) *ExportFile {
	ef.name.Set(s)
	return ef
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using ExportFiles as Table cells.
func (ef *ExportFile) Tag() string {
	return ef.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using ExportFiles as Table cells.
func (ef *ExportFile) SetTag(s string) *ExportFile {
	ef.tag.Set(s)
	return ef
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *ExportFile) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *ExportFile) SetStatus(i int) *ExportFile {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *ExportFile) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *ExportFile) SetVisible(visible bool) *ExportFile {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *ExportFile) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *ExportFile) SetEnabled(enabled bool) *ExportFile {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *ExportFile) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *ExportFile) SetCollapsed(collapsed bool) *ExportFile {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
