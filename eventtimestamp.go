// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
package golib

import "time"

// The last timestamp before an event cycle (Wait) is entered.
var _eventTimestamp time.Time

// Retrieves the timestamp of most recent event cycle (Wait).
func getEventTimestamp() time.Time {
	return _eventTimestamp
}

// Updates the timestamp to the current time.  This should be called before an event cycle (Wait).
func updateEventTimestamp() {
	_eventTimestamp = time.Now()
}
