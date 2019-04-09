package main

//go:generate ./gonum -types=ColorEnum,StatusEnum

type ColorEnum struct {
	Red  string `enum:"RED"`
	Blue string `enum:"BLUE"`
}

type StatusEnum struct {
	On  string `enum:"-"`
	Off string `enum:"-"`
}
