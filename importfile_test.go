// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ImportFileAttachedFields(t *testing.T) {
	ef := &ImportFile{}
	ef.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, ef.PrimitiveBase, "Data", "Embodiment", "Name", "Tag", "ValidExtensions")
}

func Test_ImportFileMake(t *testing.T) {
	impf := ImportFileWith{
		Data:            []byte{0, 1, 2},
		Embodiment:      "shiny",
		Name:            "abc",
		Tag:             "F",
		ValidExtensions: []string{"TXT", "CSV"},
	}.Make()

	if len(impf.Data()) != 3 {
		t.Error("could not initialize Data field")
	}

	if impf.Embodiment() != "shiny" {
		t.Error("could not initialize Embodiment field")
	}

	if impf.Name() != "abc" {
		t.Error("could not initialize Name field")
	}

	if impf.Tag() != "F" {
		t.Error("could not initialize Tag field")
	}

	if len(impf.ValidExtensions()) != 2 || impf.ValidExtensions()[0] != "TXT" || impf.ValidExtensions()[1] != "CSV" {
		t.Error("could not initialize ValidExtensions field")
	}
}

func Test_ImporttFileFieldSetting(t *testing.T) {
	impf := &ImportFile{}
	impf.PrepareForUpdates(key.NewPKey(), nil)

	impf.SetData([]byte{1, 2, 3})
	if len(impf.Data()) != 3 {
		t.Error("could not set Data field")
	}

	impf.SetEmbodiment("sleek")
	if impf.Embodiment() != "sleek" {
		t.Error("could not set Embodiment field")
	}

	impf.SetName("Hello")
	if impf.Name() != "Hello" {
		t.Error("could not set Hello field")
	}

	impf.SetTag("ABC")
	if impf.Tag() != "ABC" {
		t.Error("could not set Tag field")
	}

	impf.SetValidExtensions([]string{"TXT", "CSV"})
	if len(impf.ValidExtensions()) != 2 || impf.ValidExtensions()[0] != "TXT" || impf.ValidExtensions()[1] != "CSV" {
		t.Error("could not initialize ValidExtensions field")
	}
}
