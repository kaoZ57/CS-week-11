package main

import "fmt"

func main() {
	say(loop(0))
}

func loop(i int) int {
	for i = 0; i <= 1000000; i++ {
	}
	return i
}
