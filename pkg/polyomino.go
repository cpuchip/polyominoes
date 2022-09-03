package pkg

import "fmt"

// A struct to hold a polyomino.
//
//	A polyomino is a shape made of simple squares. A square is represented by a 2D bool array
//	where true represents a square and false represents a space.
type Polyomino struct {
	Order      int // How many squares make up the polyomino.
	M          int // The width of the polyomino.
	N          int // The length of the polyomino.
	Shape      [][]bool
	Transforms []Polyomino
}

// write a function that takes in a polyomino and returns a list of all the polyominoes that can be made from it by one order.
// To create a new polyomino, you take a square and add it systematically to the polyomino.
// Add a new row to the top of the original polyomino and place a square in the top left corner
// so that it is touching the original polyomino on it's top edge in an open space. Save the new polyomino.
// repeat the process for the next square in the top row. continue this process by adding columns on the left and right, and bottom
// only placing one square in each new polyomino, until all permutations of new polyominoes are found.
//
// The function should return a list of all the polyominoes that can be made from the original polyomino.
// The returned polyominoes should be trimmed to remove rows or columns that are all false.
//
// Example:
//
//		Input:
//			Polyomino{
//				Order: 2,
//				Shape: [][]bool{
//					{true},
//					{true},
//				},
//			}
//		Output:
//			[]Polyomino{
//				Polyomino{
//					Order: 3,
//					Shape: [][]bool{
//						{true},
//						{true},
//						{true},
//					},
//				},
//				Polyomino{
//					Order: 3,
//					Shape: [][]bool{
//						{true, true},
//						{false, true},
//					},
//				},
//				Polyomino{
//					Order: 3,
//					Shape: [][]bool{
//						{true, true},
//						{true, false},
//					},
//				},
//				Polyomino{
//					Order: 3,
//					Shape: [][]bool{
//						{false, true},
//						{true, true},
//					},
//				},
//				Polyomino{
//					Order: 3,
//					Shape: [][]bool{
//						{true, false},
//						{true, true},
//					},
//				},
//				Polyomino{
//					Order: 3,
//					Shape: [][]bool{
//						{true},
//						{true},
//						{true},
//					},
//				},
//			}
//

// see if two polyominoes are the same
func isSamePolyomino(p1, p2 Polyomino) bool {
	if p1.Order != p2.Order {
		return false
	}
	for i, row := range p1.Shape {
		for j, square := range row {
			if square != p2.Shape[i][j] {
				return false
			}
		}
	}
	return true
}

// calculate the M and N of a polyomino
func (p *Polyomino) CalculateMN() {
	p.M = len(p.Shape)
	p.N = len(p.Shape[0])
}

// Remove duplicate Polyominoes from a list of Polyominoes.
// The returned polyominoes should be unqueue to remove any duplicate polyominoes.
//
//	A duplicate polyomino is a polyomino that is the same as another polyomino when rotated or flipped.
func RemoveDuplicates(polyominoes []Polyomino) []Polyomino {
	var uniquePolyominoes []Polyomino
	for _, polyomino := range polyominoes {
		found := false
		for _, uniquePolyomino := range uniquePolyominoes {
			if polyomino.Order == uniquePolyomino.Order && isSamePolyomino(polyomino, uniquePolyomino) {
				found = true
				break
			}
		}
		if !found {
			uniquePolyominoes = append(uniquePolyominoes, polyomino)
		}
	}
	return uniquePolyominoes
}

// Print out a representation of a polyomino.
func (p Polyomino) String() string {
	var s string
	fmt.Printf("An %d ordered polyomino of Width %d and Length %d:\n", p.Order, len(p.Shape[0]), len(p.Shape))
	for _, row := range p.Shape {
		for _, square := range row {
			if square {
				s += "X"
			} else {
				s += " "
			}
		}
		s += "\n"
	}
	return s
}

// take the current polyomino and horizontally flip it.
func (p Polyomino) HorizontalFlip() Polyomino {
	var flippedPolyomino Polyomino
	flippedPolyomino.Order = p.Order
	flippedPolyomino.M = p.M
	flippedPolyomino.N = p.N
	for _, row := range p.Shape {
		var flippedRow = make([]bool, len(row))
		n := len(row) - 1
		for i := n; i >= 0; i-- {
			flippedRow[i] = row[n-i]
		}
		flippedPolyomino.Shape = append(flippedPolyomino.Shape, flippedRow)
	}
	return flippedPolyomino
}

// take the current polyomino and vertically flip it.
func (p Polyomino) VerticalFlip() Polyomino {
	var flippedPolyomino Polyomino
	flippedPolyomino.Order = p.Order
	flippedPolyomino.M = p.M
	flippedPolyomino.N = p.N
	for i := len(p.Shape) - 1; i >= 0; i-- {
		flippedPolyomino.Shape = append(flippedPolyomino.Shape, p.Shape[i])
	}
	return flippedPolyomino
}

// take the current polyomino and take the D+ rotation of it.
func (p Polyomino) DPlusReflection() Polyomino {
	var reflectedPolyomino Polyomino
	reflectedPolyomino.Order = p.Order
	reflectedPolyomino.M = p.N
	reflectedPolyomino.N = p.M
	m := len(p.Shape) - 1
	for i := m; i >= 0; i-- {
		var reflectedRow = make([]bool, len(p.Shape[i]))
		n := len(p.Shape[i]) - 1
		for j := n; j >= 0; j-- {
			reflectedRow[j] = p.Shape[i][n-j]
		}
		reflectedPolyomino.Shape = append(reflectedPolyomino.Shape, reflectedRow)
	}
	return reflectedPolyomino
}

// take the current polyomino and generate all of its transformations.
func (p Polyomino) Transformations() []Polyomino {
	p.Transforms = nil
	p.Transforms = make([]Polyomino, 0, 7)
	// Sraight ( no transformation )
	// transformations = append(transformations, p)

	// Horizontal Flip
	p.Transforms = append(p.Transforms, p.HorizontalFlip())

	// 180 degrees
	p.Transforms = append(p.Transforms, p.Transforms[0].VerticalFlip())

	// Vertical Flip
	p.Transforms = append(p.Transforms, p.Transforms[1].HorizontalFlip())

	// If the polyomino is a square, then we can do the D+ reflection
	if p.M == p.N {
		// 270 degrees
		p.Transforms = append(p.Transforms, p.Transforms[2].DPlusReflection())

		// D- symmetry
		p.Transforms = append(p.Transforms, p.Transforms[3].HorizontalFlip())

		// 90 degrees
		p.Transforms = append(p.Transforms, p.Transforms[4].VerticalFlip())

		// D+ symmetry
		p.Transforms = append(p.Transforms, p.Transforms[5].HorizontalFlip())
	}

	return p.Transforms
}

// Expand the polyomino by one row and column in each direction.
func (p Polyomino) ExpandOrder() Polyomino {
	var expandedPolyomino Polyomino
	expandedPolyomino.Order = p.Order + 1
	expandedPolyomino.M = p.M + 2
	expandedPolyomino.N = p.N + 2
	expandedPolyomino.Shape = make([][]bool, expandedPolyomino.M)
	for i := range expandedPolyomino.Shape {
		expandedPolyomino.Shape[i] = make([]bool, expandedPolyomino.N)
	}
	for i, row := range p.Shape {
		for j, square := range row {
			expandedPolyomino.Shape[i+1][j+1] = square
		}
	}
	return expandedPolyomino
}

// trim the polyomino so it has no rows or columns of all false squares on its sides.
func (p Polyomino) Trim() Polyomino {
	var trimmedPolyomino Polyomino
	trimmedPolyomino.Order = p.Order
	trimmedPolyomino.M = p.M
	trimmedPolyomino.N = p.N

	// Check to trim the top and bottom
	var trimTop bool = true
	var trimBottom bool = true
	for i := 0; i < len(p.Shape[0]); i++ {
		if p.Shape[0][i] {
			trimTop = false
			break
		}
	}
	for i := 0; i < len(p.Shape[len(p.Shape)-1]); i++ {
		if p.Shape[len(p.Shape)-1][i] {
			trimBottom = false
			break
		}
	}

	// Check to trim the left and right
	var trimLeft bool = true
	var trimRight bool = true
	for i := 0; i < len(p.Shape); i++ {
		if p.Shape[i][0] {
			trimLeft = false
			break
		}
	}
	for i := 0; i < len(p.Shape); i++ {
		if p.Shape[i][len(p.Shape[i])-1] {
			trimRight = false
			break
		}
	}

	// Trim the top and bottom
	if trimTop {
		trimmedPolyomino.M--
	} else {
		trimmedPolyomino.Shape = append(trimmedPolyomino.Shape, p.Shape[0])
	}
	trimmedPolyomino.Shape = append(trimmedPolyomino.Shape, p.Shape[1:len(p.Shape)-1]...)
	if trimBottom {
		trimmedPolyomino.M--
	} else {
		trimmedPolyomino.Shape = append(trimmedPolyomino.Shape, p.Shape[len(p.Shape)-1])
	}

	// Trim the left and right
	if trimLeft {
		trimmedPolyomino.N--
		// Loop through the rows and trim the left side one element
		for i := range trimmedPolyomino.Shape {
			trimmedPolyomino.Shape[i] = trimmedPolyomino.Shape[i][1:]
		}
	}
	if trimRight {
		trimmedPolyomino.N--
		// Loop through the rows and trim the right most element
		for i := range trimmedPolyomino.Shape {
			trimmedPolyomino.Shape[i] = trimmedPolyomino.Shape[i][:len(trimmedPolyomino.Shape[i])-1]
		}
	}

	return trimmedPolyomino
}
