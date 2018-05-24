package main

import (
	"fmt"
	"unicode/utf8"
)

var result string

func Problem17() {
	number := 999
	var letterCount int
	for i:= 1; i<= number; i++ {
		modHundred(i)
		letterCount += utf8.RuneCountInString(result)
	}
	fmt.Println(letterCount + 11) // 11 = one thousand

}

func modHundred(number int) {
	result = ""
	divided := number / 100
	if divided == 0 {
		modTen(number)
		return
	}
	modSingle(divided)
	result += "hundred"
	if number%100 != 0 {
		result += "and"
	}
	modTen(number % 100)

}

func modTen(number int) {
	divided := number / 10
	if divided == 0 {
		modSingle(number)
		return
	}
	if divided == 1 {
		switch number {
		case 10:
			result += "ten"
		case 11:
			result += "eleven"
		case 12:
			result += "twelve"
		case 13:
			result += "thirteen"
		case 14:
			result += "fourteen"
		case 15:
			result += "fifteen"
		case 16:
			result += "sixteen"
		case 17:
			result += "seventeen"
		case 18:
			result += "eighteen"
		case 19:
			result += "nineteen"
		}
		return
	}
	switch divided {
	case 2:
		result += "twenty"
	case 3:
		result += "thirty"
	case 4:
		result += "forty"
	case 5:
		result += "fifty"
	case 6:
		result += "sixty"
	case 7:
		result += "seventy"
	case 8:
		result += "eighty"
	case 9:
		result += "ninety"
	}
	modSingle(number % 10)

}

func modSingle(number int) {
	switch number {
	case 1:
		result += "one"
	case 2:
		result += "two"
	case 3:
		result += "three"
	case 4:
		result += "four"
	case 5:
		result += "five"
	case 6:
		result += "six"
	case 7:
		result += "seven"
	case 8:
		result += "eight"
	case 9:
		result += "nine"
	}
}