// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"github.com/prontogui/golib/key"
)

// A card displays a main item with optional leading, trailing, and sub item.  It is often used to
// display a summary of a record such as a contact, a product, or a transaction.
type CardWith struct {
	Embodiment   string
	LeadingItem  Primitive
	MainItem     Primitive
	SubItem      Primitive
	Tag          string
	TrailingItem Primitive
}

// Makes a new Card with specified field values.
func (w CardWith) Make() *Card {
	card := &Card{}
	card.embodiment.Set(w.Embodiment)
	card.leadingItem.Set(w.LeadingItem)
	card.mainItem.Set(w.MainItem)
	card.subItem.Set(w.SubItem)
	card.tag.Set(w.Tag)
	card.trailingItem.Set(w.TrailingItem)
	return card
}

// A card displays a main item with optional leading, trailing, and sub item.  It is often used to
// display a summary of a record such as a contact, a product, or a transaction.
type Card struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment   StringField
	leadingItem  AnyField
	mainItem     AnyField
	subItem      AnyField
	tag          StringField
	trailingItem AnyField
}

// Creates a new Card and assigns the main item.
func NewCard(mainItem Primitive) *Card {
	return CardWith{MainItem: mainItem}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (card *Card) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	card.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &card.embodiment},
			{key.FKey_LeadingItem, &card.leadingItem},
			{key.FKey_MainItem, &card.mainItem},
			{key.FKey_SubItem, &card.subItem},
			{key.FKey_Tag, &card.tag},
			{key.FKey_TrailingItem, &card.trailingItem},
		}
	})
}

// Returns a string representation of this primitive:  the mainItem.
// Implements of fmt:Stringer interface.
func (card *Card) String() string {
	var p = card.mainItem.p
	if p != nil {
		return "Card: <" + p.String() + ">"
	}
	return "Card: <nil>"
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (card *Card) Embodiment() string {
	return card.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (card *Card) SetEmbodiment(s string) *Card {
	card.embodiment.Set(s)
	return card
}

// Returns the leading item for this card.
func (card *Card) LeadingItem() Primitive {
	return card.leadingItem.Get()
}

// Sets the leading item for this card.
func (card *Card) SetLeadingItem(p Primitive) *Card {
	card.leadingItem.Set(p)
	return card
}

// Returns the main item for this card.
func (card *Card) MainItem() Primitive {
	return card.mainItem.Get()
}

// Sets the main item for this card.
func (card *Card) SetMainItem(p Primitive) *Card {
	card.mainItem.Set(p)
	return card
}

// Returns the sub item for this card.
func (card *Card) SubItem() Primitive {
	return card.subItem.Get()
}

// Sets the sub item for this card.
func (card *Card) SetSubItem(p Primitive) *Card {
	card.subItem.Set(p)
	return card
}

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (card *Card) Tag() string {
	return card.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Checks as Table cells.
func (card *Card) SetTag(s string) *Card {
	card.tag.Set(s)
	return card
}

// Returns the trailing item for this card.
func (card *Card) TrailingItem() Primitive {
	return card.trailingItem.Get()
}

// Sets the trailing item for this card.
func (card *Card) SetTrailingItem(p Primitive) *Card {
	card.trailingItem.Set(p)
	return card
}
