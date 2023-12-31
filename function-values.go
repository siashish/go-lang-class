package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// func hypot(x,y float64)float64{
// 	return math.Sqrt(x*x + y*y)
// }

// hypot := func(x, y float64)float64{
// 	return math.Sqrt(x*x + y*y)
// }

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))    // 3*3 + 4*4
	fmt.Println(compute(math.Pow)) // 3*3*3*3
}
