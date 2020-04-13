/*
 * Author: Luis Adelantado Romero
 * Date: lun abr 13 20:02:40 CEST 2020
 * Static strongly typed language
 * It compiles to machine code
*/
package main

import "fmt"

func main() {
	var v1 int
	var v2 int //Initialized to 0 by default

	v1 = 100

	fmt.Println("Value stored in v1 ::",v1);
	fmt.Println("Value stored in v2 ::",v2);

	const pi = 3.14;
	e := 2.71; //Type determined by compiler.

	fmt.Println("Pi = ", pi);
	fmt.Println("E = ", e);

	var text string

	text = "I am just playing..."
	fmt.Println(text);

	var bool boolean = true
	var byte byte;

	fmt.Println("Basic data types in Go: ");
	fmt.Println("Boolean: ");
	fmt.Println("Byte: ");
	fmt.Println("Integer: ");
	fmt.Println("Unsigned integer: ");
	fmt.Println("Float: ");
	fmt.Println("String: ")
	fmt.Println("Rune: ")
	fmt.Println("Complex: ")


}