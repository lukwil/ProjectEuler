package main

import "fmt"

func Problem3(){
	var number int64 = 600851475143
	var result int64
	var i int64
	for i = 2; i < number/2; i++ {
		// found factor
		if number%i == 0 {
			result = number / i
			// found largest prime factor
			if isPrime(result) {
				break
			}
		}
	}
	fmt.Println("Result: ", result)
}

func isPrime(number int64) bool {
	var i int64
	for i = 2; i < (number / 2); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
