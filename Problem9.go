package main

import "fmt"

func Problem9() {
	a, b, c, max := 1, 1, 1, 1000

	for a = 1; a < max; a++ {
		tmpB := max - a
		for b = a; b < tmpB; b++ {
			c = max - a - b
			if a*a+b*b == c*c {
				fmt.Printf("%d + %d = %d. The sum of %d, %d and %d is %d.", a*a, b*b, c*c, a, b, c, max)
				return
			}
		}
	}
}