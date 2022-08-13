package main

import "fmt"

func main() {
	var pass string
	str := "string"
	fmt.Println("Введите пароль:")
	fmt.Scan(&pass)
	if pass == str {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль не верный")
	}
}
