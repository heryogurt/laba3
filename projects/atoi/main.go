package main

import (
	"fmt"
)

func main() {
	var ch string
	fmt.Println("enter number: ")
	fmt.Scan(&ch)
	for i := 0; i < len(ch); i++ {
		a := ch[i] - 48
		fmt.Print(a * a)
	}
}
