package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите выражение: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.Contains(input, "\" - \"") {
			reg := regexp.MustCompile(`^"([^"]+)"`)
			match := reg.FindString(input)
			if match != "" {
				input = strings.ReplaceAll(input, match, strings.ReplaceAll(match, " ", ""))
			}
		}
		if strings.ToLower(input) == "exit" {
			fmt.Print("Программа остановлена.")
			break
		}
		result, err := evalute(input)
		if err != nil {
			fmt.Println("Ошибка: ", err)
		} else {
			if len(result) > 40 {
				result = result[:40] + "..."
			}
			fmt.Printf("Результат: \"%s\"\n", result)
		}
	}

}

func evalute(input string) (string, error) {
	tokens := strings.Fields(input)
	if len(tokens) < 3 {
		return "", fmt.Errorf("некорректное выражение")
	}
	if tokens[0][0] != '"' {
		panic("Первый операнд должен быть строкой в кавычках")
	}

	switch tokens[1] {
	case "+":
		if len(tokens[0]) > 12 || len(tokens[2]) > 12 {
			panic("не больше 10 символов")
		}
		if tokens[0][0] == '"' && tokens[2][0] == '"' {
			return strings.Trim(tokens[0], "\"") + strings.Trim(tokens[2], "\""), nil
		}
		panic("некоректные операнды")
	case "-":
		if len(tokens[0]) > 12 || len(tokens[2]) > 12 {
			panic("не больше 10 символов")
		}
		if tokens[0][0] == '"' && tokens[2][0] == '"' {
			return strings.Replace(strings.Trim(tokens[0], "\""), strings.Trim(tokens[2], "\""), "", -1), nil
		}
		panic("некорре операнды для операции вычитания")
	case "*":
		if len(tokens[0]) > 12 {
			panic("длина строки не должна превышать 10 символов")
		}
		n, err := strconv.Atoi(tokens[2])
		if err != nil || n < 1 || n > 10 {
			panic("второй операнд должен быть числом от 1 до 10")
		}
		if tokens[0][0] == '"' {
			return strings.Repeat(strings.Trim(tokens[0], "\""), n), nil
		}
		panic("некорректные операнды для операции умножения")
	case "/":
		if len(tokens[0]) > 12 {
			panic("длина строки не должна превышать 10 символов")
		}
		n, err := strconv.Atoi(tokens[2])
		if err != nil || n < 1 || n > 10 {
			panic("второй операнд должен быть числом от 1 до 10")
		}
		if tokens[0][0] == '"' {
			return strings.Trim(tokens[0], "\"")[:len(strings.Trim(tokens[0], "\""))/n], nil
		}
		panic("некорректные операнды для операции деления")
	default:
		panic("неподдерживаемая операция")
	}
}
