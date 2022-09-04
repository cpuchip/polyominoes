package pkg

import "testing"

// test horizontal flip
func TestPolyomino_HorizontalFlip(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	pUT := p.HorizontalFlip()
	if pUT.Shape[0][0] != true || pUT.Shape[0][1] != true || pUT.Shape[1][0] != true || pUT.Shape[1][1] != false {
		t.Error("HorizontalFlip failed")
	}
}

func TestPolyomino_HorizontalFlipLarger(t *testing.T) {
	p := Polyomino{
		Order: 4,
		M:     3,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
			{false, true},
		},
	}
	pUT := p.HorizontalFlip()
	if pUT.Order != p.Order {
		t.Error("HorizontalFlip failed")
	}
	if pUT.Shape[0][0] != true || pUT.Shape[0][1] != true ||
		pUT.Shape[1][0] != true || pUT.Shape[1][1] != false ||
		pUT.Shape[2][0] != true || pUT.Shape[2][1] != false {
		t.Error("HorizontalFlip failed")
	}
}

// test check if two polyominoes are the same
func TestPolyomino_IsSamePolyomino(t *testing.T) {
	p1 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	p2 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	if !isSamePolyomino(p1.Shape, p2.Shape) {
		t.Error("isSamePolyomino failed")
	}
}

// test if two polyominos are not the same
func TestPolyomino_NotSamePolyomino(t *testing.T) {
	p1 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	p2 := Polyomino{
		Order: 4,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{true, true},
		},
	}
	if isSamePolyomino(p1.Shape, p2.Shape) {
		t.Error("isSamePolyomino failed")
	}
}

func TestPolyomino_IsSamePolyominoLarger(t *testing.T) {
	p1 := Polyomino{
		Order: 3,
		M:     3,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
			{false, true},
		},
	}
	p2 := Polyomino{
		Order: 3,
		M:     3,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
			{false, true},
		},
	}
	if !isSamePolyomino(p1.Shape, p2.Shape) {
		t.Error("isSamePolyomino failed")
	}
}

// test check if two polyominoes are the same
func TestPolyomino_IsSame(t *testing.T) {
	p1 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	p1.Transformations()
	p2 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	p2.Transformations()
	if !p1.IsSame(&p2) {
		t.Error("isSamePolyomino failed")
	}
}

// test check if two polyominoes are the same
func TestPolyomino_IsSame_alt01(t *testing.T) {
	p1 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, false},
			{true, true},
		},
	}
	p1.Transformations()
	p2 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	p2.Transformations()
	if !p1.IsSame(&p2) {
		t.Error("isSamePolyomino failed")
	}
}

// test check if two polyominoes are the same
func TestPolyomino_IsSame_alt02(t *testing.T) {
	p1 := Polyomino{
		Order: 3,
		M:     1,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
		},
	}
	p1.Transformations()
	p2 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	p2.Transformations()
	if p1.IsSame(&p2) {
		t.Error("isSamePolyomino failed")
	}
}

// test vertical flip
func TestPolyomino_VerticalFlip(t *testing.T) {
	p := Polyomino{
		Order: 2,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	pUT := p.VerticalFlip()
	if pUT.Shape[0][0] != false || pUT.Shape[0][1] != true ||
		pUT.Shape[1][0] != true || pUT.Shape[1][1] != true {
		t.Error("VerticalFlip failed")
	}
}

func TestPolyomino_VerticalFlipLarger(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     2,
		N:     3,
		Shape: [][]bool{
			{true, true},
			{false, true},
			{false, true},
		},
	}
	pUT := p.VerticalFlip()
	if pUT.Order != p.Order {
		t.Error("VerticalFlip failed")
	}
	if pUT.Shape[0][0] != false || pUT.Shape[0][1] != true ||
		pUT.Shape[1][0] != false || pUT.Shape[1][1] != true ||
		pUT.Shape[2][0] != true || pUT.Shape[2][1] != true {
		t.Error("VerticalFlip failed")
	}
}

// test D+ reflection
func TestPolyomino_DPlusReflection(t *testing.T) {
	p := Polyomino{
		Order: 2,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	pUT := p.DPlusReflection()
	if pUT.Shape[0][0] != true || pUT.Shape[0][1] != false ||
		pUT.Shape[1][0] != true || pUT.Shape[1][1] != true {
		t.Error("DPlusReflection failed")
	}
}

// test Transformations large
func TestPolyomino_TransformationsLarge(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     2,
		N:     3,
		Shape: [][]bool{
			{true, true},
			{false, true},
			{false, true},
		},
	}
	pUT := p.Transformations()
	if len(pUT) != 3 {
		t.Error("Transformations failed")
	}

	// HorizontalFlip
	h := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true},
			{true, false},
			{true, false},
		},
	}
	if !isSamePolyomino(pUT[0].Shape, h.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// 180 degree
	r180 := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, false},
			{true, false},
			{true, true},
		},
	}
	if !isSamePolyomino(pUT[1].Shape, r180.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// VerticalFlip
	v := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, true},
			{false, true},
			{true, true},
		},
	}
	if !isSamePolyomino(pUT[2].Shape, v.Shape) {
		t.Error("HorizontalFlip failed")
	}
}

// test Transformations large
func TestPolyomino_TransformationsSquare(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
			{true, true, false},
			{true, false, false},
		},
	}
	pUT := p.Transformations()
	if len(pUT) != 7 {
		t.Error("Transformations failed")
	}

	// HorizontalFlip
	h := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
			{false, true, true},
			{false, false, true},
		},
	}
	if !isSamePolyomino(pUT[0].Shape, h.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// 180 degree
	r180 := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, true},
			{false, true, true},
			{true, true, true},
		},
	}
	if !isSamePolyomino(pUT[1].Shape, r180.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// VerticalFlip
	v := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, false, false},
			{true, true, false},
			{true, true, true},
		},
	}
	if !isSamePolyomino(pUT[2].Shape, v.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// 270 degree
	r270 := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
			{false, true, true},
			{false, false, true},
		},
	}
	if !isSamePolyomino(pUT[3].Shape, r270.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// D- reflection
	dMinus := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
			{true, true, false},
			{true, false, false},
		},
	}
	if !isSamePolyomino(pUT[4].Shape, dMinus.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// 90 degree
	r90 := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, false, false},
			{true, true, false},
			{true, true, true},
		},
	}
	if !isSamePolyomino(pUT[5].Shape, r90.Shape) {
		t.Error("HorizontalFlip failed")
	}
	// D+ reflection
	dPlus := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, true},
			{false, true, true},
			{true, true, true},
		},
	}
	if !isSamePolyomino(pUT[6].Shape, dPlus.Shape) {
		t.Error("HorizontalFlip failed")
	}
}

// Test the ExpandOrder function
func TestPolyomino_ExpandOrder(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
			{true, true, false},
			{true, false, false},
		},
	}
	pEO := p.ExpandOrder()
	if pEO.Order != 4 {
		t.Error("ExpandOrder failed")
	}
	if pEO.M != 5 {
		t.Error("ExpandOrder failed")
	}
	if pEO.N != 5 {
		t.Error("ExpandOrder failed")
	}
	p2 := Polyomino{
		Order: 4,
		M:     5,
		N:     5,
		Shape: [][]bool{
			{false, false, false, false, false},
			{false, true, true, true, false},
			{false, true, true, false, false},
			{false, true, false, false, false},
			{false, false, false, false, false},
		},
	}
	if !isSamePolyomino(pEO.Shape, p2.Shape) {
		t.Error("ExpandOrder failed")
	}
}

// test trim polyomino function
func TestPolyomino_Trim(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, false},
			{true, false, false},
			{false, false, false},
		},
	}
	pT := p.Trim()
	if pT.Order != 3 {
		t.Error("Trim failed")
	}
	if pT.M != 2 {
		t.Error("Trim failed")
	}
	if pT.N != 2 {
		t.Error("Trim failed")
	}
	p2 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{true, false},
		},
	}
	if !isSamePolyomino(pT.Shape, p2.Shape) {
		t.Error("Trim failed")
	}
}

// test trim polyomino function
func TestPolyomino_TrimAlt_01(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, false},
			{false, true, true},
			{false, false, true},
		},
	}
	pT := p.Trim()
	if pT.Order != 3 {
		t.Error("Trim failed")
	}
	if pT.M != 2 {
		t.Error("Trim failed")
	}
	if pT.N != 2 {
		t.Error("Trim failed")
	}
	p2 := Polyomino{
		Order: 3,
		M:     2,
		N:     2,
		Shape: [][]bool{
			{true, true},
			{false, true},
		},
	}
	if !isSamePolyomino(pT.Shape, p2.Shape) {
		t.Error("Trim failed")
	}
}

// test trim polyomino function
func TestPolyomino_TrimAlt_02(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, false},
			{true, true, true},
			{false, false, false},
		},
	}
	pT := p.Trim()
	if pT.Order != 3 {
		t.Error("Trim failed")
	}
	if pT.M != 1 {
		t.Error("Trim failed")
	}
	if pT.N != 3 {
		t.Error("Trim failed")
	}
	p2 := Polyomino{
		Order: 3,
		M:     1,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
		},
	}
	if !isSamePolyomino(pT.Shape, p2.Shape) {
		t.Error("Trim failed")
	}
}

// test trim polyomino function
func TestPolyomino_TrimAlt_03(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, true, false},
			{false, true, false},
			{false, true, false},
		},
	}
	pT := p.Trim()
	if pT.Order != 3 {
		t.Error("Trim failed")
	}
	if pT.M != 3 {
		t.Error("Trim failed")
	}
	if pT.N != 1 {
		t.Error("Trim failed")
	}
	p2 := Polyomino{
		Order: 3,
		M:     3,
		N:     1,
		Shape: [][]bool{
			{true},
			{true},
			{true},
		},
	}
	if !isSamePolyomino(pT.Shape, p2.Shape) {
		t.Error("Trim failed")
	}
}

// test trim polyomino function
func TestPolyomino_TrimAlt_04(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, true, false},
		},
	}
	pT := p.Trim()
	if pT.Order != 3 {
		t.Error("Trim failed")
	}
	if pT.M != 3 {
		t.Error("Trim failed")
	}
	if pT.N != 3 {
		t.Error("Trim failed")
	}
	p2 := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, true, false},
			{true, true, true},
			{false, true, false},
		},
	}
	if !isSamePolyomino(pT.Shape, p2.Shape) {
		t.Error("Trim failed")
	}
}

// test trim polyomino function
func TestPolyomino_TrimAlt_05(t *testing.T) {
	p := Polyomino{
		Order: 1,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	}
	pT := p.Trim()
	if pT.Order != 1 {
		t.Error("Trim failed")
	}
	if pT.M != 1 {
		t.Error("Trim failed")
	}
	if pT.N != 1 {
		t.Error("Trim failed")
	}
	p2 := Polyomino{
		Order: 1,
		M:     1,
		N:     1,
		Shape: [][]bool{
			{true},
		},
	}
	if !isSamePolyomino(pT.Shape, p2.Shape) {
		t.Error("Trim failed")
	}
}

// test the CalculateMN function
func TestPolyomino_CalculateMN(t *testing.T) {
	p := Polyomino{
		Order: 5,
		Shape: [][]bool{
			{true, true, false},
			{true, true, false},
			{false, true, false},
		},
	}
	p.CalculateMN()
	if p.M != 3 {
		t.Error("CalculateMN failed")
	}
	if p.N != 3 {
		t.Error("CalculateMN failed")
	}
}

// test the ValidSquare function with a valid square
func TestPolyomino_ValidSquare_01(t *testing.T) {
	p := Polyomino{
		Order: 3,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, false},
			{true, true, false},
			{false, true, false},
		},
	}
	if !p.ValidSquare(0, 0) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(0, 1) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(0, 2) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(1, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(1, 1) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(1, 2) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(2, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 1) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(2, 2) {
		t.Error("ValidSquare failed")
	}
}

// test the ValidSquare function with a valid square
func TestPolyomino_ValidSquare_02(t *testing.T) {
	p := Polyomino{
		Order: 2,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, false},
			{false, true, false},
			{false, true, false},
		},
	}
	// must be false
	if p.ValidSquare(0, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(0, 2) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(1, 1) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 1) {
		t.Error("ValidSquare failed")
	}
	// must be true
	if !p.ValidSquare(0, 1) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(1, 0) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(1, 2) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(2, 0) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(2, 2) {
		t.Error("ValidSquare failed")
	}
}

// test the ValidSquare function with a valid square
func TestPolyomino_ValidSquare_03(t *testing.T) {
	p := Polyomino{
		Order: 2,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{false, false, false},
			{false, true, false},
			{false, false, false},
		},
	}
	// must be false
	if p.ValidSquare(0, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(0, 2) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 2) {
		t.Error("ValidSquare failed")
	}
	// must be true
	if !p.ValidSquare(0, 1) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(1, 0) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(1, 2) {
		t.Error("ValidSquare failed")
	}
	if !p.ValidSquare(2, 1) {
		t.Error("ValidSquare failed")
	}
}

// test the ValidSquare function with a valid square
func TestPolyomino_ValidSquare_04(t *testing.T) {
	p := Polyomino{
		Order: 8,
		M:     3,
		N:     3,
		Shape: [][]bool{
			{true, true, true},
			{true, false, true},
			{true, true, true},
		},
	}
	// must be false
	if p.ValidSquare(0, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(0, 1) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(0, 2) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(1, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(1, 2) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 0) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 1) {
		t.Error("ValidSquare failed")
	}
	if p.ValidSquare(2, 2) {
		t.Error("ValidSquare failed")
	}
	// must be true
	if !p.ValidSquare(1, 1) {
		t.Error("ValidSquare failed")
	}
}
