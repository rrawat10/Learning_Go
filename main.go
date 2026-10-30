package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)
	// TODO: build a Point for each pair of coordinates, then print the
	// squared distance between them instead of this placeholder.

	point1 := Point{X: x1, Y: y1}
	point2 := Point{X: x2, Y: y2}

	dx := point1.X - point2.X
	dy := point1.Y - point2.Y
	dist := dx*dx + dy*dy

	fmt.Println(dist)
}
