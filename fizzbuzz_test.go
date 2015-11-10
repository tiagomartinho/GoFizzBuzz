package fizzbuzz

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNoFizzBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(1), "1")
	assert.Equal(t, FizzBuzz(2), "2")
}

func TestFizz(t *testing.T) {
	assert.Equal(t, FizzBuzz(3), "Fizz")
	assert.Equal(t, FizzBuzz(6), "Fizz")
}

func TestBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(5), "Buzz")
	assert.Equal(t, FizzBuzz(25), "Buzz")
}

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(15), "FizzBuzz")
}
