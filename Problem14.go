package main

import (
	"fmt"
)

func Problem14() {
	maxChainLength := 0
	maxN := 0
	var i int64
	for i = 1; i < 1000000; i++ {
		n := i
		chainLengthTmp := 0

		for n > 1 {
			chainLengthTmp++
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
		}
		if chainLengthTmp > maxChainLength {
			maxChainLength = chainLengthTmp
			maxN = int(i)
		}
	}
	fmt.Println(maxN)
}
