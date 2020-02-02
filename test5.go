package main

import "fmt"

func main() {
  switch 1 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	}
  switch 1 {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	}
}
