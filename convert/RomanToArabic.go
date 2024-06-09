package convert

func RomanToArabic(roman string) int {

	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	total := 0
	prev := 0

	for _, entry := range roman {
		value := romanMap[entry]
		if value > prev {
			total += value - 2*prev
		} else {
			total += value
		}
		prev = value
	}

	return total
}
