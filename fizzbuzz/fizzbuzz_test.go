package fizzbuzz_test

import (
	. "fizzbuzz"
	"testing"
)

func Test_FizzBuzz_Input_2_Should_Be_2(t *testing.T) {
	expected := "2"
	input := 2

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_FizzBuzz_Input_3_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	input := 3

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_FizzBuzz_Input_4_Should_Be_4(t *testing.T) {
	expected := "4"
	input := 4

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_FizzBuzz_Input_5_Should_Be_Buzz(t *testing.T) {
	expected := "Buzz"
	input := 5

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_FizzBuzz_Input_6_Should_Be_Fizz(t *testing.T) {
	expected := "Fizz"
	input := 6

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_FizzBuzz_Input_10_Should_Be_Fizz(t *testing.T) {
	expected := "Buzz"
	input := 10

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}

func Test_FizzBuzz_Input_15_Should_Be_FizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	input := 15

	actual := FizzBuzz(input)

	if expected != actual {
		t.Errorf("expected %s but it got %s", expected, actual)
	}
}
