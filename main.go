package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	name, _ := r.ReadString('\n')
	name = strings.TrimRight(name, "\r\n")
	qtyLine, _ := r.ReadString('\n')
	qty, _ := strconv.Atoi(strings.TrimSpace(qtyLine))
	priceLine, _ := r.ReadString('\n')
	price, _ := strconv.ParseFloat(strings.TrimSpace(priceLine), 64)
	total := float64(qty) * price

	// TODO: replace this Println with one fmt.Printf that lays the
	// four values out in the receipt columns described in the exercise.
	// one line: the name left-aligned in 12 columns, a space, the quantity right-aligned in 3 columns, x, the price right-aligned in 6 columns with 2 decimals, =, and the total (quantity times price) right-aligned in 8 columns with 2 decimals. Widths are minimums: longer values are printed in full.
	fmt.Printf("%-12s %3d x %6.2f = %8.2f\n", name, qty, price, total)

}
