package convert

// ArabicToRoman converts an Arabic number to a Roman numeral.
func ArabicToRoman(ar int) string {
	result := ""

	var romanMap = []struct {
		Arabic int
		Roman  string
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
		for ar >= entry.Arabic {
			result += entry.Roman
			ar -= entry.Arabic
		}
	}

	return result
}
