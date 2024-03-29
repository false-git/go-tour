package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    x1 := 0
    x2 := 1
    return func() int {
        r := x1
        x1, x2 = x2, x1 + x2
        return r
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
