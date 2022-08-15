package main

import "fmt"

func divfiveandseven(numbers []int) {
	var divfiveandseven []int
	for _, num := range numbers {
		if num%5 == 0 && num%7 == 0 {
			divfiveandseven = append(divfiveandseven, num)
		}
	}
	fmt.Print("Числа, делящиеся на 5 и 7: ")
	fmt.Println(divfiveandseven)
}

func threedigits(numbers []int) {
	var threeDiffrentDigits []int
	for _, num := range numbers {
		if num > 99 && num < 1000 {
			a := num / 100
			b := (num - a*100) / 10
			c := num - a*100 - b*10
			if a != b && b != c && a != c {
				threeDiffrentDigits = append(threeDiffrentDigits, num)
			}
		}
	}
	fmt.Print("Трехзначные числа в записи которых нет одинаковых цифр: ")
	fmt.Println(threeDiffrentDigits)
}
func simpleNumber(numbers []int) {
	var simpleNumbersSlice []int
	for _, num := range numbers {
		if num < 10 {
			if num == 1 {
				simpleNumbersSlice = append(simpleNumbersSlice, num)
			}
			c := 0
			for i := num - 1; i > 1; i-- {
				if num%i != 0 {
					c++
				}
			}
			if c == num-2 {
				simpleNumbersSlice = append(simpleNumbersSlice, num)
			}
		} else {
			c := 0
			for i := 2; i <= 9; i++ {
				if num%i != 0 {
					c++
				}
			}
			if c == 8 {
				simpleNumbersSlice = append(simpleNumbersSlice, num)
			}
		}
	}
	fmt.Print("Простые числа: ")
	fmt.Println(simpleNumbersSlice)
}
