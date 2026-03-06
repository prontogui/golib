// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"fmt"

	"github.com/prontogui/golib/key"
)

type Primitive interface {
	fmt.Stringer
	PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider)
	UnprepareForUpdates()
	LocateNextDescendant(locator *key.PKeyLocator) Primitive
	EgestUpdate(fullupdate bool, fkeys []key.FKey) map[any]any
	IngestUpdate(update map[any]any) error
}
