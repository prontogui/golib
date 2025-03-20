// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"errors"

	"github.com/prontogui/golib/pgcomm"
)

// The API for working with app-to-server communication.
type ProntoGUI interface {
	StartServing(addr string, port int) error
	StopServing()
	SetGUI(primitives ...Primitive)
	Update() (Primitive, error)
	Wait() (Primitive, error)
}

// Internal data for handling the API of this library
type _ProntoGUI struct {
	pgcomm     *pgcomm.PGComm
	synchro    *Synchro
	isgui      bool
	fullupdate bool
}

// Start up the communication server so we can accept a connection from the app.
func (pg *_ProntoGUI) StartServing(addr string, port int) error {
	pg.fullupdate = true
	return pg.pgcomm.StartServing(addr, port)
}

// Stop the communication and close any open connection.
func (pg *_ProntoGUI) StopServing() {
	pg.pgcomm.StopServing()
}

// Prepares all the primitives that will participate in the GUI.  This must be called at least once before Wait().
//
// Note:  it is possible to reuse primitive instance throughout the same GUI and even across multiple GUIs but there are
// some things to be aware of.
//
//   - If you change a field of a primitive that no longer is part of the active GUI then you could have indesirable behavior
//     or a panic situation.  This is because the primitive is still "prepped" for communication but has since been divorced
//     from the GUI model.  Right now, there is no compelling reason to "unprepare" primitives that no longer participte in
//     the GUI since this would have some impact on performance.  This could change as ProntoGUI usage evolves and perhaps
//     we might add a flexible policy in the future.  For now, don't do anything with non-participating primitives until
//     they're part of the active GUI again (by calling SetGUI).
//
//   - Updates to the primitive coming from the app side will not be discernable.  To explain this, let's say you have the
//     same Command primitive instance in three different views.  When the user clicks on the Command, you will receive
//     on update from the Wait() function that the command was clicked.  However, you be able to know which view it was
//     participating in when the click happened.
//
//   - Updates to the primitive coming from the server side will only affect one instance of the primitive on the app side.
//     For example, if you change the label of a Command primitive that is reused several times in the same GUI, only one of
//     them will visually change.
//
//   - It's best to avoid reusing primitives in the same GUI unless they are static, such as Text.
func (pg *_ProntoGUI) SetGUI(primitives ...Primitive) {
	pg.fullupdate = true
	pg.isgui = true
	pg.synchro.SetTopPrimitives(primitives...)
}

// Sends model updates to the app for rendering, waits for an event to occur in the app that requires server attention,
// and receives model updates from the app.  This function returns the primitive that ended triggered the event that ended
// the wait period.  An error is returend if something went wrong.
func (pg *_ProntoGUI) Wait() (Primitive, error) {

	updateOut, err := pg.verifyGuiIsSetThenGetNextUpdate()
	if err != nil {
		return nil, err
	}

	// Exchange updates and wait until App has an update.
	updateIn, err := pg.pgcomm.ExchangeUpdates(updateOut, false)
	if updateIn == nil || err != nil {
		// Require a full update the next time around
		pg.fullupdate = true
		return nil, err
	}

	// Update the time stamp for recognizing events such as CommandIssued.  This must be called for
	// every update received from the app.
	updateEventTimestamp()

	// Ingest the update and return the primitive that triggered the event or an error if something went wrong
	return pg.synchro.IngestUpdate(updateIn)
}

// Sends model updates to the app for rendering.  If an event occured in the app that requires server attention,
// it receives the model updates from the app returns the primitive that ended triggered the event.  An error is
// returend if something went wrong.
func (pg *_ProntoGUI) Update() (Primitive, error) {

	updateOut, err := pg.verifyGuiIsSetThenGetNextUpdate()
	if err != nil {
		return nil, err
	}

	// Exchange updates but don't wait if nothing available from the App
	updateIn, err := pg.pgcomm.ExchangeUpdates(updateOut, true)

	if err != nil {
		// Require a full update the next time around
		pg.fullupdate = true
		return nil, err
	}

	// No update available, return immediately
	if updateIn == nil {
		return nil, nil
	}

	// Update the time stamp for recognizing events such as CommandIssued.  This must be called for
	// every update received from the app.
	updateEventTimestamp()

	// Ingest the update and return the primitive that triggered the event or an error if something went wrong
	return pg.synchro.IngestUpdate(updateIn)
}

// Verifies that a GUI has been set and then gets the next update to send to the app.
func (pg *_ProntoGUI) verifyGuiIsSetThenGetNextUpdate() ([]byte, error) {
	if !pg.isgui {
		return nil, errors.New("no GUI has been set")
	}

	// Need to send a full update?
	if pg.fullupdate {
		pg.fullupdate = false
		return pg.synchro.GetFullUpdate()
	} else {
		return pg.synchro.GetPartialUpdate()
	}
}

// Creates a new ProntoGUI instance.
func NewProntoGUI() ProntoGUI {
	pg := &_ProntoGUI{}

	pg.pgcomm = pgcomm.NewPGComm()
	pg.synchro = NewSynchro()
	pg.isgui = false
	pg.fullupdate = true

	return pg
}
