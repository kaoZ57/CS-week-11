package main

import "fmt"

func main() {
	AZ := []int{}
	kk(0, &AZ)
	fmt.Print("Number 0-10 = ")
	fmt.Scan(&num)
	fmt.Println("Number =", AZ[num])
}

func kk(i int, AZ *[]int) {
	for i = 0; i < 10; i++ {
		*AZ = append(*AZ, i)
	}
}
