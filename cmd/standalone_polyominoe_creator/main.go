// This is a program to create N ordered polyominoes.
//   A polyomino is a shape made of simple squares.
//   An N ordered Polyomino is a polyomino with N number of squares that make up it's shape.
//   For a square to be apart of a polyominoe, it must be touching on it's edges and not on it's corner alone.

package main

import (
	"fmt"

	"github.com/cpuchip/polyominoes/pkg"
)

func main() {
	// Create a polyomino with 2 squares.
	p := pkg.Polyomino{
		Shape: [][]bool{
			{true},
			{true},
		},
	}
	// Print out the polyomino.
	fmt.Println(p.String())
}
