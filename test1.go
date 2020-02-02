package main

func main() {
	AZ := []int{}
}

func jj(i int, AZ *[]int) {
	for i = 0; i < 10; i++ {
		*AZ = append(*AZ, i)
	}
}
