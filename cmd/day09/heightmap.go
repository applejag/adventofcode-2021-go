package main

type Point struct {
	x int
	y int
}

type Heightmap [][]byte

func (h Heightmap) SumRiskLevels() int {
	var sum int
	for _, p := range h.GetLowPoints() {
		riskLevel := int(h.Val(p)) + 1
		log.Debug().WithInt("x", p.x).WithInt("y", p.y).WithInt("risk", riskLevel).
			Message("Low point.")
		sum += riskLevel
	}
	return sum
}

func (h Heightmap) Val(p Point) byte {
	return h[p.x][p.y]
}

func (h Heightmap) GetBasinSizes() []int {
	var sizes []int
	for _, p := range h.GetLowPoints() {
		size := h.GetBasinSize(p)
		log.Debug().WithInt("x", p.x).
			WithInt("y", p.y).
			WithInt("size", size).
			Message("Basin size.")
		sizes = append(sizes, size)
	}
	return sizes
}

func (h Heightmap) GetBasinSize(p Point) int {
	points := map[Point]Empty{p: {}}
	h.addBasinNeighborsToMapRec(p, points)
	return len(points)
}

func (h Heightmap) addBasinNeighborsToMapRec(p Point, m map[Point]Empty) {
	for _, n := range h.GetNeighbors(p) {
		if h.Val(n) == 9 {
			continue
		}
		if _, ok := m[n]; !ok {
			m[n] = Empty{}
			h.addBasinNeighborsToMapRec(n, m)
		}
	}
}

func (h Heightmap) GetLowPoints() []Point {
	var lowPoints []Point
	width, height := len(h), len(h[0])
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if h.IsPointLow(Point{x, y}) {
				lowPoints = append(lowPoints, Point{x, y})
			}
		}
	}
	return lowPoints
}

func (h Heightmap) IsPointLow(p Point) bool {
	val := h.Val(p)
	for _, n := range h.GetNeighbors(p) {
		if val >= h.Val(n) {
			return false
		}
	}
	return true
}

func (h Heightmap) GetNeighbors(p Point) []Point {
	var neighbors []Point
	if p.x > 0 {
		neighbors = append(neighbors, Point{p.x - 1, p.y})
	}
	if p.x < len(h)-1 {
		neighbors = append(neighbors, Point{p.x + 1, p.y})
	}
	if p.y > 0 {
		neighbors = append(neighbors, Point{p.x, p.y - 1})
	}
	if p.y < len(h[p.x])-1 {
		neighbors = append(neighbors, Point{p.x, p.y + 1})
	}
	return neighbors
}

func ParseHeightmap(strs []string) (Heightmap, error) {
	var lines [][]byte
	for _, str := range strs {
		var line []byte
		for _, r := range str {
			d, err := parseDigit(r)
			if err != nil {
				return Heightmap{}, err
			}
			line = append(line, d)
		}
		lines = append(lines, line)
	}
	return Heightmap(lines), nil
}
