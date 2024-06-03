package main

import (
	"fmt"
)

func ArabicToRoman(ar int) (ro string) {

	result := ""

	var romanMap = []struct {
		Arab int
		Roma string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, entry := range romanMap {
		if ar >= entry.Arab {
			result += entry.Roma
			ar -= entry.Arab
		}
	}

	return result
}

func main() {
	var numb int
	fmt.Println("Введите число для перевода на римский")
	fmt.Scanln(&numb)

	if numb <= 0 {
		fmt.Println("Введите число больше 0")
		return
	}

	translait := ArabicToRoman(numb)
	fmt.Println(numb, "на римском будет", translait)
}
