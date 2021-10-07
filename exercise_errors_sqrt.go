package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x) // automatically Error() function same as Stringer
	}
	z := 1.0
	for i := 0; i < 20; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Println(z);
	}
	return z, nil
}

func ExerciseErrors() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(-99))
}
