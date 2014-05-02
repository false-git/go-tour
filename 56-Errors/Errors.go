package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64
func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    }
    z := 1.0
    z2 := 1.0
    for i := 0; i < 100000000; i++ {
        z2 = z - (z * z - f) / (2 * z)
        if math.Abs(z2 - z) < 0.000000000000001 {
            break
        }
        z = z2;
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
