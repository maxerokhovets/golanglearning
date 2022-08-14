package main

import "fmt"

func threenine(numbers []int) {
	var threeninediv []int
	for _, num := range numbers {
		if num%3 == 0 || num%9 == 0 {
			threeninediv = append(threeninediv, num)
		}
	}
	fmt.Print("Числа делящиеся на 3 или на 9: ")
	fmt.Println(threeninediv)
}
