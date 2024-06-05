package go_romawi_to_integer

import (
	"strings"
)

// Fungsi untuk memeriksa 4 karakter berurutan
func checkQuadSymbol(romawi string) error {
	romawiLength := len(romawi)
	if romawiLength > 3 {
		for i := 0; i < romawiLength-3; i++ {
			if romawi[i] == romawi[i+1] && romawi[i] == romawi[i+2] && romawi[i] == romawi[i+3] {
				return ErrorInvalidRomawi
			}
		}
	}
	return nil
}

// Fungsi untuk memeriksa 2 karakter berurutan untuk V, L, dan D
func checkDoubleVLD(romawi string) error {
	romawiLength := len(romawi)
	if romawiLength > 1 {
		for i := 0; i < romawiLength-1; i++ {
			if (romawi[i] == 'V' && romawi[i+1] == 'V') ||
				(romawi[i] == 'L' && romawi[i+1] == 'L') ||
				(romawi[i] == 'D' && romawi[i+1] == 'D') {
				return ErrorInvalidRomawi
			}
		}
	}
	return nil
}

// Fungsi untuk melakukan check string I V X L C D M
func isRomawiValid(romawi string) error {
	validRomawi := "IVXLCDM"
	for _, r := range romawi {
		if !strings.ContainsRune(validRomawi, r) {
			return ErrorInvalidRomawi
		}
	}
	return nil
}

func RomawiToInteger(romawi string) (int, error) {
	length := len(romawi)
	if length == 0 {
		return 0, ErrorInvalidRomawi
	}

	romawi = strings.ToUpper(romawi)

	if err := isRomawiValid(romawi); err != nil {
		return 0, err
	}

	if err := checkQuadSymbol(romawi); err != nil {
		return 0, err
	}

	if err := checkDoubleVLD(romawi); err != nil {
		return 0, err
	}

	romawiDict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i := 0; i < length; i++ {
		if i+1 < length && romawiDict[romawi[i]] < romawiDict[romawi[i+1]] {
			result += romawiDict[romawi[i+1]] - romawiDict[romawi[i]]
			i++
		} else {
			result += romawiDict[romawi[i]]
		}
	}

	return result, nil
}
