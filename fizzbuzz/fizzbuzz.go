package fizzbuzz

import "fmt"

func FizzBuzz(number int) string {
	words := []string{"FizzBuzz", "Buzz", "Fizz"}
	mod := []int{15, 5, 3}
	for index := range mod {
		if number%mod[index] == 0 {
			return words[index]
		}
	}
	return fmt.Sprintf("%d", number)
}
