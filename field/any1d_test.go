package field

import (
	"reflect"
	"testing"

	"github.com/prontogui/golib/primitive"
)

func Test_Any1DSetAndGet(t *testing.T) {
	f := Any1D{}

	actuals_i, _ := generateTestData1D()
	f.Set(actuals_i)

	expected_i, _ := generateTestData1D()

	if !reflect.DeepEqual(f.Get(), expected_i) {
		t.Fatal("cannot set any array and get the same value back.")
	}
}

func Test_Any1DPrepareForUpdates(t *testing.T) {
	f := Any1D{}

	values_i, values_p := generateTestData1D()

	f.Set(values_i)

	f.PrepareForUpdates(10, 50, getTestOnsetFunc())

	verifyStashUpdateInfo(t, &f.Reserved)

	for i, p := range values_p {
		if !p.prepped {
			t.Errorf("array element (%d) was not prepared correctly", i)
		}
	}

	f.Set(values_i)

	if !testOnsetCalled {
		t.Error("onset was not called")
	}
}

func Test_Any1DEgestValue(t *testing.T) {
	f := Any1D{}
	f.Set([]primitive.Interface{&TestPrimitive{s: "abc"}, &TestPrimitive{s: "xyz"}})
	v := f.EgestValue()
	a, ok := v.([]any)
	if !ok {
		t.Fatal("cannot convert value to []any")
	}
	if len(a) != 2 {
		t.Fatal("wrong number of elements returned.  Expecting 2 elements")
	}
	m1, ok := a[0].(map[any]any)
	if !ok {
		t.Fatal("cannot convert element to map[any]any")
	}
	m1v, ok := m1["s"].(string)
	if !ok {
		t.Fatal("cannot convert element map item to string")
	}
	if m1v != "abc" {
		t.Fatal("wrong string value for element")
	}
	m2, ok := a[1].(map[any]any)
	if !ok {
		t.Fatal("cannot convert element to map[any]any")
	}
	m2v, ok := m2["s"].(string)
	if !ok {
		t.Fatal("cannot convert element map item to string")
	}
	if m2v != "xyz" {
		t.Fatal("wrong string value for element")
	}
}

func createAny1DForTest() (*Any1D, []*TestPrimitive) {
	f := &Any1D{}
	p1 := &TestPrimitive{}
	p2 := &TestPrimitive{}
	f.Set([]primitive.Interface{p1, p2})
	return f, []*TestPrimitive{p1, p2}
}

func Test_Any1DIngestUpdate(t *testing.T) {

	f, tps := createAny1DForTest()

	m1 := map[any]any{"s": "Hello"}
	m2 := map[any]any{"s": "World"}

	err := f.IngestValue([]any{m1, m2})
	if err != nil {
		t.Fatalf("unexpected error returned:  %s", err.Error())
	}

	if tps[0].s != "Hello" {
		t.Fatal("primitive #1 not updated correctly")
	}

	if tps[1].s != "World" {
		t.Fatal("primitive #2 not updated correctly")
	}
}

func Test_Any1DIngestUpdateInvalid1(t *testing.T) {

	f := Any1D{}
	p1 := &TestPrimitive{}
	p2 := &TestPrimitive{}
	f.Set([]primitive.Interface{p1, p2})

	err := f.IngestValue(3453)
	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "invalid update" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}

func Test_Any1DIngestUpdateInvalid2(t *testing.T) {

	f := Any1D{}
	p1 := &TestPrimitive{}
	p2 := &TestPrimitive{}
	f.Set([]primitive.Interface{p1, p2})

	err := f.IngestValue([]any{"Hello", "World"})

	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "invalid update" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}

func Test_Any1DIngestUpdateInvalidNumPrimitives(t *testing.T) {

	f := Any1D{}
	p1 := &TestPrimitive{}
	p2 := &TestPrimitive{}
	f.Set([]primitive.Interface{p1, p2})

	m1 := map[any]any{"s": "Hello"}

	err := f.IngestValue([]any{m1})

	if err == nil {
		t.Fatal("no error returned for an invalid update")
	}
	if err.Error() != "number of primitives in update does not equal existing primitives" {
		t.Fatalf("wrong error was returned:  %s", err.Error())
	}
}
