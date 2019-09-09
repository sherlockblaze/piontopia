package utils

import "testing"

func TestUpdateParameter(t *testing.T) {
	type TestStruct struct {
		A int
		B string
	}

	a1 = TestStruct{}
	a2 = TestStruct{}

	UpdateParameter(a1, a2)
}
