package main

type fizz struct{}

func NewFizz() *fizz {
	return &fizz{}
}

func (fizz *fizz) isValid(number int) bool {
	if number%3 == 0 {
		return true
	}
	return false
}
func (fizz *fizz) GetName() string {
	return "Fizz"
}
