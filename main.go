package main

import (
	"Kata-Academy-Test/convert"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите пример")
	// Считываем строку через bufio с пробелами
	primerStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// Отправляем на обработку
	calcResult := PrimaryProcessing(primerStr)
	fmt.Println(calcResult)
}

/*
Эта функция отвечает за проверку токенов римских чисел,
чтобы не допустить ввода неверного формата римского числа,
например "IIII" или "VV".
Так как по условия задачи на вход доступны только числа от 1 до 10,
можно обойтись циклом, который сравнивает список допустимых чисел
с входными данными.
*/
func contains(slice []string, item string) bool {
	for _, str := range slice {
		if str == item {
			return true
		}
	}
	return false
}

func PrimaryProcessing(primerStr string) string {
	// Удаляем пробелы и переводим в верхний регистр
	primerStr = strings.ToUpper(strings.TrimSpace(primerStr))
	// Регулярное выражение для поиска знаков +, -, *, / и удаления пробелов вокруг оператора
	re := regexp.MustCompile(`\s*([+-/*])\s*`)
	// Найти все части строки, разделенные знаками
	parts := re.Split(primerStr, -1)
	// Найти все знаки в строке
	signs := re.FindAllString(primerStr, -1)

	// Проверка на валидный формат
	if len(parts) != 2 || len(signs) != 1 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}

	// Создаем массив для хранения результата
	num1, num2 := parts[0], parts[1]
	operator := strings.TrimSpace(signs[0])

	var x1, x2 int
	var err1, err2 error

	// Проверка формата и преобразование
	x1, err1 = strconv.Atoi(num1)
	x2, err2 = strconv.Atoi(num2)

	// Допустимые римские числа (все числа, которые можно составить из "I", "V" и "X")
	validRomanNumerals := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI",
		"XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX"}

	isRoman := false

	if err1 != nil || err2 != nil {
		// Если хотя бы одна из ошибок не nil, значит числа могут быть римскими
		if err1 != nil && err2 != nil {
			// Проверяем, являются ли оба числа допустимыми римскими числами
			if contains(validRomanNumerals, num1) && contains(validRomanNumerals, num2) {
				x1 = convert.RomanToArabic(num1)
				x2 = convert.RomanToArabic(num2)
				isRoman = true
			} else {
				panic("Выдача паники, так как используются неверные римские числа.")
			}
		} else {
			panic("Выдача паники, так как используются одновременно разные системы счисления")
		}
	}

	if err1 == nil && err2 == nil {
		// Если оба числа арабские, проверяем диапазон
		if (x1 <= 0 || x1 > 10) || (x2 <= 0 || x2 > 10) {
			panic("Выдача паники, так как диапазон доступных чисел: от 1 до 10")
		}
	} else if isRoman {
		// Если оба числа римские, проверяем диапазон после преобразования
		if (x1 <= 0 || x1 > 10) || (x2 <= 0 || x2 > 10) {
			panic("Выдача паники, так как диапазон доступных чисел: от I до X")
		}
		if (x1 < x2) && operator == "-" {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
	} else {
		// Если форматы чисел не совпадают, возвращаем ошибку
		panic("Выдача паники, так как используются одновременно разные системы счисления")
	}

	// Выполнение арифметической операции
	var result int
	switch operator {
	case "+":
		result = x1 + x2
	case "-":
		result = x1 - x2
	case "*":
		result = x1 * x2
	case "/":
		if x2 != 0 {
			result = x1 / x2 // Деление, отбросив остаток
		} else {
			panic("Выдача паники, так как на ноль делить нельзя!")
		}
	default:
		panic("Выдача паники, так как строка не является математической операцией.")
	}

	// Возврат результата в римских или арабских цифрах в зависимости от входных данных
	if isRoman {
		return convert.ArabicToRoman(result)
	}

	return strconv.Itoa(result)
}
