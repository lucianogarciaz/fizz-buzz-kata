package main

type Interpreter interface {
	isValid(number int) bool
	GetName() string
}
