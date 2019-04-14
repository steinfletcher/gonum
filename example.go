package main

//go:generate ./gonum -types=ColorEnum,StatusEnum,SushiEnum,ErrorsEnum

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

type ErrorsEnum struct {
	InvalidCredentials string `enum:"INVALID_CREDENTIALS,The username or password is not recognised"`
	AccountLocked      string `enum:"ACCOUNT_LOCKED,Your account has been locked. Contact support at me@admin.com"`
}
