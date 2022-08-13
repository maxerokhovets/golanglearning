package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Введите n:")
	fmt.Scan(&n)
	var numbers []int
	for i := 0; i < n; i++ {
		fmt.Println("Введите чило: ")
		var num int
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}
	evenoddnumb(numbers)
	maxminnumber(numbers)
}

func evenoddnumb(numbers []int) {
	var even, odd []int
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	fmt.Print("Четные числа: ")
	fmt.Println(even)
	fmt.Print("Нечетные числа: ")
	fmt.Println(odd)

}

func maxminnumber(numbers []int) {
	max := numbers[0]
	min := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Print("Максимальное число: ")
	fmt.Println(max)
	fmt.Print("Минимальное число: ")
	fmt.Println(min)

}
