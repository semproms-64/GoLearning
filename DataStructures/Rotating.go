package main

import "fmt"

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Print("Insert n = ")
	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println("Error")
	}

	showArray(data)
	RotateArray(data, i)
	showArray(data)

}

func showArray(data []int) {
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i], " ")
	}

	fmt.Println()
}

func RotateArray(data []int, k int) {
	n := len(data)
	ReverseArray(data, 0, k-1)
	ReverseArray(data, k, n-1)
	ReverseArray(data, 0, n-1)
}

func ReverseArray(data []int, start int, end int) {
	i := start
	j := end

	for i < j {
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}
