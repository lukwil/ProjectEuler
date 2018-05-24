package main

import (
	"fmt"
	"math"
)

func Problem12() {
	i := 1
	for {
		sum := 0
		for j := 1; j <= i; j++ {
			sum += j
		}
		divisors := 0
		for j := 1; j <= int(math.Sqrt(float64(sum))); j++ {
			if sum%j == 0 {
				divisors += 2

			}
		}
		// Correction if the number is a perfect square
		if int(math.Sqrt(float64(sum)))*int(math.Sqrt(float64(sum))) == sum {
			divisors--
		}

		if divisors > 500 {
			fmt.Println(sum)
			return
		}
		i++
	}
}