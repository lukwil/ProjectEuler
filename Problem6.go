package main

import "fmt"

func Problem6() {
	sumOfSquares := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
	}

	squareOfSum := 0
	for i := 1; i <= 100; i++ {
		squareOfSum += i
	}
	squareOfSum *= squareOfSum
	fmt.Println(squareOfSum - sumOfSquares)
}