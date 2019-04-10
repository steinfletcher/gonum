package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	var a = Red

	name := a.Name()

	if !reflect.DeepEqual(name, "RED") {
		t.Fail()
	}
}

func TestNewColorFromValue(t *testing.T) {
	color, _ := NewColor("RED")

	if color != Red {
		t.Fail()
	}
}

func TestNewColorFromValue_Failure(t *testing.T) {
	_, err := NewColor("NOT_A_COLOR")

	if err == nil {
		t.Fatal("expected err")
	}
	if err.Error() != "'NOT_A_COLOR' is not a valid value for type" {
		t.Fatalf("incorrect err: %s", err.Error())
	}
}

func TestNames(t *testing.T) {
	values := ColorNames()

	if !reflect.DeepEqual(values, []string{"RED", "BLUE"}) {
		t.Fail()
	}
}

func TestEquality(t *testing.T) {
	var r1 = Red
	var r2 = Red
	var g = Blue

	if r1 == g {
		t.Fail()
	}

	if r1 != r2 {
		t.Fail()
	}
}

func TestJSONMarshal(t *testing.T) {
	type A struct {
		X string `json:"x"`
		Y Color  `json:"y"`
	}

	a := A{
		X: "x",
		Y: Red,
	}

	bytes, err := json.Marshal(&a)
	if err != nil {
		t.Fail()
	}

	if string(bytes) != `{"x":"x","y":"RED"}` {
		t.Fail()
	}
}

func TestUnmarshalJSON(t *testing.T) {
	color := new(Color)

	err := color.UnmarshalJSON([]byte(`"BLUE"`))

	if err != nil {
		t.Fail()
	}
	if *color != Blue {
		t.Fatalf("expected 'BLUE' but got '%v'", color)
	}
	// assert we are not mutating the global instances
	if Red.name != "RED" {
		panic("mutation of name")
	}
	if Red.value != "Red" {
		panic("mutation of value")
	}
}
