package main

import (
	"bufio"
	"calculator/convert"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите пример")
	example, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	result := calculator(example)
	fmt.Println(result)

}

// фунция проверки правильности римских значений
func romanTrue(slice []string, item string) bool {

	for _, entry := range slice {
		for item == entry {
			return true
		}
	}
	return false
}

func calculator(example string) (result_example string) {
	//Обработаем строчку
	care_example := strings.ToUpper(strings.TrimSpace(example))
	// Сделаем регулярное выражение
	re := regexp.MustCompile(`\s*([+-/*])\s*`)
	parts := re.Split(care_example, -1)
	operators := re.FindAllString(care_example, -1)

	//Проверяем формат
	if len(parts) < 2 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	if parts[0] == "" || parts[1] == "" {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	if (len(parts) != 2) || (len(operators) != 1) {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию.")
	}

	num1, num2 := parts[0], parts[1]
	oper := strings.TrimSpace(operators[0])

	var err1, err2 error
	var x1, x2 int

	//Пробуем изменить тип на int
	x1, err1 = strconv.Atoi(num1)
	x2, err2 = strconv.Atoi(num2)

	if err1 == nil && err2 == nil {
		if x1 > 10 || x1 < 1 || x2 > 10 || x2 < 1 {
			panic("Выдача паники, так как диапазон доступных чисел: от 1 до 10.")
		}
	}

	//Список доступных римских значений
	romanTrueMap := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI",
		"XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX"}

	isRoman := false

	//Сценарий для римский чисел
	if err1 != nil && err2 != nil {
		if romanTrue(romanTrueMap, num1) && romanTrue(romanTrueMap, num2) {
			x1, x2 = convert.RomanToArabic(num1), convert.RomanToArabic(num2)
			isRoman = true
			if x1 > 10 || x1 < 1 || x2 > 10 || x2 < 1 {
				panic("Выдача паники, так как диапазон доступных чисел: от I до X.")
			}
			if x1 < x2 && oper == "-" {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
			}
			if (x1 == x2 && oper == "-") || (x1 < x2 && oper == "/") {
				panic("Выдача паники, так как в римской системе нет нуля.")
			}

		} else {
			panic("Выдача паники, так как используются неправильные римские значения.")
		}
	}

	if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	var result int

	switch oper {
	case "+":
		result = x1 + x2
	case "-":
		result = x1 - x2
	case "*":
		result = x1 * x2
	case "/":
		result = x1 / x2
	default:
		panic("Неизвестный оператор")
	}

	if isRoman {
		result_example = convert.ArabicToRoman(result)
		return result_example
	} else {
		result_example = strconv.Itoa(result)
	}

	return result_example
}
