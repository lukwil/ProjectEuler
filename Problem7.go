package main

import (
	"fmt"
	"math"
)

func Problem7() {
	n := 1
	i := 3
	for {
		if isPrime(i) {
			n++
		}
		if n == 10001 {
			fmt.Println(i)
			break
		}
		i++
	}
}

func isPrime(number int) bool {
	//var i int64 = 2
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}