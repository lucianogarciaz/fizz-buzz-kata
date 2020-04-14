package main

type buzz struct{}

func NewBuzz() *buzz {
	return &buzz{}
}

func (buzz *buzz) isValid(number int) bool {
	if number%5 == 0 {
		return true
	}
	return false
}
func (buzz *buzz) GetName() string {
	return "Buzz"
}
