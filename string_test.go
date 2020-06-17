package main

import "testing"

type testPairFindIndex struct {
	char   string
	result int
}

type testPairGetString struct {
	idx    int
	result string
}

var testsFind = []testPairFindIndex{
	{"0", 0},
	{"5", 5},
	{"9", 9},
}

var testsGet = []testPairGetString{
	{0, "0123456789"},
	{5, "56789"},
	{9, "9"},
}

func TestFindIndexOfChar(t *testing.T) {
	s := inputString("0123456789")

	for _, pair := range testsFind {
		actual := s.findIndexOfChar(pair.char)
		if actual != pair.result {
			t.Error(
				"For", pair.char,
				"expected", pair.result,
				"got", actual,
			)
		}
	}
}

func TestGetStringStartFromIndex(t *testing.T) {
	s := inputString("0123456789")

	for _, pair := range testsGet {
		actual := s.getStringStartFromIndex(pair.idx)
		if actual != pair.result {
			t.Error(
				"For", pair.idx,
				"expected", pair.result,
				"got", actual,
			)
		}
	}
}
