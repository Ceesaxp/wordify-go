package main

import (
	"fmt"
	"testing"
)

func TestPowerName(t *testing.T) {
	args := []struct {
		pw       int
		lang     string
		expected string
	}{
		{0, "en", ""},
		{1, "ru", "тысяча"},
		//{0, "fr", "Language fr is not defined."},
		{3, "en", "billion"},
		{3, "ru", "миллиард"},
		{99, "en", "The power 99 is too big."},
	}

	for _, arg := range args {
		r := PowerName(arg.pw, arg.lang)
		if r != arg.expected {
			t.Errorf("got %q, wanted %q", r, arg.expected)
		}
	}
}

func TestHundreds(t *testing.T) {
	args := []struct {
		lang     string
		expected string
	}{
		{"en", "hundred"},
		{"ru", "сто"},
		{"gr", "Language gr is not defined."},
		{"cn", "Language cn is not defined."},
	}

	for _, arg := range args {
		r := Hundreds(arg.lang)
		if r != arg.expected {
			t.Errorf("got %q, wanted %q", r, arg.expected)
		}
	}
}

func TestAmountToWords(t *testing.T) {
	args := []struct {
		amount    float64
		precision int
		expected  []string
	}{
		{0, 0, []string{"zero ", "dollars"}},
		{0, 1, []string{"zero ", "dollars"}},
		{0, 2, []string{"zero ", "dollars"}},
		{0, 3, []string{"zero ", "dollars"}},
		{0, 4, []string{"precision must be nore than than 3"}},
		{0, -1, []string{"precision must be positive"}},
		{1, 0, []string{"one", "dollars"}},
		{1, 1, []string{"one", "dollars"}},
		{1, 2, []string{"one", "dollars"}},
		{10, 0, []string{"ten", "dollars"}},
		{10, 1, []string{"ten", "dollars"}},
		{10, 2, []string{"ten", "dollars"}},
		{100, 0, []string{"one hundred", "dollars"}},
		{100, 1, []string{"one hundred", "dollars"}},
		{100, 2, []string{"one hundred", "dollars"}},
		{1000, 0, []string{"one thousand", "dollars"}},
		{12345, 0, []string{"twelve thousand", "three hundred", "forty five", "dollars"}},
	}

	for _, arg := range args {
		r, _ := AmountToWords(arg.amount, arg.precision)
		fmt.Print(r)
		if r[0] != arg.expected[0] {
			t.Errorf("got %q, wanted %q", r, arg.expected)
		}
	}
}
