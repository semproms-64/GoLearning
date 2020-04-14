/*
 * Author: Luis Adelantado Romero
 * Date: lun abr 13 20:02:40 CEST 2020
 * Static strongly typed language
 * It compiles to machine code
 */
package main

import "fmt"

func swapRune(r rune) rune {
	switch {
	case 97 <= rune && r <= 122:
		return r - 32
	case 65 <= r && r <= 90:
		return r + 32
	default:
		return r
	}
}

func main() {
	var v1 int
	var v2 int //Initialized to 0 by default

	v1 = 100

	fmt.Println("Value stored in v1 ::", v1)
	fmt.Println("Value stored in v2 ::", v2)

	const pi = 3.14
	e := 2.71 //Type determined by compiler.

	fmt.Println("Pi = ", pi)
	fmt.Println("E = ", e)

	var text string

	text = "I am just playing..."
	fmt.Println(text)

	var X uint8 = 225
	var Y int16 = 32767
	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)
	str := "Hello world"

	fmt.Println("Swap rune: ", swapRune('a'))

	fmt.Println("Basic data types in Go: ")

	fmt.Println("Unsigned integer: ", X)
	fmt.Println("Integer: ", Y)
	fmt.Println("Complex number 128: ", a)
	fmt.Println("Complex number 64: ", b)
	fmt.Println("String: ", str)

}
