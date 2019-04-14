package main

//go:generate ./gonum -types=ColorEnum,StatusEnum,SushiEnum

type ColorEnum struct {
	Red       string `enum:"RED"`
	LightBlue string `enum:"LIGHT_BLUE"`
}

type StatusEnum struct {
	On  string `enum:"-"`
	Off string `enum:"-"`
}

type SushiEnum struct {
	Maki    string `enum:"MAKI,Rice and filling wrapped in seaweed"`
	Temaki  string `enum:"TEMAKI,Hand rolled into a cone shape"`
	Sashimi string `enum:"SASHIMI,Fish or shellfish served alone without rice"`
}
