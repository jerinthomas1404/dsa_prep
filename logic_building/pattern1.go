package main

import (
	"fmt"
)

func printPattern1(n int64) {
	/*
		1. outer loop n times
		2. n times we're printing the value
		3. it should be printed in inner loop
		4. symmetry is not present
	*/

	for _ = range n {
		for _ = range n {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
