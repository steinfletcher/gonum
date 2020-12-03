# gonum

**I don't recommend using this in production, but rather as an example of how one could go about using AST to generate code.**

[![Build Status](https://travis-ci.org/steinfletcher/gonum.svg?branch=master)](https://travis-ci.org/steinfletcher/gonum)

`gonum` is an enum generator for Go. It is inspired by the powerful enum types found in Java. `gonum` has the following capabilities

* Reference and compare enums using values
* Provide a display value for the enumerated fields
* Generate an enum instance from a string factory method
* Generate a slice of display values
* JSON support
* Enum instances can be passed as errors since they implement `Error() string`

## Install

From a github release

```bash
curl https://raw.githubusercontent.com/steinfletcher/gonum/master/download.sh | sh
mv bin/gonum /usr/local/bin
```

OR

```bash
go get -u github.com/steinfletcher/gonum
```

## Example

To define an enum, create a `struct` with the suffix `Enum`. You can define a display value in the `struct` tag. Adding a hyphen will assign the field name to the display value.

You can then generate the enum as follows.

```go
//go:generate gonum -types=ColorEnum,StatusEnum,SushiEnum

// generate an enum with display values. The display values are used for JSON serialization/deserialization
type ColorEnum struct {
	Red       string `enum:"RED"`
	LightBlue string `enum:"LIGHT_BLUE"`
}

// generate an enum with default display values. The display values are set to the field names, e.g. `On` and `Off`
type StatusEnum struct {
	On  string `enum:"-"`
	Off string `enum:"-"`
}

// generate an enum with display values and descriptions.
type SushiEnum struct {
	Maki    string `enum:"MAKI,Rice and filling wrapped in seaweed"`
	Temaki  string `enum:"TEMAKI,Hand rolled into a cone shape"`
	Sashimi string `enum:"SASHIMI,Fish or shellfish served alone without rice"`
}
```

When a description is defined the json is serialized as follows (not yet implemented)

```json
{
  "sushi": {
    "name": "SASHIMI",
    "description": "Fish or shellfish served alone without rice"
  }
}
```

## Consumer api

The generated code would yield the following consumer api

### Create an enum value

```go
a := Red // OR var a Color = Red
```

### Create an enum from a factory method

```go
var name Color = NewColor("RED")
```

### Get the display value

```go
var name string = a.Name() // "RED"
```

### Get all display values

```go
var names []string = ColorNames() // []string{"RED", "BLUE"}
```

### Get all values

```go
var values []Color = ColorValues() // []string{Red, Blue}
```

### Pass as an error

Enums implement `Error() string` which means they can be passed as errors.

```go
var a error = Red
```

## Developing

```bash
go build gonum.go
go generate
go test .
```

OR

```bash
make test
```
