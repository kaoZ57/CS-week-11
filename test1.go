package main

import "fmt"

func main() {
	AZ := []int{}
	jj(0, &AZ)
	fmt.Println(AZ)
}

func jj(i int, AZ *[]int) {
	for i = 0; i < 10; i++ {
		*AZ = append(*AZ, i)
	}
}
