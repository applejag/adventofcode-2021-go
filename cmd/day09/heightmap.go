package main

type Heightmap [][]byte

func (h Heightmap) SumRiskLevels() int {
	var sum int
	width, height := len(h), len(h[0])
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if h.IsPointLow(x, y) {
				riskLevel := int(h[x][y]) + 1
				log.Debug().WithInt("x", x).WithInt("y", y).WithInt("risk", riskLevel).
					Message("Low point.")
				sum += riskLevel
			}
		}
	}
	return sum
}

func (h Heightmap) IsPointLow(x, y int) bool {
	val := h[x][y]
	for _, n := range h.GetNeighbors(x, y) {
		if val >= n {
			return false
		}
	}
	return true
}

func (h Heightmap) GetNeighbors(x, y int) []byte {
	var neighbors []byte
	if x > 0 {
		neighbors = append(neighbors, h[x-1][y])
	}
	if x < len(h)-1 {
		neighbors = append(neighbors, h[x+1][y])
	}
	if y > 0 {
		neighbors = append(neighbors, h[x][y-1])
	}
	if y < len(h[x])-1 {
		neighbors = append(neighbors, h[x][y+1])
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
