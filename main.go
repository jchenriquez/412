package main

import (
	"fmt"
	"math"
)

func fizz(result string, n int) string {
	if math.Mod(float64(n), 3) == 0 {
		return fmt.Sprintf("%s%s", result, "Fizz")
	} else {
		return result
	}
}

func buzz(result string, n int) string {
	if math.Mod(float64(n), 5) == 0 {
		return fmt.Sprintf("%s%s", result, "Buzz")
	} else {
		return result
	}
}

func fizzBuzz(n int) []string {

	if n < 0 {
		return nil
	}

	result := make([]string, 0, n)

	for i := 1; i <= n; i++ {

		trans := buzz(fizz("", i), i)

		if len(trans) == 0 {
			result = append(result, fmt.Sprintf("%d", i))
		} else {
			result = append(result, trans)
		}
	}

	return result

}

func main() {
	fmt.Printf("%v\n", fizzBuzz(15))
}
