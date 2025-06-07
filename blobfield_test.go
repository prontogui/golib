// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"os"
	"reflect"
	"testing"

	"github.com/prontogui/golib/key"
)

func Test_BlobSetAndGet(t *testing.T) {
	f := BlobField{}

	f.Set([]byte{34, 200, 90, 1, 0})

	if !reflect.DeepEqual(f.Get(), []byte{34, 200, 90, 1, 0}) {
		t.Fatal("cannot set blob and get the same value back.")
	}
}

func Test_BlobPrepareForUpdates(t *testing.T) {
	f := BlobField{}

	f.PrepareForUpdates(10, key.NewPKey(50), 0, getTestOnsetFunc())

	verifyFieldPreppedForUpdate(t, &f.FieldBase)

	f.Set([]byte{1, 2, 3})

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_BlobEgestValue(t *testing.T) {

	f := BlobField{}
	f.Set([]byte{10, 20, 30})

	v := f.EgestValue()
	ba, ok := v.([]byte)
	if !ok {
		t.Fatal("unable to convert value to []byte")
	}
	if !reflect.DeepEqual(ba, []byte{10, 20, 30}) {
		t.Fatal("incorrect value returned")
	}
}

func Test_BlobIngestUpdate(t *testing.T) {

	f := BlobField{}
	err := f.IngestValue([]byte{1, 2, 3})

	if err != nil {
		t.Fatal("error returned from IngestValue.  Not expected an error.")
	}

	bytes := f.Get()

	if len(bytes) != 3 {
		t.Fatal("ingesting value for Blob doesn't return correct number of bytes")
	}

	if bytes[0] != 1 {
		t.Fatal("element 0 of ingested bytes is not the correct value")
	}

	if bytes[1] != 2 {
		t.Fatal("element 1 of ingested bytes is not the correct value")
	}

	if bytes[2] != 3 {
		t.Fatal("element 2 of ingested bytes is not the correct value")
	}
}

func Test_BlobIngestWrongValueType(t *testing.T) {

	f := BlobField{}
	err := f.IngestValue("something")

	if err == nil {
		t.Fatal("no error returned from IngestValue.  Expecting an error due to wrong value type.")
	}
}

func Test_BlobLoadFromFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "blobfield_load_test_file")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	content := []byte{1, 2, 3, 4, 5}
	if _, err := tmpFile.Write(content); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tmpFile.Close() // Ensure file is closed before reading

	var f BlobField
	err = f.LoadFromFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("LoadFromFile returned error: %v", err)
	}
	if !reflect.DeepEqual(f.Get(), content) {
		t.Errorf("blob content mismatch: got %v, want %v", f.Get(), content)
	}
}

func Test_BlobLoadFromFile_FileNotExist(t *testing.T) {
	var f BlobField
	err := f.LoadFromFile("nonexistent_file_123456789")
	if err == nil {
		t.Fatal("expected error when loading from non-existent file, got nil")
	}
}

func Test_BlobSaveToFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "blobfield_save_test_file")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFileName)

	content := []byte{9, 8, 7, 6, 5}
	var f BlobField
	f.Set(content)

	err = f.SaveToFile(tmpFileName)
	if err != nil {
		t.Fatalf("SaveToFile returned error: %v", err)
	}

	readContent, err := os.ReadFile(tmpFileName)
	if err != nil {
		t.Fatalf("failed to read file after SaveToFile: %v", err)
	}
	if !reflect.DeepEqual(readContent, content) {
		t.Errorf("file content mismatch: got %v, want %v", readContent, content)
	}
}

func Test_BlobSaveToFile_FileCreateError(t *testing.T) {
	var f BlobField
	f.Set([]byte{1, 2, 3})

	// Try to save to an invalid directory path
	invalidPath := "/invalid_dir/should_fail_blobfield"
	err := f.SaveToFile(invalidPath)
	if err == nil {
		t.Fatal("expected error when saving to invalid file path, got nil")
	}
}
