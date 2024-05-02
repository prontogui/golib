package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TristateWith struct {
	label string
	state int
}

func (w TristateWith) Make() *Tristate {
	tri := &Tristate{}
	tri.label.Set(w.label)
	tri.state.Set(w.state)
	return tri
}

type Tristate struct {
	// Mix-in the common guts for primitives (B-side fields, implements primitive interface, etc.)
	Reserved

	label   field.String
	state   field.Integer
	changed field.Event
}

func (tri *Tristate) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	tri.AttachField("Label", &tri.label)
	tri.AttachField("State", &tri.state)
	tri.AttachField("Changed", &tri.changed)

	// Prepare all fields for updates
	tri.Reserved.PrepareForUpdates(pkey, onset)
}

func (tri *Tristate) Label() string {
	return tri.label.Get()
}

func (tri *Tristate) SetLabel(s string) {
	tri.label.Set(s)
}

func (tri *Tristate) State() int {
	return tri.state.Get()
}

func (tri *Tristate) SetState(i int) {
	tri.state.Set(i)
}

func (tri *Tristate) Changed() bool {
	return tri.changed.Get()
}