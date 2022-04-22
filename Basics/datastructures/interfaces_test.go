package datastructures

import (
	"fmt"
	"testing"
)

// interface Greeter
type Greeter interface {
	greet() string
}

// Human
type Human struct {
	Name string
}

func (h *Human) greet() string {
	return fmt.Sprintf("Hi, I am %v!", h.Name)
}

// ---------------------------------

func callInterface(w Greeter) {
	fmt.Println(w.greet())
}

func Test_calling_an_interface(t *testing.T) {
	h := Human{Name: "John"}

	// pass human obj to method taking a Greeter
	callInterface(&h)
}
