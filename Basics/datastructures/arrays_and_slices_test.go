package arrays_and_slices

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
}

func Test_maps(t *testing.T) {
	m := map[string]int{"Mario": 42, "Plendl": 43}
	fmt.Println(m)

	// index into map
	fmt.Println(m["Plendl"])

	// delete an entry
	delete(m, "Mario")

	fmt.Println(m)
}
