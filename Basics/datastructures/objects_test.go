package datastructures

import "testing"

func Test_constructing_an_object(t *testing.T) {
	m := newUserController()

	if m.userIDPattern == nil {
		t.Error("userIDPattern not initialized!")
	}
}
