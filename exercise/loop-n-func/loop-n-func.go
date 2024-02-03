package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	
	for i:=0;i<10;i++ {
		prev := z
		z -= (z*z - x) / (2*z)
		
		fmt.Println(i+1, z)
		
		if math.Abs(prev - z) < 0.0001 {
			fmt.Println("no diff., break loop.")
			return z
		}

	}

	return z
}

func main() {
	fmt.Println(Sqrt(169))
}
