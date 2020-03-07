package fizz_buzz_kata

import "testing"

type dataTest struct {
	input    int
	expected string
}

func Test_if_module_three_return_Fizz(t *testing.T) {
	for _, testData := range setUpData() {
		if FizzBuzz(testData.input) != testData.expected {
			t.Error("The expected value was: ", testData.expected, " instead ", testData.input)
		}
	}
}
func setUpData() []dataTest {
	return []dataTest{
		{1, ""},
		{2, ""},
		{3, "Fizz"},
		{4, ""},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, ""},
		{8, ""},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, ""},
		{12, "Fizz"},
		{13, ""},
		{14, ""},
		{15, "FizzBuzz"},
		{16, ""},
		{17, ""},
		{18, "Fizz"},
		{19, ""},
		{20, "Buzz"},
		{21, "Fizz"},
		{22, ""},
		{23, ""},
		{24, "Fizz"},
		{25, "Buzz"},
		{26, ""},
		{27, "Fizz"},
		{28, ""},
		{29, ""},
		{30, "FizzBuzz"},
	}
}
