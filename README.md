# gonum

## What's this repo about?

Go doesn't support powerful enum types that you find in other languages like Java. I am considering writing a code generator that will generate a powerful enum type.

I haven't thought this through in great detail, but here are my initial plans.

## Capabilities

It would be nice to have these capabilities:

* Reference an enum instance statically, e.g. it should not be necessary to create an instance. This could be achieved with const values.

* Support a display value. It's sometimes useful to have a human readable representation of the enum for display purposes.

* Provide a factory method to create an enum value from a string

* Provide a method to get all instances of the enum name as a slice of strings

* Provide a method to get all instances of the enum value as a slice of the enum type

## Design

This might be an abuse of structs, but one approach would be to declare the enum as a struct. This is because metadata can be added through tags. I prefer this to recording metadata as comments (which is what https://github.com/abice/go-enum does). Using struct tags provide support for future extensions

```go
type ColorEnum struct {
	Red  string `enum:"name=RED"`
	Blue string `enum:"name=BLUE"`
}
```

## Generated code

This would yield the following generated code

```go
// package private singleton instance holding metadata
var colorEnumInstance = ColorEnum{
	Red:  "RED",
	Blue: "BLUE",
}

type Color uint

const (
	Red = iota
	Blue
)

func (g Color) Name() string {
	switch g {
	case Red:
		return colorEnumInstance.Red
	case Blue:
		return colorEnumInstance.Blue
	default:
		panic(fmt.Sprintf("Could not map enum '%v'", g))
	}
}

// String is an alias of Name to satisfy the Stringer interface
func (g Color) String() string {
	return g.Name()
}


func NewColor(name string) Color {
	switch value {
	case "RED":
		return Red
	case "BLUE":
		return Blue
	default:
		panic(fmt.Sprintf("Could not create enum from name '%v'", value))
	}
}

func ColorValues() []Color {
	return []string{
		Red,
		Blue,
	}
}

func ColorNames() []string {
	return []string{
		"RED",
		"BLUE",
	}
}
```

## Consumer api

The generated code would yield the following consumer api

```go
var a Color = Red

var name string = a.Name() // "RED"

var color Color = NewColor("BLUE") // from string helper to create an enum

var values []Color = ColorValues() // []string{Red, Blue}

var names []string = ColorNames() // []string{"RED", "BLUE"}
```
