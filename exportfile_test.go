// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ExportFileAttachedFields(t *testing.T) {
	ef := &ExportFile{}
	ef.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, ef.PrimitiveBase, "Data", "Embodiment", "Name", "Tag")
}

func Test_ExportFileMake(t *testing.T) {
	ef := ExportFileWith{Data: []byte{0, 1, 2}, Embodiment: "shiny", Name: "abc", Tag: "F"}.Make()

	if len(ef.Data()) != 3 {
		t.Error("could not initialize Data field")
	}

	if ef.Embodiment() != "shiny" {
		t.Error("could not initialize Embodiment field")
	}

	if ef.Name() != "abc" {
		t.Error("could not initialize Name field")
	}

	if ef.Tag() != "F" {
		t.Error("could not initialize Tag field")
	}
}

func Test_ExportFileFieldSetting(t *testing.T) {
	ef := &ExportFile{}
	ef.PrepareForUpdates(key.NewPKey(), nil)

	ef.SetData([]byte{1, 2, 3})
	if len(ef.Data()) != 3 {
		t.Error("could not set Data field")
	}

	ef.SetEmbodiment("sleek")
	if ef.Embodiment() != "sleek" {
		t.Error("could not set Embodiment field")
	}

	ef.SetName("Hello")
	if ef.Name() != "Hello" {
		t.Error("could not set Hello field")
	}

	ef.SetTag("ABC")
	if ef.Tag() != "ABC" {
		t.Error("could not set Tag field")
	}
}

func Test_ExportReset(t *testing.T) {
	ef := &ExportFile{}
	ef.PrepareForUpdates(key.NewPKey(), nil)

	ef.SetData([]byte{1, 2, 3})
	ef.SetExported(true)
	ef.Reset()

	if len(ef.Data()) != 0 {
		t.Error("data wasn't cleared")
	}

	if ef.Exported() != false {
		t.Error("exported flag wasn't set to false")
	}
}
