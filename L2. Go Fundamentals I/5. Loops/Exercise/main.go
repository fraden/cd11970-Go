package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Fizz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

func main() {
	fmt.Println(fizzbuzz(20))

}
