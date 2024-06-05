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
		panic("Неверный формат")
	}

	// Создаем массив для хранения результата
	num1, num2 := parts[0], parts[1]
	operator := strings.TrimSpace(signs[0])

	var x1, x2 int

	// Проверка формата и преобразование
	x1, err1 := strconv.Atoi(num1)
	x2, err2 := strconv.Atoi(num2)

	isRoman := false

	if err1 != nil || err2 != nil {
		// Если хотя бы одна из ошибок не nil, значит числа могут быть римскими
		x1 = convert.RomanToArabic(num1)
		x2 = convert.RomanToArabic(num2)
		isRoman = true
	}

	if err1 == nil && err2 == nil {
		// Если оба числа арабские, проверяем диапазон
		if (x1 <= 0 || x1 > 10) || (x2 <= 0 || x2 > 10) {
			panic("Диапазон доступных чисел: от 1 до 10")
		}
	} else if isRoman {
		// Если оба числа римские, проверяем диапазон после преобразования
		if (x1 <= 0 || x1 > 10) || (x2 <= 0 || x2 > 10) {
			panic("Диапазон доступных чисел: от I до X")
		}
	} else {
		// Если форматы чисел не совпадают, возвращаем ошибку
		panic("Числа могут быть только одинакового формата!")
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
			panic("На ноль делить нельзя!")
		}
	default:
		panic("Неизвестная операция")
	}

	// Возврат результата в римских или арабских цифрах в зависимости от входных данных
	if isRoman {
		return convert.ArabicToRoman(result)
	}

	return strconv.Itoa(result)
}
