package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"calculator/convert"
)

func main() {
	var example string
	fmt.Println("Введите пример")
	example, _ = bufio.NewReader(os.Stdin).ReadString('\n') //Используем bufio для получения входных данных
	result := calc(example)
	fmt.Println(result)
}

// Функция проверки римского числа на правильность ввода
func trueRoman(slice []string, item string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}
	return false
}

// Функция принимает строку, и вовзращает строку, потому-что результатом может римское число.
func calc(example string) (result_example string) {
	// Обрабатываем, проверяем и добываеи нужную инфу из строки
	example_care := strings.ToUpper(strings.TrimSpace(example))
	re := regexp.MustCompile(`\s*([+-/*])\s*`)
	opernands := re.Split(example_care, -1)
	operators := re.FindAllString(example_care, -1)
	if len(opernands) < 2 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	//создаем переменные (первое число, второе, оператор)
	num1, num2 := opernands[0], opernands[1]
	operator := strings.TrimSpace(operators[0])
	//Проверяем на кол.во операторов и опернандов

	if len(opernands) != 2 || len(operators) != 1 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию.")
	}

	if num1 == "" || num2 == "" {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	//Создаем переменные для хранения чисел и ошибок. Используем ошибки для обработки дальнейших сценариев.
	var err1, err2 error
	var x1, x2 int

	// Пробуем перевести опернанлды в int
	x1, err1 = strconv.Atoi(num1)
	x2, err2 = strconv.Atoi(num2)

	// Список римских чисел, которые можно составить из "I", "V" и "X"
	romanMap := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI",
		"XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX"}

	// Флаг, чтобы мы могли вывести результат в нужно формате.
	isRoman := false

	// Если перевести в int с помощью strconv.Atoi не удалось, пробуем перевести как римскую строку, используя нашу функцию RomanToArabic
	// и проверяем на допустимость значений.
	if err1 != nil && err2 != nil {
		if trueRoman(romanMap, num1) && trueRoman(romanMap, num2) {
			x1, x2 = convert.RomanToArabic(num1), convert.RomanToArabic(num2)
			isRoman = true //  Меняем значение флага на true, чтобы идти по римскому сценарию.
			if x1 > 10 || x1 <= 0 || x2 > 10 || x2 <= 0 {
				panic("Выдача паники, так как доступный диапазон от I до X.")
			}
			if x1 < x2 {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
			}
			if (x1 == x2) && (operator == "-") {
				panic("Выдача паники, так как в римской системе нет нуля")
			}
		} else {
			panic("Выдача паники, так как в примере используются неизвестные значения")
		}
	}
	// Если строка перевелась в int c помощью strconv.Atoi без ошибок, првоеряем на допустимость значений.
	if err1 == nil && err2 == nil {
		if x1 > 10 || x1 <= 0 || x2 > 10 || x2 <= 0 {
			panic("Выдача паники, так как доступный диапазон от 1 до 10")
		}
	}
	// Проверяем чтобы числа не были из разных систем счисления.
	if (err1 == nil && err2 != nil) || (err1 != nil && err2 == nil) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
	// Создаем переменную, в которую будем записывать результат математической операции.
	var result int

	switch operator {
	case "+":
		result = x1 + x2
	case "-":
		result = x1 - x2
	case "*":
		result = x1 * x2
	case "/":
		result = x1 / x2
	default:
		panic("Выдача паники, так как испольузется неизвестный оператор")
	}
	// Если флаг isRoman = true, значит на входе были римские числа, значит результатом должно быть римское число,
	if isRoman {
		// переводим число с арабского на римский используя нашу функцию ArabicToRoman
		result_example := convert.ArabicToRoman(result)
		return result_example // Выводим результат вычеслений римских чисел
	} else {
		result_example = strconv.Itoa(result) // если isRoman = false, выводим результат вычислений арабских чисел в формате строки.
	}
	return result_example // Возвращаем результат функции.
}
