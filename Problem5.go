package main

import "fmt"

func Problem5() {
	i := 20
	for {
		for j := 1; j < 20; j++ {
			if i%j != 0 {
				break
			}
			if j == 19 {
				fmt.Printf("Found it: %d!\n", i)
				return
			}
		}
		i += 20
	}
}