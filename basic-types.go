package main

import (
	"fmt"
	"math/cmplx"
)

// bool
// string
// int int8 int16 int32 int64  // int8 = -128 to 127
// uint uint8 uint16 uint32 uint64 uintptr // uint8 = 0 to 256
// byte // alias for uint8
// rune // alias for int32
// represents a Unicode code point
// float32 float64
// complex64 complex128

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i) // (x+y)pow2 = xpow2 + 2xy + ypow2
)

func main() {
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)
}
