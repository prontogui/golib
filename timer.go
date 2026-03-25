// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package golib

import (
	"github.com/prontogui/golib/key"
)

// A timer is an invisible primitive that fires an event, triggering an update
// to the server.  This is useful for low-precision GUI updates that originate
// on the server side.  An example is updating "live" readings from a running
// process on the server.
type TimerWith struct {
	Embodiment string
	PeriodMs   int
	Status     int
	Tag        string
}

// Makes a new Timer with specified field values.
func (w TimerWith) Make() *Timer {
	timer := &Timer{}

	timer.embodiment.Set(w.Embodiment)
	timer.periodMs.Set(w.PeriodMs)
	timer.status.Set(w.Status)
	timer.tag.Set(w.Tag)
	return timer
}

// A timer is an invisible primitive that fires an event, triggering an update
// to the server.  This is useful for low-precision GUI updates that originate
// on the server side.  An example is updating "live" readings from a running
// process on the server.
type Timer struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	periodMs   IntegerField
	status     IntegerField
	tag        StringField
	timerFired EventField
}

// Create a new Timer with period in milliseconds.
func NewTimer(periodMs int) *Timer {
	return TimerWith{PeriodMs: periodMs}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (tmr *Timer) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, etsprovider EventTimestampProvider) {

	tmr.InternalPrepareForUpdates(pkey, onset, etsprovider, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &tmr.embodiment},
			{key.FKey_PeriodMs, &tmr.periodMs},
			{key.FKey_Status, &tmr.status},
			{key.FKey_Tag, &tmr.tag},
			{key.FKey_TimerFired, &tmr.timerFired},
		}
	})
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (tmr *Timer) Embodiment() string {
	return tmr.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (tmr *Timer) SetEmbodiment(s string) *Timer {
	tmr.embodiment.Set(s)
	return tmr
}

// Returns the time period in milliseconds after which the timer fires an event.  If the
// period is -1 (or any negative number) then the timer is disabled.  A period of 0
// will cause the timer to fire immediately after the primitive is updated.
func (tmr *Timer) PeriodMs() int {
	return tmr.periodMs.Get()
}

// Sets the time period in milliseconds after which the timer fires an event.  If the
// period is -1 (or any negative number) then the timer is disabled.  A period of 0
// will cause the timer to fire immediately after the primitive is updated.
func (tmr *Timer) SetPeriodMs(i int) *Timer {
	tmr.periodMs.Set(i)
	return tmr
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Timers inside containers.
func (tmr *Timer) Tag() string {
	return tmr.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Timers inside containers.
func (tmr *Timer) SetTag(s string) *Timer {
	tmr.tag.Set(s)
	return tmr
}

// Returns true if the command was issued during the current Wait cycle.
func (tmr *Timer) Issued() bool {
	return tmr.timerFired.Issued()
}

// Returns the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *Timer) Status() int {
	return p.status.Get()
}

// Sets the status of the primitive: 0 = visible and enabled,  1 = visible and disabled,
// 2 = hidden and disabled, 3 = collapsed and disabled.
func (p *Timer) SetStatus(i int) *Timer {
	p.status.Set(i)
	return p
}

// Returns the visibility of the group.  This is derived from the Status field.
func (p *Timer) Visible() bool {
	status := p.status.Get()
	return status == 0 || status == 1
}

// Sets the visibility of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 2 (hidden).
func (p *Timer) SetVisible(visible bool) *Timer {
	if visible {
		p.status.Set(0)
	} else {
		p.status.Set(2)
	}
	return p
}

// Returns the enabled status of the primitive.  This is derived from the Status field.
func (p *Timer) Enabled() bool {
	return p.status.Get() == 0
}

// Sets the enabled status of the primitive.  Setting this to true will set Status to 0 (visible & enabled)
// and setting this to false will set Status to 1 (disabled).
func (p *Timer) SetEnabled(enabled bool) *Timer {
	if enabled {
		p.status.Set(0)
	} else {
		p.status.Set(1)
	}
	return p
}

// Returns the collapsed status of the primitive.  This is derived from the Status field.
func (p *Timer) Collapsed() bool {
	return p.status.Get() == 3
}

// Sets the collapsed status of the primitive.  Setting this to true will set Status to 3 (collapsed)
// and setting this to false will set Status to 0 (visible & enabled).
func (p *Timer) SetCollapsed(collapsed bool) *Timer {
	if collapsed {
		p.status.Set(3)
	} else {
		p.status.Set(0)
	}
	return p
}
