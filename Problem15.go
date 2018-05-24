package main

import (
	"fmt"
	"math/big"
)

func Problem15() {
	// Number of lattice paths using the binomial coefficient.

	var nFac, kFac, kFacMul, result big.Int
	nFac.MulRange(1, 40)        // n!
	kFac.MulRange(1, 20)        // k!
	kFacMul.Mul(&kFac, &kFac)   // k!**2
	result.Div(&nFac, &kFacMul) // n! / k! * k!
	fmt.Println(&result)
}