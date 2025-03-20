// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ChoiceAttachedFields(t *testing.T) {
	cmd := &Choice{}
	cmd.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, cmd.PrimitiveBase, "Choice", "Choices", "Embodiment", "Tag")
}

func Test_ChoiceMake(t *testing.T) {
	choice := ChoiceWith{
		Choice:     "Apple",
		Choices:    []string{"Apple", "Orange"},
		Embodiment: "checkmark",
		Tag:        "F",
	}.Make()

	if choice.Choice() != "Apple" {
		t.Error("Could not initialize Choice field.")
	}

	if !reflect.DeepEqual(choice.Choices(), []string{"Apple", "Orange"}) {
		t.Error("Could not initialize Choices field.")
	}

	if choice.Embodiment() != "checkmark" {
		t.Error("Could not initialize Embodiment field.")
	}

	if choice.Tag() != "F" {
		t.Error("Could not initialize Tag field.")
	}
}

func Test_ChoiceFieldSettings(t *testing.T) {
	choice := &Choice{}
	choice.PrepareForUpdates(key.NewPKey(), nil)

	choice.SetChoice("Big Fish")
	if choice.Choice() != "Big Fish" {
		t.Error("Could not set Choice field.")
	}

	choice.SetChoices([]string{"mary", "john", "paul"})
	if !reflect.DeepEqual(choice.Choices(), []string{"mary", "john", "paul"}) {
		t.Error("Could not set Choices field.")
	}

	choice.SetChoicesVA("nancy", "tom", "bob")
	if !reflect.DeepEqual(choice.Choices(), []string{"nancy", "tom", "bob"}) {
		t.Error("Could not set Choices field using variadic arguments")
	}

	choice.SetEmbodiment("checkmark")
	if choice.Embodiment() != "checkmark" {
		t.Error("Could not set Embodiment field")
	}

	choice.SetTag("ABC")
	if choice.Tag() != "ABC" {
		t.Error("Could not set Tag field.")
	}
}

func Test_ChoiceSetIndex(t *testing.T) {

	testFunc := func(index int, expectedChoice string) {
		choice := NewChoice("Apple", "Orange", "Peach")

		if choice.SetChoiceIndex(0) != choice {
			t.Error("SetChoiceIndex didn't return a reference to choice.")
		}

		actualChoice := choice.Choice()

		if actualChoice != "Apple" {
			t.Errorf("Returned Choice = %s for index = %d.  Expecting Choice to be %s", actualChoice, index, expectedChoice)
		}
	}

	// Test cases...
	testFunc(0, "Apple")
	testFunc(1, "Orange")
	testFunc(2, "Peach")
	testFunc(-1, "")
	testFunc(3, "")
}

func Test_ChoiceIndex(t *testing.T) {

	testFunc := func(setChoice string, expectedIndex int) {
		choice := NewChoice("Apple", "Orange", "Peach")

		choice.SetChoice(setChoice)
		actualIndex := choice.ChoiceIndex()

		if actualIndex != expectedIndex {
			t.Errorf("ChoiceIndex returned %d.  Expecting %d.", actualIndex, expectedIndex)
		}
	}

	// Test cases...
	testFunc("Apple", 0)
	testFunc("Orange", 1)
	testFunc("Peach", 2)
	testFunc("", -1)
}
