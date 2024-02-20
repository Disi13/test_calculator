package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		panic("Неверное количество аргументов. Введите два числа и операцию.")
	}

	num1, err := parseNumber(args[0])
	if err != nil {
		panic(err)
	}

	num2, err := parseNumber(args[2])
	if err != nil {
		panic(err)
	}

	operator := args[1]

	result := 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			panic("Деление на ноль недопустимо.")
		}
		result = num1 / num2
	default:
		panic("Неверная операция. Допустимые операции: +, -, *, /")
	}

	if isRoman(args[0]) && isRoman(args[2]) {
		romanResult := convertToRoman(result)
		fmt.Println("Результат:", romanResult)
	} else {
		fmt.Println("Результат:", result)
	}
}

func parseNumber(input string) (int, error) {
	num, err := strconv.Atoi(input)
	if err == nil {
		if num < 1 || num > 10 {
			return 0, fmt.Errorf("Недопустимое число. Введите число от 1 до 10.")
		}
		return num, nil
	}

	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	num, ok := romanNumerals[input]
	if !ok {
		return 0, fmt.Errorf("Недопустимое число. Введите число от 1 до 10.")
	}

	return num, nil
}

func isRoman(input string) bool {
	romanNumerals := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}

	_, ok := romanNumerals[input]
	return ok
}

func convertToRoman(num int) string {
	romanNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	return romanNumerals[num]
}
