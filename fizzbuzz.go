package fizzbuzz

import (
	"strconv"
)

const (
	FizzDiv  = 3
	FizzName = "Fizz"
	BuzzDiv  = 5
	BuzzName = "Buzz"
)

func FizzBuzz(number int) (accu string) {
	if isNeitherFizzOrBuzz(number) {
		return strconv.Itoa(number)
	}
	return Fizz(number) + Buzz(number)
}

func Fizz(number int) string {
	if isFizz(number) {
		return FizzName
	}
	return ""
}

func Buzz(number int) string {
	if isBuzz(number) {
		return BuzzName
	}
	return ""
}

func isNeitherFizzOrBuzz(number int) bool {
	return !isFizz(number) && !isBuzz(number)
}

func isFizz(number int) bool {
	return (number % FizzDiv) == 0
}

func isBuzz(number int) bool {
	return (number % BuzzDiv) == 0
}
