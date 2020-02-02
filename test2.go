package main

import (
	"fmt"
	"strings"
)

func main() {
	var vr1 string
	var num1 int
	var num2 int
	fmt.Println("Plus or Minus =")
	fmt.Scan(&vr1)
X:
	fmt.Println("number =")
	fmt.Println(&num1, &num2)
	if strings.Contains(vr1, "Plus") == true {
		plus(num1, num2)
	} else {
		goto X
	}
	if strings.Contains(vr1, "Minus") == true {
		minus(num1, num2)
	} else {
		goto X
	}
}

func plus(n1, n2 int) {
	X := n1 + n2
	fmt.Println(X)
}
