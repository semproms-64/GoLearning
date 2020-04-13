package main

import "fmt"

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8}

	suma := SumData(data)

	fmt.Println(suma)
}

func SumData(data []int) int {
	size := len(data)
	total := 0

	for index := 0; index < size; index++ {
		total = total + data[index]
	}

	return total
}
