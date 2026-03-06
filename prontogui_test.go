// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"context"
	"testing"
)

func Test_NewProntoGUI(t *testing.T) {
	pg := NewProntoGUI()
	if pg == nil {
		t.Fatal("NewProntoGUI returned nil")
	}
}

func Test_StartAndStopServing(t *testing.T) {
	pg := NewProntoGUI()
	err := pg.StartServing("", 0)
	if err != nil {
		t.Fatalf("StartServing failed: %v", err)
	}
	pg.StopServing()
}

func Test_AcceptSession_AfterStop(t *testing.T) {
	pg := NewProntoGUI()
	err := pg.StartServing("", 0)
	if err != nil {
		t.Fatalf("StartServing failed: %v", err)
	}
	pg.StopServing()

	_, err = pg.AcceptSession(context.Background())
	if err == nil {
		t.Fatal("expected error from AcceptSession after StopServing")
	}
}
