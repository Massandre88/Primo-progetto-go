package main

import (
	"golang.org/x/tour/pic"
	//"math"
)
func Pic(dx, dy int) [][]uint8 {
	v := make([][]uint8,dy)
	for i := range v {
		v[i] = make([]uint8,dx)
		for j := range v[i] {
			v[i][j]=uint8((i+j)/2)
		}
	}
	return v
}

func main() {
	pic.Show(Pic)
}
