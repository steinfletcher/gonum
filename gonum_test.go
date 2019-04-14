package main

import (
	"encoding/json"
	"log"
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

	if !reflect.DeepEqual(values, []string{"RED", "LIGHT_BLUE"}) {
		t.Fail()
	}
}

func TestEquality(t *testing.T) {
	var r1 = Red
	var r2 = Red
	var g = LightBlue

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
	type A struct {
		X string `json:"x"`
		Y Color  `json:"y"`
	}
	data := []byte(`{"x":"x","y":"RED"}`)

	a := new(A)
	err := json.Unmarshal(data, a)
	if err != nil {
		t.Fail()
	}

	expected := A{X: "x", Y: Red}
	if *a != expected {
		t.Fatalf("expected x")
	}
}

func TestUnmarshalJSON_Error(t *testing.T) {
	type A struct {
		X string `json:"x"`
		Y Color  `json:"y"`
	}
	data := []byte(`{"x":"x","y":"NOT_A_COLOR"}`)

	a := new(A)
	err := json.Unmarshal(data, a)

	if err == nil {
		log.Fatal("expected an error")
	}
	if err.Error() != "'NOT_A_COLOR' is not a valid value for type" {
		t.Fatalf("expected err: %v", err)
	}
}

func TestSupportsDescription(t *testing.T) {
	e := Maki

	if e.Description() != "Rice and filling wrapped in seaweed" {
		t.Fatalf("expected description but got %s", e.Description())
	}
}

func TestSupportsSerializationWithDescription(t *testing.T) {
	e := Maki

	bytes, err := e.MarshalJSON()

	if err != nil {
		t.Fatal(err)
	}
	actual := string(bytes)
	expected := `{"name":"MAKI","description":"Rice and filling wrapped in seaweed"}`
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("incorrect json, \nexpected: %s\nreceived: %s", expected, actual)
	}
}

func TestSupportsUsageAsError(t *testing.T) {
	var e error = AccountLocked

	if e.Error() != "ACCOUNT_LOCKED" {
		t.Fail()
	}
}

