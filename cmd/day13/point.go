package main

type Point struct {
	x, y int
}

func (p Point) Folded(f Fold) Point {
	if f.Axis == AxisX {
		return p.FoldedX(f.Pos)
	} else {
		return p.FoldedY(f.Pos)
	}
}

func (p Point) FoldedX(x int) Point {
	if p.x <= x {
		return p
	}
	return Point{x + x - p.x, p.y}
}

func (p Point) FoldedY(y int) Point {
	if p.y <= y {
		return p
	}
	return Point{p.x, y + y - p.y}
}
