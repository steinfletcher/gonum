package main

import (
	"reflect"
	"testing"
)

func TestValue(t *testing.T) {
	var a Color = Red

	name := a.Value()

	if !reflect.DeepEqual(name, "RED") {
		t.Fail()
	}
}

func TestNewColorFromValue(t *testing.T) {
	color := NewColor("RED")

	if color != Red {
		t.Fail()
	}
}

func TestValues(t *testing.T) {
	values := ColorValues()

	if !reflect.DeepEqual(values, []string{"RED", "BLUE"}) {
		t.Fail()
	}
}

func TestEquality(t *testing.T) {
	var r1 Color = Red
	var r2 Color = Red
	var g Color = Blue

	if r1 == g {
		t.Fail()
	}

	if r1 != r2 {
		t.Fail()
	}
}
