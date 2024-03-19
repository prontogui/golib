package field

import (
	"testing"
)

func Test_BooleanSetAndGetFalse(t *testing.T) {
	f := Boolean{}

	f.Set(false)

	if f.Get() != false {
		t.Fatal("cannot set boolean to false and get the same value back.")
	}
}

func Test_BooleanSetAndGetTrue(t *testing.T) {
	f := Boolean{}

	f.Set(true)

	if f.Get() != true {
		t.Fatal("cannot set boolean to true and get the same value back.")
	}
}