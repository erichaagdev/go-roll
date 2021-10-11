package main

import (
	"testing"
)

func TestParseNumberEmpty(t *testing.T) {
	defer shouldPanic(t)
	parseNumber([]rune(""))
}

func TestParseNumberSingleDigit(t *testing.T) {
	want := parseNumber([]rune("1"))
	if string(want) != "1" {
		t.FailNow()
	}
}

func TestParseNumberMultiDigit(t *testing.T) {
	want := parseNumber([]rune("1234"))
	if string(want) != "1234" {
		t.FailNow()
	}
}

func TestParseNumberNonDigit(t *testing.T) {
	defer shouldPanic(t)
	parseNumber([]rune("a"))
}

func TestParseNumberDigitAndNonDigit(t *testing.T) {
	want := parseNumber([]rune("1a"))
	if string(want) != "1" {
		t.FailNow()
	}
}

func TestParseDiceEmpty(t *testing.T) {
	defer shouldPanic(t)
	parseNumber([]rune(""))
}

func TestParseDiceSingleDigit(t *testing.T) {
	want := parseDice([]rune("d8"))
	if string(want) != "8" {
		t.FailNow()
	}
}

func TestParseDiceMultiDigit(t *testing.T) {
	want := parseDice([]rune("d20"))
	if string(want) != "20" {
		t.FailNow()
	}
}

func TestParseDiceUnexpectedChar(t *testing.T) {
	defer shouldPanic(t)
	parseDice([]rune("a"))
}

func TestParseDiceUnexpectedChar2(t *testing.T) {
	defer shouldPanic(t)
	parseDice([]rune("da20"))
}

func TestParseDiceDigitAndNonDigit(t *testing.T) {
	want := parseDice([]rune("d20a"))
	if string(want) != "20" {
		t.FailNow()
	}
}

func TestParseOperatorPlus(t *testing.T) {
	want := parseOperator([]rune("+"))
	if string(want) != "+" {
		t.FailNow()
	}
}

func TestParseOperatorMinus(t *testing.T) {
	want := parseOperator([]rune("-"))
	if string(want) != "-" {
		t.FailNow()
	}
}

func TestParseOperatorUnexpected(t *testing.T) {
	defer shouldPanic(t)
	parseOperator([]rune("a"))
}

func TestParseSyntax(t *testing.T) {
	want := parseSyntax([]rune("1d2"))
	if want < 1 || want > 2 {
		t.FailNow()
	}
}

func TestParseSyntaxPlus(t *testing.T) {
	want := parseSyntax([]rune("1d2+1"))
	if want < 2 || want > 3 {
		t.FailNow()
	}
}

func TestParseSyntaxMinus(t *testing.T) {
	want := parseSyntax([]rune("1d2-1"))
	if want < 0 || want > 1 {
		t.FailNow()
	}
}

func TestParseSyntaxUnexpected(t *testing.T) {
	defer shouldPanic(t)
	parseSyntax([]rune("1d2a"))
}

func TestParseSyntaxExpectedNumber(t *testing.T) {
	defer shouldPanic(t)
	parseSyntax([]rune("1d2+"))
}

func shouldPanic(t *testing.T) {
	if recover() == nil {
		t.Errorf("The code did not panic")
	}
}
