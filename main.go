package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanMap = map[string]int{
	"I":        1,
	"II":       2,
	"III":      3,
	"IV":       4,
	"V":        5,
	"VI":       6,
	"VII":      7,
	"VIII":     8,
	"IX":       9,
	"X":        10,
	"XI":       11,
	"XII":      12,
	"XIII":     13,
	"XIV":      14,
	"XV":       15,
	"XVI":      16,
	"XVII":     17,
	"XVIII":    18,
	"XIX":      19,
	"XX":       20,
	"XXI":      21,
	"XXII":     22,
	"XXIII":    23,
	"XXIV":     24,
	"XXV":      25,
	"XXVI":     26,
	"XXVII":    27,
	"XXVIII":   28,
	"XXIX":     29,
	"XXX":      30,
	"XXXI":     31,
	"XXXII":    32,
	"XXXIII":   33,
	"XXXIV":    34,
	"XXXV":     35,
	"XXXVI":    36,
	"XXXVII":   37,
	"XXXVIII":  38,
	"XXXIX":    39,
	"XL":       40,
	"XLI":      41,
	"XLII":     42,
	"XLIII":    43,
	"XLIV":     44,
	"XLV":      45,
	"XLVI":     46,
	"XLVII":    47,
	"XLVIII":   48,
	"XLIX":     49,
	"L":        50,
	"LI":       51,
	"LII":      52,
	"LIII":     53,
	"LIV":      54,
	"LV":       55,
	"LVI":      56,
	"LVII":     57,
	"LVIII":    58,
	"LVIX":     59,
	"LX":       60,
	"LXI":      61,
	"LXII":     62,
	"LXIII":    63,
	"LXIV":     64,
	"LXV":      65,
	"LXVI":     66,
	"LXVII":    67,
	"LXVIII":   68,
	"LXIX":     69,
	"LXX":      70,
	"LXXI":     71,
	"LXXII":    72,
	"LXXIII":   73,
	"LXXIV":    74,
	"LXXV":     75,
	"LXXVI":    76,
	"LXXVII":   77,
	"LXXVIII":  78,
	"LXXIX":    79,
	"LXXX":     80,
	"LXXXI":    81,
	"LXXXII":   82,
	"LXXXIII":  83,
	"LXXXIV":   84,
	"LXXXV":    85,
	"LXXXVI":   86,
	"LXXXVII":  87,
	"LXXXVIII": 88,
	"LXXXIX":   89,
	"XC":       90,
	"XCI":      91,
	"XCII":     92,
	"XCIII":    93,
	"XCIV":     94,
	"XCV":      95,
	"XCVI":     96,
	"XCVII":    97,
	"XCVIII":   98,
	"XCIX":     99,
	"C":        100,
}

var arabicMap = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	11:  "XI",
	12:  "XII",
	13:  "XIII",
	14:  "XIV",
	15:  "XV",
	16:  "XVI",
	17:  "XVII",
	18:  "XVIII",
	19:  "XIX",
	20:  "XX",
	21:  "XXI",
	22:  "XXII",
	23:  "XXIII",
	24:  "XXIV",
	25:  "XXV",
	26:  "XXVI",
	27:  "XXVII",
	28:  "XXVIII",
	29:  "XXIX",
	30:  "XXX",
	31:  "XXXI",
	32:  "XXXII",
	33:  "XXXIII",
	34:  "XXXIV",
	35:  "XXXV",
	36:  "XXXVI",
	37:  "XXXVII",
	38:  "XXXVIII",
	39:  "XXXIX",
	40:  "XL",
	41:  "XLI",
	42:  "XLII",
	43:  "XLIII",
	44:  "XLIV",
	45:  "XLV",
	46:  "XLVI",
	47:  "XLVII",
	48:  "XLVIII",
	49:  "XLIX",
	50:  "L",
	51:  "LI",
	52:  "LII",
	53:  "LIII",
	54:  "LIV",
	55:  "LV",
	56:  "LVI",
	57:  "LVII",
	58:  "LVIII",
	59:  "LVIX",
	60:  "LX",
	61:  "LXI",
	62:  "LXII",
	63:  "LXIII",
	64:  "LXIV",
	65:  "LXV",
	66:  "LXVI",
	67:  "LXVII",
	68:  "LXVIII",
	69:  "LXIX",
	70:  "LXX",
	71:  "LXXI",
	72:  "LXXII",
	73:  "LXXIII",
	74:  "LXXIV",
	75:  "LXXV",
	76:  "LXXVI",
	77:  "LXXVII",
	78:  "LXXVIII",
	79:  "LXXIX",
	80:  "LXXX",
	81:  "LXXXI",
	82:  "LXXXII",
	83:  "LXXXIII",
	84:  "LXXXIV",
	85:  "LXXXV",
	86:  "LXXXVI",
	87:  "LXXXVII",
	88:  "LXXXVIII",
	89:  "LXXXIX",
	90:  "XC",
	91:  "XCI",
	92:  "XCII",
	93:  "XCIII",
	94:  "XCIV",
	95:  "XCV",
	96:  "XCVI",
	97:  "XCVII",
	98:  "XCVIII",
	99:  "XCIX",
	100: "C",
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
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return 0
		}
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
			panic("Некорректный ввод.")
		}

		var num1, num2 int
		var err error

		if isRomanNumeral(inputParts[0]) && isRomanNumeral(inputParts[2]) {
			num1, err = romanToArabic(inputParts[0])
			num2, err = romanToArabic(inputParts[2])
			if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				panic("Неверное число")
			}
		} else {
			num1, err = strconv.Atoi(inputParts[0])
			num2, err = strconv.Atoi(inputParts[2])
			if err != nil || num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				panic("Только от 1 до 10")
			}
		}

		result := calculate(num1, num2, inputParts[1])

		if isRomanNumeral(inputParts[0]) && isRomanNumeral(inputParts[2]) {
			if result <= 0 {
				panic("Результат не может быть нулевым или отрицательным.")
			}
			fmt.Printf("Результат: %s\n", arabicMap[result])
		} else {
			fmt.Printf("Результат: %d\n", result)
		}
	}
}
