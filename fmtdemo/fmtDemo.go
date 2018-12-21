package fmtdemo

import (
	"fmt"
)

type Animal struct {
	Name string
	Age  uint
}

func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func ExampleStringer() {
	a := Animal{
		Name: "Bird",
		Age:  2,
	}
	fmt.Println(a)
}
