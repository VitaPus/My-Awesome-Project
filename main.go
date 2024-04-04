package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
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

var arabicMap = map[int]string{
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

func isRomanNumeral(str string) bool {
	for r := range romanMap {
		if r == str {
			return true
		}
	}
	return false
}

func romanToArabic(roman string) (int, error) {
	total := 0
	var prevValue int

	for i := len(roman) - 1; i >= 0; i-- {
		character := string(roman[i])
		value := romanMap[character]

		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}

	return total, nil
}

func calculate(num1, num2 int, operator string) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		return num1 / num2
	default:
		return 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ввод: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Программа завершена.")
			break
		}

		inputParts := strings.Fields(input)

		if len(inputParts) != 3 {
			fmt.Println("Некорректный ввод.")
			continue
		}

		var num1, num2 int
		var err error

		if isRomanNumeral(inputParts[0]) && isRomanNumeral(inputParts[2]) {
			num1, err = romanToArabic(inputParts[0])
			num2, err = romanToArabic(inputParts[2])
		} else {
			num1, err = strconv.Atoi(inputParts[0])
			num2, err = strconv.Atoi(inputParts[2])
		}

		if err != nil {
			fmt.Println("Ошибка преобразования чисел.")
			continue
		}

		result := calculate(num1, num2, inputParts[1])

		if isRomanNumeral(inputParts[0]) && isRomanNumeral(inputParts[2]) {
			fmt.Printf("Результат: %s\n", arabicMap[result])
		} else {
			fmt.Printf("Результат: %d\n", result)
		}
	}
}
