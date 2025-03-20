// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_TimerAttach(t *testing.T) {
	txt := &Timer{}
	txt.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, txt.PrimitiveBase, "Embodiment", "PeriodMs", "Tag")
}

func Test_TimerMake(t *testing.T) {
	timer := TimerWith{
		Embodiment: "block",
		PeriodMs:   20,
		Tag:        "F",
	}.Make()

	if timer.Embodiment() != "block" {
		t.Error("could not initialize Embodiment field")
	}

	if timer.PeriodMs() != 20 {
		t.Error("could not initialize PeriodMs field")
	}

	if timer.Tag() != "F" {
		t.Error("could not initialize Tag field")
	}
}

func Test_TimerFieldSetting(t *testing.T) {
	timer := &Timer{}

	timer.SetEmbodiment("block")
	if timer.Embodiment() != "block" {
		t.Error("could not set Embodiment fields.")
	}

	timer.SetPeriodMs(50)
	if timer.PeriodMs() != 50 {
		t.Error("could not set PeriodMs field.")
	}

	timer.SetTag("ABC")
	if timer.Tag() != "ABC" {
		t.Error("could not set Tag field.")
	}
}
