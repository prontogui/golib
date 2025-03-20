// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_CheckAttachedFields(t *testing.T) {
	check := &Check{}
	check.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, check.PrimitiveBase, "Checked", "Embodiment", "Label", "Tag")
}

func Test_CheckMake(t *testing.T) {
	check := CheckWith{
		Checked:    true,
		Embodiment: "such-and-such",
		Label:      "Option",
		Tag:        "F",
	}.Make()

	if !check.Checked() {
		t.Error("Could not initialize Checked field.")
	}

	if check.Embodiment() != "such-and-such" {
		t.Error("Could not initialize Embodiment field.")
	}

	if check.Label() != "Option" {
		t.Error("Could not initialize Label field.")
	}

	if check.Tag() != "F" {
		t.Error("Could not initialize Tag field.")
	}
}

func Test_CheckFieldSettings(t *testing.T) {

	check := &Check{}
	check.PrepareForUpdates(key.NewPKey(), nil)

	check.SetChecked(true)
	if !check.Checked() {
		t.Error("Could not set Checked field.")
	}

	check.SetEmbodiment("checkmark")
	if check.Embodiment() != "checkmark" {
		t.Error("Could not set Embodiment field.")
	}

	check.SetLabel("Option 1")
	if check.Label() != "Option 1" {
		t.Error("Could not set Label field.")
	}

	check.SetTag("ABC")
	if check.Tag() != "ABC" {
		t.Error("Could not set Tag field.")
	}
}
