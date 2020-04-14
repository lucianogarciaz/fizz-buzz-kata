package main

func Factory(number int) Interpreter {
	if (&fizzBuzz{}).isValid(number) {
		return NewFizzBuzz()
	}
	if (&fizz{}).isValid(number) {
		return NewFizz()
	}
	if (&buzz{}).isValid(number) {
		return NewBuzz()
	}
	return nil
}
