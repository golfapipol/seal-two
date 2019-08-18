package fizzbuzz

import "fmt"

func FizzBuzz(number int) string {
	mod := map[int]string{
		15: "FizzBuzz",
		5:  "Buzz",
		3:  "Fizz",
	}
	for index := range mod {
		if number%index == 0 {
			return mod[index]
		}
	}
	return fmt.Sprintf("%d", number)
}
