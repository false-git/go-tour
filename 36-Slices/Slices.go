package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    r := make([][]uint8, dy)
    for i := range r {
        r[i] = make([]uint8, dx)
        for j := range r[i] {
            r[i][j] = uint8((i + j) / 2)
        }
    }
    return r
}

func main() {
    pic.Show(Pic)
}
