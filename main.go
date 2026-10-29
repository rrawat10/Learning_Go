package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	// TODO: keep only the whole part of dividing a by b.
	quotient := 0
    quotient = a/b
	// TODO: what is left over after that division.
	remainder := 0
    remainder = a%b

	fmt.Println(quotient)
	fmt.Println(remainder)
}
