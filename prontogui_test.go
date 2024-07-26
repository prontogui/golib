// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"fmt"
	"testing"
)

func Test_BasicGUI(t *testing.T) {

	// The following code just builds a very simple GUI and goes into operation.
	// This was developed as a convenient way to test changes to pgcomm.go ang prontogui.go

	// Initialize ProntoGUI
	pgui := NewProntoGUI()
	err := pgui.StartServing("127.0.0.1", 50053)

	if err != nil {
		fmt.Printf("Error trying to start server:  %s", err.Error())
		return
	}

	// Big and bold heading for the GUI
	guiHeading := TextWith{
		Content: "Simple App",
	}.Make()

	cmd := CommandWith{Label: "OK"}.Make()

	pgui.SetGUI(guiHeading, cmd)

	// Loop while handling the events occuring in the GUI
	for {
		// Wait for something to happen in the GUI
		_, err := pgui.Wait()
		if err != nil {
			fmt.Printf("error from Wait() is:  %s\n", err.Error())
			break
		}
	}

	f := StringField{}

	f.Set("abc")

	if f.Get() != "abc" {
		t.Fatal("cannot set string and get the same value back.")
	}

}
