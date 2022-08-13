package main

import "fmt"

func main() {
	fmt.Println("Введите имя:")
	var name string
	fmt.Scan(&name)
	fmt.Print("Hello, ")
	fmt.Print(name)
}
