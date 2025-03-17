package main

import "fmt"

func printPattern3(n int64) {
	/*
		1. outer loop n times
		2. n times we're printing the value
		3. it should be printed in inner loop
		4. symmetry is not present
	*/
	var outer_loop, inner_loop int64
	for outer_loop = 0; outer_loop < n; outer_loop++ {
		for inner_loop = 0; inner_loop <= outer_loop; inner_loop++ {
			fmt.Print(inner_loop + 1)
		}
		fmt.Println()
	}
}
