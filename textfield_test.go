// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TextFieldAttach(t *testing.T) {
	txt := &TextField{}
	txt.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, txt.PrimitiveBase, "Embodiment", "Tag", "TextEntry")
}

func Test_TextFieldMake(t *testing.T) {
	txt := TextFieldWith{
		Embodiment: "block",
		Tag:        "F",
		TextEntry:  "This is a piece of text",
	}.Make()

	if txt.Embodiment() != "block" {
		t.Error("could not initialize Embodiment field")
	}

	if txt.Tag() != "F" {
		t.Error("could not initialize Tag field")
	}

	if txt.TextEntry() != "This is a piece of text" {
		t.Error("could not initialize TextEntry field")
	}
}

func Test_TextFieldFieldSetting(t *testing.T) {
	txt := &TextField{}

	txt.SetEmbodiment("block")
	if txt.Embodiment() != "block" {
		t.Error("could not set Embodiment fields")
	}

	txt.SetTag("ABC")
	if txt.Tag() != "ABC" {
		t.Error("Could not set Tag field.")
	}

	txt.SetTextEntry("This is some nice content.")
	if txt.TextEntry() != "This is some nice content." {
		t.Error("could not set Content field")
	}
}
