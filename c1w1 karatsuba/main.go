package main

import (
	"fmt"
)


func main() {
	s1 := "3141592653589793238462643383279502884197169399375105820974944592"
	s2 := "2718281828459045235360287471352662497757247093699959574966967627"

	a := fromString(s1)
	b := fromString(s2)
	res := mult(a, b)

	fmt.Println(res)
}
