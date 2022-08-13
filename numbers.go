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
