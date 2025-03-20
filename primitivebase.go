// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/key"
)

// PrimitiveBase fields for primitive updates.
type PrimitiveBase struct {
	pkey   key.PKey
	fields []FieldRef
}

type FieldRef struct {
	// The field's key
	fkey key.FKey

	// Reference to the field itself
	field Field
}

func (r *PrimitiveBase) InternalPrepareForUpdates(pkey key.PKey, onset key.OnSetFunction, getFields func() []FieldRef) {

	r.pkey = pkey

	// Attach fields (if not done already)
	if len(r.fields) == 0 {
		r.fields = getFields()
	}

	// Prepare each field for updates
	fieldPKeyIndex := 0
	for _, f := range r.fields {
		if f.field.PrepareForUpdates(f.fkey, pkey, fieldPKeyIndex, onset) {
			fieldPKeyIndex = fieldPKeyIndex + 1
		}
	}
}

func (r *PrimitiveBase) LocateNextDescendant(locator *key.PKeyLocator) Primitive {
	return nil
}

func (r *PrimitiveBase) findField(fkey key.FKey) Field {

	var found Field
	for _, f := range r.fields {
		if f.fkey == fkey {
			found = f.field
			break
		}
	}
	return found
}

func (r *PrimitiveBase) EgestUpdate(fullupdate bool, fkeys []key.FKey) map[any]any {

	update := map[any]any{}

	if fullupdate {
		for _, v := range r.fields {
			fieldvalue := v.field.EgestValue()

			if fieldvalue != nil {
				update[key.FieldnameFor(v.fkey)] = fieldvalue
			}
		}
	} else {
		for _, fkey := range fkeys {

			field := r.findField(fkey)
			if field == nil {
				panic("field not found in primitive")
			}

			fieldvalue := field.EgestValue()

			if fieldvalue != nil {
				update[key.FieldnameFor(fkey)] = fieldvalue
			}
		}
	}

	return update
}

func (r *PrimitiveBase) IngestUpdate(update map[any]any) error {

	for k, v := range update {
		var ok bool

		ks, ok := k.(string)
		if !ok {
			return errors.New("invalid key type.  Expecting a string")
		}

		fkey := key.FKeyFor(ks)
		if fkey == key.INVALID_FIELDNAME {
			return errors.New("invalid field name")
		}

		var found Field
		for _, f := range r.fields {
			if f.fkey == fkey {
				found = f.field
				break
			}
		}

		if found == nil {
			return errors.New("no matching field name in primitive")
		}

		err := found.IngestValue(v)
		if err != nil {
			return err
		}
	}

	return nil
}

// Returns the index of this primitive in a parent container specified by parentLevel as follows:
// parentLevel = 0, immediate parent container
// parentLevel = 1, grandparent
// parentLevel = 2, great grandparent
// And so on.
//
// It returns -1 if parentLevel is a negative number or is invalid given the depth where the primitive belongs.
func (r *PrimitiveBase) IndexOf(parentLevel int) int {
	len := r.pkey.Len()

	if parentLevel >= 0 && parentLevel < len {
		return r.pkey.IndexAtLevel(len - parentLevel - 1)
	}
	return -1
}

// Default implementation of fmt:Stringer interface.
func (r *PrimitiveBase) String() string {
	return ""
}
