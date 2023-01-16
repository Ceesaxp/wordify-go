package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

var numbersInWordsEN = map[int]string{
	0:   "zero",
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	10:  "ten",
	11:  "eleven",
	12:  "twelve",
	13:  "thirteen",
	14:  "fourteen",
	15:  "fifteen",
	16:  "sixteen",
	17:  "seventeen",
	18:  "eighteen",
	19:  "nineteen",
	20:  "twenty",
	30:  "thirty",
	40:  "forty",
	50:  "fifty",
	60:  "sixty",
	70:  "seventy",
	80:  "eighty",
	90:  "ninety",
	100: "hundred",
	200: "two hundred",
	300: "three hundred",
	400: "four hundred",
	500: "five hundred",
	600: "six hundred",
	700: "seven hundred",
	800: "eight hundred",
	900: "nine hundred",
}

var numbersInWordsRU = map[int]string{
	0:   "ноль",
	1:   "один",
	2:   "два",
	3:   "три",
	4:   "четыре",
	5:   "пять",
	6:   "шесть",
	7:   "семь",
	8:   "восемь",
	9:   "девять",
	10:  "десять",
	11:  "одиннадцать",
	12:  "двенадцать",
	13:  "тринадцать",
	14:  "четырнадцать",
	15:  "пятнадцать",
	16:  "шестнадцать",
	17:  "семнадцать",
	18:  "восемнадцать",
	19:  "девятнадцать",
	20:  "двадцать",
	30:  "тридцать",
	40:  "сорок",
	50:  "пятьдесят",
	60:  "шестьдесят",
	70:  "семьдесят",
	80:  "восемьдесят",
	90:  "девяносто",
	100: "сто",
	200: "двести",
	300: "триста",
	400: "четыреста",
	500: "пятьсот",
	600: "шестьсот",
	700: "семьсот",
	800: "восемьсот",
	900: "девятьсот",
}

type Language struct {
	// Name of the language
	Lang string
	// Words for powers of 10^3
	Powers []string
}

type Options struct {
	// Language of the output
	Language string
	// Currency name
	Currency string
	// Currency name in plural form
	CurrencyPlural string
	// Currency name in plural form for numbers ending with 1
	CurrencyPlural1 string
	// Precision of the output
	Precision int
}

type CliOptions struct {
	// Language of the output
	Language string
	// Amount is currency
	IsCurrency bool
	// Currency name
	Currency string
	// Currency decimals name
	CurrencyDecimals string
}

/*
ReadLanguagesFromFile : Reads languages definitions from a file powers.json, returns Language struct for the given language
*/
func ReadLanguagesFromFile(lang string) (Language, error) {
	var jsonDefinition []Language
	var langStruct Language

	file, err := os.ReadFile("powers.json")
	if err != nil {
		log.Fatal(err)
	}
	// Extract powers definition only for the requested language lang
	err = json.Unmarshal(file, &jsonDefinition)
	if err != nil {
		log.Fatal(err)
	}
	for _, l := range jsonDefinition {
		if l.Lang == lang {
			langStruct = l
			return langStruct, nil
		}
	}
	return langStruct, errors.New("Language " + lang + " is not defined.")
}

/*
PowerName : Returns names for powers of 10^3
Parameters:
  - pw (int) – the power of 10^3
  - lang (string) – the language to use

Returns:
  - string – the name of the power
*/
func PowerName(pw int, lang string) string {
	langStruct, err := ReadLanguagesFromFile(lang)
	if err != nil {
		log.Fatal(err)
	}

	powers := langStruct.Powers
	if pw < len(powers) {
		return powers[pw]
	} else {
		return "The power " + fmt.Sprint(pw) + " is too big."
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
func PowerShift(num int, pw int) (int, int) {
	shiftBy := int(math.Pow(10, float64(pw)))
	reminder := num % shiftBy
	base := num / shiftBy
	return base, reminder
}

// AmountToWords : Convert a number to words
// Parameters: amount (float64) – the amount to be converted to words
// Returns: string – the amount in words or error if amount is not a number
func AmountToWords(amount float64, precision int) ([]string, error) {
	if precision > 3 {
		return nil, errors.New("precision must be nore than than 3")
	}
	if precision < 0 {
		return nil, errors.New("precision must be positive")
	}

	var triplets []int // A slice to hold the triplets of the amount, e.g. 123456789 -> [123, 456, 789]
	whole := int(amount)
	decimal := int(math.Round((amount - float64(whole)) * math.Pow(10, float64(precision))))

	triplets = SplitIntoTriplets(whole)
	var words []string
	p := 0
	for i := len(triplets) - 1; i >= 0; i-- {
		words = append(words, TripletToWords(triplets[i], p))
		p += 1
	}

	// Reverse the words slice
	for i := 0; i < len(words)/2; i++ {
		j := len(words) - i - 1
		words[i], words[j] = words[j], words[i]
	}

	words = append(words, CurrencyName("en"))

	if decimal != 0 {
		words = append(words, "and "+TripletToWords(decimal, 0))
		words = append(words, DecimalsName("en"))
	}

	return words, nil
}

func CurrencyName(lang string) string {
	currency := map[string]string{
		"en": "dollars",
		"ru": "рублей",
	}
	if _, key := currency[lang]; key {
		return currency[lang]
	} else {
		return "Language " + lang + " is not defined."
	}
}

func DecimalsName(lang string) string {
	if lang == "en" {
		return " cents"
	} else {
		return " копеек"
	}
}

func TripletToWords(triplet int, index int) string {
	pwName := PowerName(index, "en") // FIXME: hardcoded language
	if triplet <= 20 {
		return fmt.Sprintf("%s %s", numbersInWordsEN[triplet], pwName)
	} else {
		// Split the triplet into hundreds, tens and units
		hundreds, reminder := PowerShift(triplet, 2)
		tens, units := PowerShift(reminder, 1)
		// Convert the triplet into words
		if hundreds == 0 {
			return fmt.Sprintf("%s %s %s", numbersInWordsEN[tens*10], numbersInWordsEN[units], pwName)
		} else {
			return fmt.Sprintf("%s %s %s %s",
				numbersInWordsEN[hundreds*100], numbersInWordsEN[tens*10], numbersInWordsEN[units], pwName)
		}
	}
}

func SplitIntoTriplets(number int) []int {
	shift := 3 // shift by 3 digits, may need to make a parameter
	if number < int(math.Pow(10, float64(shift))) {
		return []int{number}
	} else {
		base, rem := PowerShift(number, shift)
		return append(SplitIntoTriplets(base), rem)
	}
}

func main() {
	var opts Options
	opts.Language = "en"
	opts.Precision = 2

	result, err := AmountToWords(12_345_678_987_654.12, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
