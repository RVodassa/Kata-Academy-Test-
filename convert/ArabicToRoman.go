package convert

func ArabicToRoman(ar int) string {

	ArabicMap := []struct {
		Roman  string
		Arabic int
	}{
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	result := ""

	for _, entry := range ArabicMap {
		for ar >= entry.Arabic {
			result += entry.Roman
			ar -= entry.Arabic
		}
	}

	return result
}
