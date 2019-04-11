// Code generated by "gonum -types=ColorEnum,StatusEnum"; DO NOT EDIT.

package main

import "encoding/json"
import "errors"
import "fmt"

var colorInstance = ColorEnum{
	Red:  "RED",
	Blue: "BLUE",
}

// Color is the enum that instances should be created from
type Color struct {
	name  string
	value string
}

// Enum instances
var Red = Color{name: "RED", value: "Red"}
var Blue = Color{name: "BLUE", value: "Blue"}

// NewColor generates a new Color from the given display value (name)
func NewColor(value string) (Color, error) {
	switch value {
	case "RED":
		return Red, nil
	case "BLUE":
		return Blue, nil
	default:
		return Color{}, errors.New(
			fmt.Sprintf("'%s' is not a valid value for type", value))
	}
}

// Name returns the enum display value
func (g Color) Name() string {
	switch g {
	case Red:
		return Red.name
	case Blue:
		return Blue.name
	default:
		panic("Could not map enum")
	}
}

// String returns the enum display value and is an alias of Name to implement the Stringer interface
func (g Color) String() string {
	return g.Name()
}

// Names returns the displays values of all enum instances as a slice
func ColorNames() []string {
	return []string{
		"RED",
		"BLUE",
	}
}

// Values returns all enum instances as a slice
func ColorValues() []Color {
	return []Color{
		Red,
		Blue,
	}
}

// MarshalJSON provides json marshalling support by implementing the Marshaler interface
func (g Color) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.Name())
}

// MarshalJSON provides json unmarshalling support by implementing the Unmarshaler interface
func (g *Color) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	instance, createErr := NewColor(v)
	if createErr != nil {
		return createErr
	}

	g.name = instance.name
	g.value = instance.value

	return nil
}

var statusInstance = StatusEnum{
	On:  "On",
	Off: "Off",
}

// Status is the enum that instances should be created from
type Status struct {
	name  string
	value string
}

// Enum instances
var On = Status{name: "On", value: "On"}
var Off = Status{name: "Off", value: "Off"}

// NewStatus generates a new Status from the given display value (name)
func NewStatus(value string) (Status, error) {
	switch value {
	case "On":
		return On, nil
	case "Off":
		return Off, nil
	default:
		return Status{}, errors.New(
			fmt.Sprintf("'%s' is not a valid value for type", value))
	}
}

// Name returns the enum display value
func (g Status) Name() string {
	switch g {
	case On:
		return On.name
	case Off:
		return Off.name
	default:
		panic("Could not map enum")
	}
}

// String returns the enum display value and is an alias of Name to implement the Stringer interface
func (g Status) String() string {
	return g.Name()
}

// Names returns the displays values of all enum instances as a slice
func StatusNames() []string {
	return []string{
		"On",
		"Off",
	}
}

// Values returns all enum instances as a slice
func StatusValues() []Status {
	return []Status{
		On,
		Off,
	}
}

// MarshalJSON provides json marshalling support by implementing the Marshaler interface
func (g Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.Name())
}

// MarshalJSON provides json unmarshalling support by implementing the Unmarshaler interface
func (g *Status) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	instance, createErr := NewStatus(v)
	if createErr != nil {
		return createErr
	}

	g.name = instance.name
	g.value = instance.value

	return nil
}
