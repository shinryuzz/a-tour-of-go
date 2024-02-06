package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %1.0f", e)
}

func Sqrt(x float64) (float64, error) {
	z := 1.0

	if x < 0 {
		return z,  ErrNegativeSqrt(x)
	}
	
	for i:=0;i<10;i++ {
		prev := z
		z -= (z*z - x) / (2*z)
		if math.Abs(prev - z) < 0.0001 {
			return z, nil
		}

	}
	return z, nil
}

func main() {
	if i, err := Sqrt(2); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%v\n", i)
	}
	if i, err := Sqrt(-2); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%v\n", i)
	}
	
}
