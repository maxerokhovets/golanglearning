package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(i)
		fmt.Print(" ")
		fmt.Print(os.Args[i])
		//fmt.Println()
	}
}
