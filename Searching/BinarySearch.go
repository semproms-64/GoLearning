package main

import (
	"fmt"
)

var pos int

func main() {
	data := []int{1, 2, 8, 25, 50, 125, 512, 1024}

	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		fmt.Println("Error")
	}

	find := BinarySearch(data, i)

	if find {
		fmt.Println("Element found in ", pos, " position.")
	} else {
		fmt.Println("Element not found!")
	}

}

func BinarySearch(data []int, value int) bool {
	size := len(data)
	var mid int

	low := 0
	high := size - 1

	for low <= high {
		mid = low + (high-low)/2

		if data[mid] == value {
			pos = mid
			return true
		} else {
			if data[mid] < value {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return false
}
