package main

import "testing"

type testPairValidation struct {
	str    string
	result bool
}

var testsValidation = []testPairValidation{
	{"[]", true},
	{"{}", true},
	{"[", false},
	{"]", false},
	{"{", false},
	{"}", false},
	{"[}", false},
	{"{]", false},
}

func TestValidate(t *testing.T) {
	for _, pair := range testsValidation {
		s := inputValidation(pair.str)
		actual := s.validate()
		if bool(*actual) != pair.result {
			t.Error(
				"For", pair.str,
				"expected", pair.result,
				"got", actual,
			)
		}
	}

	sHalo := "haloooo"
	s := inputValidation(sHalo)
	actual := s.validate()
	if actual != nil {
		t.Error(
			"For", sHalo,
			"expected", nil,
			"got", actual,
		)
	}
}
