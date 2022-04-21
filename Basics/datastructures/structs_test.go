package datastructures

import (
	"fmt"
	"testing"
)

func Test_struct_init(t *testing.T) {
	type Person struct {
		Name      string
		Firstname string
		Age       int
	}

	john := Person{Name: "Doe", Firstname: "John", Age: 42}
	fmt.Println(john)

	john2 := Person{
		Name:      "Doe",
		Firstname: "John",
		Age:       42,  // comma is required
	}
	fmt.Println(john2)
}
