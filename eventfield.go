// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"time"

	"github.com/prontogui/golib/key"
)

// The EventField is used to capture an event condition happening on the App side.  For example,
// when the user clicks on a Command, we need to know on the server side the moment this happens
// during a Wait cycle.
//
// Unlike other field types, this field doesn't store a value per se, but tracks if a value has been
// injested during the current Wait cycle.  This information is only valid if the Issued() method is
// called during the current Wait cycle.  If another Wait cycle is entered and no value is injested
// by this field, then it will no longer report Issued() = true.
//
// The field type solved a nagging problem when dealing with handling events on the service side,
// especially when Command primitivess are clicked/issued.  The Command doesn't have any state
// that changes like a Check or TextField might have, and the information it produces is momentary
// in nature.
//
// This field requires initialization after creation by setting its TimestampProvider member.  The
// timestamp provider supplies a time value that originates from the time package and uniquely
// represents the current Wait cycle.  A good implementation will keep a central timestamp generated
// from time.Now(), called at the beginning of each Wait cycle, and a central function for returning
// the timestamp.  The central function is assigned to every EventField during construction.
//
// We avoided creating a global timestamp in this module because in the future there could be multiple, concurrent
// Wait cycles going on.  It is also fragile to write unit tests this way.
type EventField struct {
	FieldBase

	// A valid timestamp has been saved.
	validTimestamp bool

	// The timestamp returned from provider at the last time a value was injested.
	// This is used to know if a value was injested during the current Wait cycle.
	eventTimestamp time.Time

	// A function that returns a timestamp uniquely tied to the current Wait cycle.
	TimestampProvider func() time.Time
}

// Makes sure we have a timestamp provider.
func (f *EventField) _checkMissingFunc() {
	if f.TimestampProvider == nil {
		panic("EventField does not have timestampFunc member assigned.  Cannot retrieve timestamp.")
	}
}

// Returns true if the event was assigned during the current Wait cycle, as determined from the
// timestamp provider.
func (f *EventField) Issued() bool {
	if !f.validTimestamp {
		return false
	}

	f._checkMissingFunc()

	// Was a new value injested during the current event cycle?
	return f.TimestampProvider().Sub(f.eventTimestamp) == 0
}

func (f *EventField) PrepareForUpdates(fkey key.FKey, pkey key.PKey, fieldPKeyIndex int, onset key.OnSetFunction) (isContainer bool) {
	f.StashUpdateInfo(fkey, pkey, fieldPKeyIndex, onset)
	return false
}

func (f *EventField) EgestValue() any {
	return false
}

func (f *EventField) IngestValue(value any) error {
	f._checkMissingFunc()
	f.validTimestamp = true
	f.eventTimestamp = f.TimestampProvider()
	return nil
}
