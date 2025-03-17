package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pattern 3")
	var n int64
	fmt.Print("Enter the value of n: ")
	fmt.Scanf("%d", &n)
	printPattern3(n)
}
