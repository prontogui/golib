// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC
package golib

import (
	"time"
)

type EventTimestampProvider func() time.Time

/*
var _eventTimestampMu sync.Mutex

// The last timestamp before an event cycle (Wait) is entered.
var _eventTimestamp time.Time

// Retrieves the timestamp of most recent event cycle (Wait).
func getEventTimestamp() time.Time {
	_eventTimestampMu.Lock()
	defer _eventTimestampMu.Unlock()
	return _eventTimestamp
}

// Updates the timestamp to the current time.  This should be called before an event cycle (Wait).
func updateEventTimestamp() {
	_eventTimestampMu.Lock()
	defer _eventTimestampMu.Unlock()
	_eventTimestamp = time.Now()
}
*/
