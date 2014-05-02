package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    z := 1.0
    z2 := 1.0
    for i := 0; i < 100000000; i++ {
        z2 = z - (z * z - x) / (2 * z)
        if math.Abs(z2 - z) < 0.000000000000001 {
            break
        }
        z = z2;
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2), math.Sqrt(2) == Sqrt(2))
}
