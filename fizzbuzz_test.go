package main

import "testing"

type dataTest struct {
	input    int
	expected interface{}
}

func Test_FizzBuzzNames(t *testing.T) {
	for _, testData := range setUpData() {
		object := Factory(testData.input)
		if object != nil && object.GetName() != testData.expected {
			t.Error("The expected value was: ", testData.expected, " instead ", object.GetName())
		}
	}
}
func setUpData() []dataTest {
	return []dataTest{
		{1, 1},
		{2, 2},
		{3, "Fizz"},
		{4, 4},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, 7},
		{8, 8},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, 11},
		{12, "Fizz"},
		{13, 13},
		{14, 14},
		{15, "FizzBuzz"},
		{16, 16},
		{17, 17},
		{18, "Fizz"},
		{19, 19},
		{20, "Buzz"},
		{21, "Fizz"},
		{22, 22},
		{23, 23},
		{24, "Fizz"},
		{25, "Buzz"},
		{26, 26},
		{27, "Fizz"},
		{28, 28},
		{29, 29},
		{30, "FizzBuzz"},
	}
}
