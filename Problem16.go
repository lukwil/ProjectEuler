package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func Problem16() {
	var number big.Int
	two := big.NewInt(2)
	exp := big.NewInt(1000)
	mod := big.NewInt(0)
	number.Exp(two, exp, mod)
	numberString := number.Text(10) // 10 = decimal base
	sumOfDigits := 0
	for _, element := range numberString {
		res, _ := strconv.Atoi(string(element))
		sumOfDigits += res
	}
	fmt.Println(sumOfDigits)
}