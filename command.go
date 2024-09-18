// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

type CommandWith struct {
	Embodiment string
	Label      string
	Status     int
	Tag        string
}

// Makes a new Command with specified field values.
func (w CommandWith) Make() *Command {
	cmd := &Command{}
	cmd.embodiment.Set(w.Embodiment)
	cmd.label.Set(w.Label)
	cmd.status.Set(w.Status)
	cmd.tag.Set(w.Tag)
	return cmd
}

type Command struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	label      StringField
	status     IntegerField
	tag        StringField
}

func (cmd *Command) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	cmd.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &cmd.embodiment},
			{key.FKey_Label, &cmd.label},
			{key.FKey_Status, &cmd.status},
			{key.FKey_Tag, &cmd.tag},
		}
	})
}

func (cmd *Command) Embodiment() string {
	return cmd.embodiment.Get()
}

func (cmd *Command) SetEmbodiment(s string) {
	cmd.embodiment.Set(s)
}

func (cmd *Command) Label() string {
	return cmd.label.Get()
}

func (cmd *Command) SetLabel(s string) {
	cmd.label.Set(s)
}

func (cmd *Command) Status() int {
	return cmd.status.Get()
}

func (cmd *Command) SetStatus(i int) {
	cmd.status.Set(i)
}

func (cmd *Command) Tag() string {
	return cmd.tag.Get()
}

func (cmd *Command) SetTag(s string) {
	cmd.tag.Set(s)
}
