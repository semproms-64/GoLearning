package main

import "fmt"

func esPar(numb rune) bool {
	return numb%2 == 0
}

func main() {

	var b [4]rune
	b[0] = 'a'
	b[1] = 'b'
	b[2] = 'c'
	b[3] = 'd'

	fmt.Printf("%c\n", b)
	for i, number := range b {
		if !esPar(number) {
			number = number - 32
		}
		fmt.Printf("%d: %c\n", i, number)
	}

	var n_tr int
	n_tr = 25

	for n_tr > 1 {
		fmt.Printf("%d\n", n_tr)
		n_tr--
	}

}
