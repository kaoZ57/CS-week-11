package main

import "fmt"

func main() {
	say(loop(0))
}

func say(N int) {
	fmt.Printf("\rnumber = %v", N)
}

func loop(i int) int {
	for i = 0; i <= 1000000; i++ {
	}
	return i
}
