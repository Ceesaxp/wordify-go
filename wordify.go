package main

import (
	"fmt"
)

/*
PowerName : Returns names for powers of 10^3
Parameters:
  - pw (int) – the power of 10^3
  - lang (string) – the language to use

Returns:
  - string – the name of the power

#FIXME: move language definitions to a separate file
*/
func PowerName(pw int, lang string) string {
	switch lang {
	case "en", "ru":
		powersEn := []string{"", "thousand", "million", "billion", "trillion",
			"quadrillion", "quintillion", "sextillion", "septillion", "octillion",
			"nonillion", "decillion", "undecillion", "duodecillion", "tredecillion",
			"quattuordecillion", "quindecillion", "sexdecillion", "septendecillion",
			"octodecillion", "novemdecillion", "vigintillion"}
		powersRu := []string{"", "тысяча", "миллион", "миллиард", "триллион", "квадриллион",
			"квинтиллион", "секстиллион", "септиллион", "октиллион", "нониллион",
			"дециллион", "ундециллион", "дуодециллион", "тредециллион", "кваттуордециллион",
			"квиндециллион", "сексдециллион", "септендециллион", "октодециллион",
			"новемдециллион", "вигинтиллион"}

		powersLang := map[string][]string{
			"en": powersEn,
			"ru": powersRu,
		}

		if pw < len(powersLang["en"]) {
			return powersLang[lang][pw]
		} else {
			return "The power " + fmt.Sprint(pw) + " is too big."
		}

	default:
		return "Language " + lang + " is not defined."
	}
}

func Hundreds(lang string) string {
	hundreds := map[string]string{
		"en": "hundred",
		"ru": "сто",
	}
	if _, key := hundreds[lang]; key {
		return hundreds[lang]
	} else {
		return "Language " + lang + " is not defined."
	}
}

// PowerShift : given a number, shift it to the left by a power of 10^pw
// Parameters: num (int) – the number to be shifted
// pw (int) – the power to shift the number by
// Returns: []int – a tuple of the shifted number and the remainder
func PowerShift(num int, pw int) []int {
	base := num % (10 ^ pw)
	rem := num - base*(10^pw)
	return []int{base, rem}
}

// AmountToWords : Convert a number to words
// Parameters: amount (float64) – the amount to be converted to words
// Returns: string – the amount in words or error if amount is not a number
func AmountToWords(amount float64) (string, error) {
	precision := 2 // may need to move it to a parameter or a global variable
	dec := 0
	pw := 10.0 ^ precision
	whole := 0

	// check if amount has decimal part
	if amount != float64(int(amount)) {
		whole = int(amount)
		dec = int((amount - float64(whole)) * pw)
	} else {
		whole = int(amount)
		dec = 0
	}

	fmt.Println(whole)
	fmt.Println(dec)

	return "whole: " + fmt.Sprint(whole) + ", decimal: " + fmt.Sprint(dec), nil
}

func main() {
	amount, err := AmountToWords(10.0)
	if err != nil {
		panic(err)
	}
	fmt.Println(amount)
}
