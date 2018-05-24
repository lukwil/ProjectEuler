package main

import (
	"fmt"
	"math"
)

func Problem10() {
	result := sieveOfEratosthenes(2000000)
	var sum int64
	for _, i := range result{
		sum += i
	}
	fmt.Println(sum)
}

func sieveOfEratosthenes(number int64) []int64 {
	var result []int64
	sieve := make([]bool, number+1)
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(number))); i++ {
		if !sieve[i] {
			result = append(result, i)
			for j := i * i; j <= number; j += i {
				sieve[j] = true
			}
		}
	}
	for i := int64(math.Sqrt(float64(number))) + 1; i <= number; i++ {
		if !sieve[i] {
			result = append(result, i)
		}
	}
	return result
}