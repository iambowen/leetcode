package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	results := []string{""}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			results = append(results, "FizzBuzz")
		} else if i%3 == 0 {
			results = append(results, "Fizz")
		} else if i%5 == 0 {
			results = append(results, "Buzz")
		} else {
			results = append(results, strconv.Itoa(i))
		}

	}
	return results[1:]
}

func main() {
	fmt.Println(fizzBuzz(15))
}
