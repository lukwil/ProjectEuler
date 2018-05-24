package main

import (
	"fmt"
	"strconv"
)

func Problem4() {
	var result int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product := strconv.Itoa(i * j)
			var front, back string
			if len(product) == 6 {
				front = product[:len(product)-3]
				back = reverse(product[len(product)-3:])
			} else {
				front = product[:len(product)-3]
				back = reverse(product[len(product)-2:])
			}
			if front == back && i*j > result{result = i*j}
		}
	}
	fmt.Println(result)
}

func reverse(s string) string {
    var reverse string
    for i := len(s)-1; i >= 0; i-- {
        reverse += string(s[i])
    }
    return reverse 
}