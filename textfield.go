package golib

import (
	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type TextFieldWith struct {
	TextEntry  string
	Embodiment string
}

func (w TextFieldWith) Make() *TextField {
	textField := &TextField{}
	textField.textEntry.Set(w.TextEntry)
	textField.embodiment.Set(w.Embodiment)
	return textField
}

type TextField struct {
	// Mix-in the common guts for primitives
	Reserved

	textEntry  field.String
	embodiment field.String
}

func (txt *TextField) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	txt.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &txt.embodiment},
			{key.FKey_TextEntry, &txt.textEntry},
		}
	})
}

func (txt *TextField) TextEntry() string {
	return txt.textEntry.Get()
}

func (txt *TextField) SetTextEntry(s string) {
	txt.textEntry.Set(s)
}

func (txt *TextField) Embodiment() string {
	return txt.embodiment.Get()
}

func (txt *TextField) SetEmbodiment(s string) {
	txt.embodiment.Set(s)
}