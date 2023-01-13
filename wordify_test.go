package main

import "testing"

func TestPowerName(t *testing.T) {
	args := []struct {
		pw       int
		lang     string
		expected string
	}{
		{0, "en", ""},
		{1, "ru", "тысяча"},
		{0, "fr", "Language fr is not defined."},
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
		{"fr", "Language fr is not defined."},
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
		amount   float64
		expected string
	}{
		{123.00, "whole: 123, decimal: 0"},
		{123.01, "whole: 123, decimal: 1"},
		{123.10, "whole: 123, decimal: 10"},
		{0.00, "whole: 0, decimal: 0"},
	}

	for _, arg := range args {
		r, _ := AmountToWords(arg.amount)
		if r != arg.expected {
			t.Errorf("got %q, wanted %q", r, arg.expected)
		}
	}
}
