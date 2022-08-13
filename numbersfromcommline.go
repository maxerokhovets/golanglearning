package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var summ, counter int
	var prod = 1
	for _, arg := range os.Args {
		argInt, err := strconv.Atoi(arg)
		if err != nil {
			counter++
		} else {
			summ += argInt
			prod *= argInt
		}
	}
	fmt.Print("Сумма аргументов: ")
	fmt.Println(summ)
	fmt.Print("Произведение аргументов: ")
	fmt.Println(prod)
	fmt.Print("Аргументов не удалось конвертировать в цифры: ")
	fmt.Println(counter)
}
