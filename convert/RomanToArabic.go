package convert

// RomanToArabic converts a Roman numeral to an Arabic number.
func RomanToArabic(ro string) int {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prev := 0

	for _, char := range ro {
		value := romanMap[char]
		if value > prev {
			total += value - 2*prev
		} else {
			total += value
		}
		prev = value
	}

	return total
}
