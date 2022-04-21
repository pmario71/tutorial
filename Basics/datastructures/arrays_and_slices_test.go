package datastructures

import (
	"fmt"
	"testing"
)

func Test_ArrayInitializing(t *testing.T) {
	arr := []int{1, 2, 3}

	fmt.Println(arr)
}

func Test_Slice_from_Array(t *testing.T) {
	arr := []int{1, 2, 3}

	fmt.Println(arr[1:])
	fmt.Println(arr[:2])

	// append slice to array
	arr2 := []int{4, 5}
	arr2 = append(arr, arr2...)

	fmt.Println(arr2)
}

func Test_maps(t *testing.T) {
	m := map[string]int{"John": 42, "Doe": 43}
	fmt.Println(m)

	// index into map
	fmt.Println(m["Doe"])

	// delete an entry
	delete(m, "John")

	fmt.Println(m)
}
