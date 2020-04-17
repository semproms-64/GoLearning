package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a[0])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var b [4]rune
	b[0] = 'a'
	b[1] = 'b'
	b[2] = 'c'
	b[3] = 'd'
}
