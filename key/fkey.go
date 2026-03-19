// Copyright 2024-2026 ProntoGUI, LLC
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// ProntoGUI™ is a trademark of ProntoGUI, LLC

package key

type FKey uint8

const (
	INVALID_FIELDNAME = 255
	INVALID_FKEY      = ""
)

const (

	// ADD NEW FIELDS TO THIS BLOCK - ALPHABETICAL ORDER PLEASE!
	FKey_Checked FKey = iota
	FKey_Choice
	FKey_Choices
	FKey_ChoiceLabels
	FKey_CommandIssued
	FKey_Content
	FKey_Data
	FKey_Embodiment
	FKey_Expanded
	FKey_Exported
	FKey_FrameItems
	FKey_GroupItems
	FKey_HeaderRow
	FKey_Icon
	FKey_IconID
	FKey_ID
	FKey_Image
	FKey_Imported
	FKey_Issued
	FKey_Item
	FKey_Label
	FKey_LabelItem
	FKey_LeadingItem
	FKey_Level
	FKey_ListItems
	FKey_MainItem
	FKey_ModelItem
	FKey_ModelRow
	FKey_Name
	FKey_NumericEntry
	FKey_PeriodMs
	FKey_Ref
	FKey_Rows
	FKey_SelectedItems
	FKey_SelectedRows
	FKey_SelectionChanged
	FKey_SelectionMode
	FKey_Showing
	FKey_State
	FKey_Status
	FKey_SubItem
	FKey_Tag
	FKey_TextEntry
	FKey_TimerFired
	FKey_Title
	FKey_TrailingItem
	FKey_ValidExtensions

	// RESERVED CONSTANT
	FKey_MAXIMUMKEYS
)

var _fkeyToName []string
var _nameToFKey map[string]FKey

func init() {
	_fkeyToName = make([]string, FKey_MAXIMUMKEYS)

	// ADD NEW FIELDS TO THIS BLOCK - ALPHABETICAL ORDER PLEASE!
	_fkeyToName[FKey_Checked] = "Checked"
	_fkeyToName[FKey_Choice] = "Choice"
	_fkeyToName[FKey_Choices] = "Choices"
	_fkeyToName[FKey_ChoiceLabels] = "ChoiceLabels"
	_fkeyToName[FKey_CommandIssued] = "CommandIssued"
	_fkeyToName[FKey_Content] = "Content"
	_fkeyToName[FKey_Data] = "Data"
	_fkeyToName[FKey_Embodiment] = "Embodiment"
	_fkeyToName[FKey_Expanded] = "Expanded"
	_fkeyToName[FKey_Exported] = "Exported"
	_fkeyToName[FKey_FrameItems] = "FrameItems"
	_fkeyToName[FKey_GroupItems] = "GroupItems"
	_fkeyToName[FKey_HeaderRow] = "HeaderRow"
	_fkeyToName[FKey_Icon] = "Icon"
	_fkeyToName[FKey_IconID] = "IconID"
	_fkeyToName[FKey_ID] = "ID"
	_fkeyToName[FKey_Image] = "Image"
	_fkeyToName[FKey_Imported] = "Imported"
	_fkeyToName[FKey_Issued] = "Issued"
	_fkeyToName[FKey_Item] = "Item"
	_fkeyToName[FKey_Label] = "Label"
	_fkeyToName[FKey_LabelItem] = "LabelItem"
	_fkeyToName[FKey_LeadingItem] = "LeadingItem"
	_fkeyToName[FKey_Level] = "Level"
	_fkeyToName[FKey_ListItems] = "ListItems"
	_fkeyToName[FKey_MainItem] = "MainItem"
	_fkeyToName[FKey_ModelItem] = "ModelItem"
	_fkeyToName[FKey_ModelRow] = "ModelRow"
	_fkeyToName[FKey_Name] = "Name"
	_fkeyToName[FKey_NumericEntry] = "NumericEntry"
	_fkeyToName[FKey_PeriodMs] = "PeriodMs"
	_fkeyToName[FKey_Ref] = "Ref"
	_fkeyToName[FKey_Rows] = "Rows"
	_fkeyToName[FKey_SelectedItems] = "SelectedItems"
	_fkeyToName[FKey_SelectedRows] = "SelectedRows"
	_fkeyToName[FKey_SelectionChanged] = "SelectionChanged"
	_fkeyToName[FKey_SelectionMode] = "SelectionMode"
	_fkeyToName[FKey_Showing] = "Showing"
	_fkeyToName[FKey_State] = "State"
	_fkeyToName[FKey_Status] = "Status"
	_fkeyToName[FKey_SubItem] = "SubItem"
	_fkeyToName[FKey_Tag] = "Tag"
	_fkeyToName[FKey_TextEntry] = "TextEntry"
	_fkeyToName[FKey_TimerFired] = "TimerFired"
	_fkeyToName[FKey_Title] = "Title"
	_fkeyToName[FKey_TrailingItem] = "TrailingItem"
	_fkeyToName[FKey_ValidExtensions] = "ValidExtensions"

	_nameToFKey = make(map[string]FKey, FKey_MAXIMUMKEYS)

	for fkey, fname := range _fkeyToName {
		_nameToFKey[fname] = FKey(fkey)
	}
}

func FKeyFor(fieldname string) FKey {

	fkey, ok := _nameToFKey[fieldname]

	if ok {
		return fkey
	}

	return INVALID_FIELDNAME
}

func FieldnameFor(fkey FKey) string {

	if fkey >= FKey_MAXIMUMKEYS {
		return INVALID_FKEY
	}

	return _fkeyToName[fkey]
}
