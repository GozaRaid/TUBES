package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	x := math.Abs(x1 - x2)
	y := math.Abs(y1 - y2)
	fmt.Print(x + y)
}
