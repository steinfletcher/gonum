# gonum

`gonum` is an enum generator for Go. It is inspired by the powerful enum types found in Java. `gonum` has the following capabilities

* Reference an compare enums using values
* Provide a display value for the enumerated fields
* Generate an enum instance from a string factory method
* Generate a slice of display values
* JSON marshalling and unmarshalling support

## Install

```bash
go get -u github.com/steinfletcher/gonum
```

## Example

To define an enum, create a `struct` with the suffix `Enum`. You can define a display value in the `struct` tag. Adding a hyphen will assign the field name to the display value.

You can then generate the enum as follows.

```go
//go:generate ./gonum -types=ColorEnum,StatusEnum

type ColorEnum struct {
	Red  string `enum:"RED"`
	Blue string `enum:"BLUE"`
}

type StatusEnum struct {
	On  string `enum:"-"`
	Off string `enum:"-"`
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
