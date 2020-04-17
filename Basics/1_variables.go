package main

import (
	"fmt"
	"math"
)

func main() {

	/*-----------------------------------------------------------------------------------------*/
	/*Numeric types: Integers*/
	var myInt8 int8 = 97
	var myInt = 120
	var myUint uint = 500
	var myHexNumber = 0xFF
	var myOctalNumber = 034

	fmt.Printf("%d, %d, %d, %#x, %#o\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber)

	/*Integer type alias*/
	var myByte byte = 'A'
	var myRune rune = '♬'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)

	var numb byte = 120
	fmt.Printf("%c\n", numb)

	var numb_1 rune = 500
	fmt.Printf("%c\n", numb_1)

	var caracter byte = 'a'
	fmt.Printf("Result: %c\n", caracter-32)
	var caracter2 byte = 'z'
	fmt.Printf("Result2: %c\n", caracter2-32)

	/*-----------------------------------------------------------------------------------------*/
	/*Numeric types: Floating Point Types*/
	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2

	a++
	b += 10
	var res2 = a ^ b

	var r = 3.5
	var res3 = math.Pi * r * r

	fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)

	/*-----------------------------------------------------------------------------------------*/
	/*Boolean*/
	var truth = 3 <= 5
	var falsehood = 10 != 10

	var res4 = 10 > 20 && 5 == 5
	var res5 = 2*2 == 4 || 10%3 == 0

	fmt.Println(truth, falsehood, res4, res5)

	/*-----------------------------------------------------------------------------------------*/
	/*Complex Number*/

	/*-----------------------------------------------------------------------------------------*/
	/*Strings*/
}
