// Copyright 2025 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package key

const INVALID_INDEX = -1

type PKey []int

func EmptyPKey() PKey {
	return PKey{}
}

// Create a PKey from the supplied indices.
func NewPKey(indices ...int) PKey {
	pk := make([]int, len(indices))
	copy(pk, indices)
	return pk
}

// Create a new PKey from a list of indices represented as any type.
func NewPKeyFromAny(indices ...any) PKey {
	pk := make([]int, len(indices))
	for level, index := range indices {
		pk[level] = int(index.(uint64))
	}
	return pk
}

// Return true if this PKey is equal to the other PKey.
func (pk PKey) EqualTo(topk PKey) bool {
	if len(pk) != len(topk) {
		return false
	}
	for i, v := range pk {
		if v != topk[i] {
			return false
		}
	}
	return true
}

// Returns a new PKey with the supplied index added to the end.
func (pk PKey) AddLevel(index int) PKey {
	origlen := len(pk)
	newpk := make([]int, origlen+1)
	copy(newpk, pk)
	newpk[origlen] = index
	return newpk
}

// Returns true if this PKey descends from the other PKey.
func (pk PKey) DescendsFrom(frompkey PKey) bool {
	if len(pk) <= len(frompkey) {
		return false
	}

	for level, index := range frompkey {
		if pk[level] != index {
			return false
		}
	}
	return true
}

// Returns the index at supplied level.
func (pk PKey) IndexAtLevel(level int) int {
	if level < 0 || level >= len(pk) {
		return INVALID_INDEX
	}
	return pk[level]
}

// Returns the number of levels in the PKey.
func (pk PKey) Len() int {
	return len(pk)
}

// A helper class to locate a Primitive in the model.
type PKeyLocator struct {
	// The PKey being located.
	PKey PKey

	// The current level of the locator.
	LocationLevel int
}

// A helper object to locate a Primitive in the model.
func NewPKeyLocator(pkey PKey) *PKeyLocator {
	loc := &PKeyLocator{}
	loc.PKey = pkey
	loc.LocationLevel = -1
	return loc
}

// Advance the level and return the index at that level.
func (loc *PKeyLocator) NextIndex() int {

	if loc.LocationLevel >= (len(loc.PKey) - 1) {
		panic("PKeyLocator went out of bounds")
	}

	loc.LocationLevel = loc.LocationLevel + 1

	return loc.PKey[loc.LocationLevel]
}

// Return true if the locator is at the last level and therefore
// primitive hase been located.
func (loc *PKeyLocator) Located() bool {
	return loc.LocationLevel == (len(loc.PKey) - 1)
}
