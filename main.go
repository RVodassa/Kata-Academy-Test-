package main

import (
	"Kata-Academy-Test/convert"
	"fmt"
)

func main() {
	var choice int
	fmt.Println("Выберите действие:")
	fmt.Println("1. Преобразовать арабское число в римское")
	fmt.Println("2. Преобразовать римское число в арабское")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var numb int
		fmt.Println("Введите арабское число больше 0:")
		fmt.Scanln(&numb)

		if numb <= 0 {
			fmt.Println("Число должно быть больше 0")
			return
		}

		translation := convert.ArabicToRoman(numb)
		fmt.Printf("Арабское число: %d, Римское число: %s\n", numb, translation)

	case 2:
		var romanNum string
		fmt.Println("Введите римское число:")
		fmt.Scanln(&romanNum)

		translation := convert.RomanToArabic(romanNum)
		fmt.Printf("Римское число: %s, Арабское число: %d\n", romanNum, translation)

	default:
		fmt.Println("Неверный выбор")
	}
}
