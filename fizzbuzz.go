package fizz_buzz_kata

func FizzBuzz(number int) string {
	if isFizzBuzz(number) {
		return "FizzBuzz"
	} else if isBuzz(number) {
		return "Buzz"
	} else if number%3 == 0 {
		return "Fizz"
	} else {
		return ""
	}
}
func isBuzz(number int) bool {
	if number%5 == 0 {
		return true
	} else {
		return false
	}
}
func isFizzBuzz(number int) bool {
	if number%3 == 0 && number%5 == 0 {
		return true
	} else {
		return false
	}
}
