// Code generated by "gonum -types=ColorEnum,StatusEnum,SushiEnum,ErrorsEnum"; DO NOT EDIT.
// See https://github.com/steinfletcher/gonum
package main

import "encoding/json"
import "errors"
import "fmt"

type colorInstanceJsonDescriptionModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var colorInstance = ColorEnum{
	Red:       "RED",
	LightBlue: "LIGHT_BLUE",
}

// Color is the enum that instances should be created from
type Color struct {
	name        string
	value       string
	description string
}

// Enum instances
var Red = Color{name: "RED", value: "Red", description: ""}
var LightBlue = Color{name: "LIGHT_BLUE", value: "LightBlue", description: ""}

// NewColor generates a new Color from the given display value (name)
func NewColor(value string) (Color, error) {
	switch value {
	case "RED":
		return Red, nil
	case "LIGHT_BLUE":
		return LightBlue, nil
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
	case LightBlue:
		return LightBlue.name
	default:
		return ""
	}
}

// String returns the enum display value and is an alias of Name to implement the Stringer interface
func (g Color) String() string {
	return g.Name()
}

// Error returns the enum name and implements the Error interface
func (g Color) Error() string {
	return g.Name()
}

// Description returns the enum description if present. If no description is defined an empty string is returned
func (g Color) Description() string {
	switch g {
	case Red:
		return ""
	case LightBlue:
		return ""
	default:
		return ""
	}
}

// ColorNames returns the displays values of all enum instances as a slice
func ColorNames() []string {
	return []string{
		"RED",
		"LIGHT_BLUE",
	}
}

// ColorValues returns all enum instances as a slice
func ColorValues() []Color {
	return []Color{
		Red,
		LightBlue,
	}
}

// MarshalJSON provides json serialization support by implementing the Marshaler interface
func (g Color) MarshalJSON() ([]byte, error) {
	if g.Description() != "" {
		m := colorInstanceJsonDescriptionModel{
			Name:        g.Name(),
			Description: g.Description(),
		}
		return json.Marshal(m)
	}
	return json.Marshal(g.Name())
}

// UnmarshalJSON provides json deserialization support by implementing the Unmarshaler interface
func (g *Color) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	var value string
	switch v.(type) {
	case map[string]interface{}:
		value = v.(map[string]interface{})["name"].(string)
	case string:
		value = v.(string)
	}

	instance, createErr := NewColor(value)
	if createErr != nil {
		return createErr
	}

	g.name = instance.name
	g.value = instance.value
	g.description = instance.description

	return nil
}

type statusInstanceJsonDescriptionModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var statusInstance = StatusEnum{
	On:  "On",
	Off: "Off",
}

// Status is the enum that instances should be created from
type Status struct {
	name        string
	value       string
	description string
}

// Enum instances
var On = Status{name: "On", value: "On", description: ""}
var Off = Status{name: "Off", value: "Off", description: ""}

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
		return ""
	}
}

// String returns the enum display value and is an alias of Name to implement the Stringer interface
func (g Status) String() string {
	return g.Name()
}

// Error returns the enum name and implements the Error interface
func (g Status) Error() string {
	return g.Name()
}

// Description returns the enum description if present. If no description is defined an empty string is returned
func (g Status) Description() string {
	switch g {
	case On:
		return ""
	case Off:
		return ""
	default:
		return ""
	}
}

// StatusNames returns the displays values of all enum instances as a slice
func StatusNames() []string {
	return []string{
		"On",
		"Off",
	}
}

// StatusValues returns all enum instances as a slice
func StatusValues() []Status {
	return []Status{
		On,
		Off,
	}
}

// MarshalJSON provides json serialization support by implementing the Marshaler interface
func (g Status) MarshalJSON() ([]byte, error) {
	if g.Description() != "" {
		m := statusInstanceJsonDescriptionModel{
			Name:        g.Name(),
			Description: g.Description(),
		}
		return json.Marshal(m)
	}
	return json.Marshal(g.Name())
}

// UnmarshalJSON provides json deserialization support by implementing the Unmarshaler interface
func (g *Status) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	var value string
	switch v.(type) {
	case map[string]interface{}:
		value = v.(map[string]interface{})["name"].(string)
	case string:
		value = v.(string)
	}

	instance, createErr := NewStatus(value)
	if createErr != nil {
		return createErr
	}

	g.name = instance.name
	g.value = instance.value
	g.description = instance.description

	return nil
}

type sushiInstanceJsonDescriptionModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var sushiInstance = SushiEnum{
	Maki:    "MAKI",
	Temaki:  "TEMAKI",
	Sashimi: "SASHIMI",
}

// Sushi is the enum that instances should be created from
type Sushi struct {
	name        string
	value       string
	description string
}

// Enum instances
var Maki = Sushi{name: "MAKI", value: "Maki", description: "Rice and filling wrapped in seaweed"}
var Temaki = Sushi{name: "TEMAKI", value: "Temaki", description: "Hand rolled into a cone shape"}
var Sashimi = Sushi{name: "SASHIMI", value: "Sashimi", description: "Fish or shellfish served alone without rice"}

// NewSushi generates a new Sushi from the given display value (name)
func NewSushi(value string) (Sushi, error) {
	switch value {
	case "MAKI":
		return Maki, nil
	case "TEMAKI":
		return Temaki, nil
	case "SASHIMI":
		return Sashimi, nil
	default:
		return Sushi{}, errors.New(
			fmt.Sprintf("'%s' is not a valid value for type", value))
	}
}

// Name returns the enum display value
func (g Sushi) Name() string {
	switch g {
	case Maki:
		return Maki.name
	case Temaki:
		return Temaki.name
	case Sashimi:
		return Sashimi.name
	default:
		return ""
	}
}

// String returns the enum display value and is an alias of Name to implement the Stringer interface
func (g Sushi) String() string {
	return g.Name()
}

// Error returns the enum name and implements the Error interface
func (g Sushi) Error() string {
	return g.Name()
}

// Description returns the enum description if present. If no description is defined an empty string is returned
func (g Sushi) Description() string {
	switch g {
	case Maki:
		return "Rice and filling wrapped in seaweed"
	case Temaki:
		return "Hand rolled into a cone shape"
	case Sashimi:
		return "Fish or shellfish served alone without rice"
	default:
		return ""
	}
}

// SushiNames returns the displays values of all enum instances as a slice
func SushiNames() []string {
	return []string{
		"MAKI",
		"TEMAKI",
		"SASHIMI",
	}
}

// SushiValues returns all enum instances as a slice
func SushiValues() []Sushi {
	return []Sushi{
		Maki,
		Temaki,
		Sashimi,
	}
}

// MarshalJSON provides json serialization support by implementing the Marshaler interface
func (g Sushi) MarshalJSON() ([]byte, error) {
	if g.Description() != "" {
		m := sushiInstanceJsonDescriptionModel{
			Name:        g.Name(),
			Description: g.Description(),
		}
		return json.Marshal(m)
	}
	return json.Marshal(g.Name())
}

// UnmarshalJSON provides json deserialization support by implementing the Unmarshaler interface
func (g *Sushi) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	var value string
	switch v.(type) {
	case map[string]interface{}:
		value = v.(map[string]interface{})["name"].(string)
	case string:
		value = v.(string)
	}

	instance, createErr := NewSushi(value)
	if createErr != nil {
		return createErr
	}

	g.name = instance.name
	g.value = instance.value
	g.description = instance.description

	return nil
}

type errorsInstanceJsonDescriptionModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var errorsInstance = ErrorsEnum{
	InvalidCredentials: "INVALID_CREDENTIALS",
	AccountLocked:      "ACCOUNT_LOCKED",
}

// Errors is the enum that instances should be created from
type Errors struct {
	name        string
	value       string
	description string
}

// Enum instances
var InvalidCredentials = Errors{name: "INVALID_CREDENTIALS", value: "InvalidCredentials", description: "The username or password is not recognised"}
var AccountLocked = Errors{name: "ACCOUNT_LOCKED", value: "AccountLocked", description: "Your account has been locked. Contact support at me@admin.com"}

// NewErrors generates a new Errors from the given display value (name)
func NewErrors(value string) (Errors, error) {
	switch value {
	case "INVALID_CREDENTIALS":
		return InvalidCredentials, nil
	case "ACCOUNT_LOCKED":
		return AccountLocked, nil
	default:
		return Errors{}, errors.New(
			fmt.Sprintf("'%s' is not a valid value for type", value))
	}
}

// Name returns the enum display value
func (g Errors) Name() string {
	switch g {
	case InvalidCredentials:
		return InvalidCredentials.name
	case AccountLocked:
		return AccountLocked.name
	default:
		return ""
	}
}

// String returns the enum display value and is an alias of Name to implement the Stringer interface
func (g Errors) String() string {
	return g.Name()
}

// Error returns the enum name and implements the Error interface
func (g Errors) Error() string {
	return g.Name()
}

// Description returns the enum description if present. If no description is defined an empty string is returned
func (g Errors) Description() string {
	switch g {
	case InvalidCredentials:
		return "The username or password is not recognised"
	case AccountLocked:
		return "Your account has been locked. Contact support at me@admin.com"
	default:
		return ""
	}
}

// ErrorsNames returns the displays values of all enum instances as a slice
func ErrorsNames() []string {
	return []string{
		"INVALID_CREDENTIALS",
		"ACCOUNT_LOCKED",
	}
}

// ErrorsValues returns all enum instances as a slice
func ErrorsValues() []Errors {
	return []Errors{
		InvalidCredentials,
		AccountLocked,
	}
}

// MarshalJSON provides json serialization support by implementing the Marshaler interface
func (g Errors) MarshalJSON() ([]byte, error) {
	if g.Description() != "" {
		m := errorsInstanceJsonDescriptionModel{
			Name:        g.Name(),
			Description: g.Description(),
		}
		return json.Marshal(m)
	}
	return json.Marshal(g.Name())
}

// UnmarshalJSON provides json deserialization support by implementing the Unmarshaler interface
func (g *Errors) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	var value string
	switch v.(type) {
	case map[string]interface{}:
		value = v.(map[string]interface{})["name"].(string)
	case string:
		value = v.(string)
	}

	instance, createErr := NewErrors(value)
	if createErr != nil {
		return createErr
	}

	g.name = instance.name
	g.value = instance.value
	g.description = instance.description

	return nil
}
