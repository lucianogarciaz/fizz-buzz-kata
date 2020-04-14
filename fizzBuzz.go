package main

type fizzBuzz struct{}

func NewFizzBuzz() *fizzBuzz {
	return &fizzBuzz{}
}

func (fizzBuzz *fizzBuzz) isValid(number int) bool {
	if number%3 == 0 && number%5 == 0 {
		return true
	}
	return false
}
func (fizzBuzz *fizzBuzz) GetName() string {
	return "FizzBuzz"
}
