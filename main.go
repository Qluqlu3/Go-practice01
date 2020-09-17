package main

import "fmt"

var text string = "PPP"
var number = 123

const content = "ASDF"

func main() {

	sub := 9090
	var inText = "111"
	const inNumber int = 1230000

	const (
		a1 = iota
		a2
		a3
		a4
		a5 = iota
		a6
		a7
	)

	type MyInt int

	var myNum MyInt = 2

	fmt.Println("AAAA")
	fmt.Println(text)
	fmt.Println(number)
	fmt.Println(content)
	fmt.Println(sub)
	fmt.Println(inText)
	fmt.Println(inNumber)
	fmt.Println(a1, a2, a3, a4, a5, a6, a7)
	// fmt.Println(myNum + inNumber)NG
}
