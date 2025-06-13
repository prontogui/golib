// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"bytes"
	"os"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_ImageAttachedFields(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)
	verifyAllFieldsAttached(t, image.PrimitiveBase, "Embodiment", "Image", "Tag")
}

func Test_ImageMake1(t *testing.T) {
	image, err := ImageWith{
		Embodiment: "black-white",
		Image:      []byte{0, 1, 2},
		Tag:        "F",
	}.Make()

	if err != nil {
		t.Error("unexpected error returned")
	}

	if image.Embodiment() != "black-white" {
		t.Error("could not initialize Embodiment field")
	}

	if len(image.Image()) != 3 {
		t.Error("could not initialize Image field")
	}

	if image.Tag() != "F" {
		t.Error("could not initialize Tag field")
	}
}

func Test_ImageMake2(t *testing.T) {
	image, err := ImageWith{Embodiment: "black-white", FromFile: "gopher.png"}.Make()

	if err != nil {
		t.Error("unexpected error returned")
	}

	if image.Embodiment() != "black-white" {
		t.Error("Could not initialize Embodiment field.")
	}

	if len(image.Image()) == 0 {
		t.Error("Could not initialize Image field using FromFile.")
	}
}

func Test_ImageFieldSetting(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)

	image.SetEmbodiment("black-white")
	if image.Embodiment() != "black-white" {
		t.Error("Could not set Embodiment field.")
	}

	image.SetImage([]byte{0, 1, 2})
	if len(image.Image()) != 3 {
		t.Error("Could not set Image field.")
	}

	image.SetTag("ABC")
	if image.Tag() != "ABC" {
		t.Error("Could not set Tag field.")
	}
}

func Test_ImageLoadFromFile_Error(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)

	// Try to load from a non-existent file, should return an error
	err := image.LoadFromFile("nonexistent_file.img")
	if err == nil {
		t.Error("expected error when loading from non-existent file, got nil")
	}
}

func Test_ImageSaveToFile_Error(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)

	// Set some image data
	image.SetImage([]byte{1, 2, 3})

	// Try to save to an invalid path, should return an error
	err := image.SaveToFile("/invalid_path/test.img")
	if err == nil {
		t.Error("expected error when saving to invalid path, got nil")
	}
}

func Test_ImageLoadAndSaveToFile_Success(t *testing.T) {
	image := &Image{}
	image.PrepareForUpdates(key.NewPKey(), nil)

	// Create a temporary file with some data
	tmpfile, err := os.CreateTemp("", "testimg-*.img")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	data := []byte{10, 20, 30, 40}
	if _, err := tmpfile.Write(data); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpfile.Close()

	// Load from the temp file
	err = image.LoadFromFile(tmpfile.Name())
	if err != nil {
		t.Errorf("unexpected error loading from file: %v", err)
	}
	if got := image.Image(); !bytes.Equal(got, data) {
		t.Errorf("image data mismatch after loading from file, got %v, want %v", got, data)
	}

	// Save to another temp file
	tmpfile2, err := os.CreateTemp("", "testimg-save-*.img")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpfile2.Close()
	defer os.Remove(tmpfile2.Name())

	err = image.SaveToFile(tmpfile2.Name())
	if err != nil {
		t.Errorf("unexpected error saving to file: %v", err)
	}

	// Read back and compare
	savedData, err := os.ReadFile(tmpfile2.Name())
	if err != nil {
		t.Fatalf("failed to read saved file: %v", err)
	}
	if !bytes.Equal(savedData, data) {
		t.Errorf("saved file data mismatch, got %v, want %v", savedData, data)
	}
}
