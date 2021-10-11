package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/*
	<dice_roll> ::= <roll_expression> | <roll_expression> <operator> <number>
	<roll_expression> ::= <number> <dice>
	<number> ::= [1-9] [0-9]*
	<dice> ::= "d" <number>
	<operator> ::= "+" | "-"
*/

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	input := strings.Join(os.Args[1:], "")
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	number := parseSyntax([]rune(input))
	fmt.Println(number)
}

func parseSyntax(input []rune) int {
	result, remainder := parseRollExpression(input)
	if remainder == len(input) {
		return result
	}

	operator := parseOperator(input[remainder:])
	if remainder+1 == len(input) {
		panic("expected number")
	}

	number := parseNumber(input[remainder+1:])
	return doOperator(operator, result, runesToInt(number))
}

func parseRollExpression(input []rune) (result int, consumed int) {
	times := parseNumber(input)
	dice := parseDice(input[len(times):])
	result = roll(runesToInt(times), runesToInt(dice))
	consumed = len(times) + len(dice) + 1
	return
}

func parseNumber(input []rune) (token []rune) {
	if len(input) == 0 || !unicode.IsDigit(input[0]) || runeToInt(input[0]) == 0 {
		panic("expected number: 1-9")
	}
	token = input[:1]
	for i := 1; i < len(input) && unicode.IsDigit(input[i]); i++ {
		token = input[:i+1]
	}
	return
}

func parseDice(input []rune) []rune {
	if len(input) == 0 || input[0] != 'd' {
		panic("expected: 'd'")
	}
	return parseNumber(input[1:])
}

func parseOperator(input []rune) rune {
	if len(input) == 0 || (input[0] != '+' && input[0] != '-') {
		panic("expected operator")
	}
	return input[0]
}

func roll(times int, dice int) (result int) {
	for i := 0; i < times; i++ {
		result += random.Intn(dice) + 1
	}
	return
}

func runesToInt(input []rune) (result int) {
	result, _ = strconv.Atoi(string(input))
	return
}

func runeToInt(input rune) (result int) {
	result, _ = strconv.Atoi(string(input))
	return
}

func doOperator(operator rune, a int, b int) int {
	switch operator {
	case '+':
		return a + b
	case '-':
		return a - b
	default:
		panic("unknown operator")
	}
}
