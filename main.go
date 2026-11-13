package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
	} else if len(os.Args) < 4 {
		fmt.Printf("Недостаточно аргументов.\n")
		return
	} else {
		fmt.Printf("Слишком много аргументов.\n")
		return
	}

	num1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Вводимое значение должно быть численным!\n")
		return
	}

	num2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Вводимое значение должно быть численным!\n")
		return
	}

	res := os.Args[2]

	switch res {
	case "+":
		res := num1 + num2
		fmt.Printf("Результат: %d\n", res)
	case "-":
		res := num1 - num2
		fmt.Printf("Результат: %d\n", res)
	case "/":
		res := num1 / num2
		fmt.Printf("Результат: %d\n", res)
	case "*":
		res := num1 * num2
		fmt.Printf("Результат: %d\n", res)
	default:
		fmt.Printf("Введите что-то из перечисленного: +, -, /, *\n")
		return
	}

}
