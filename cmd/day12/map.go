package main

import "github.com/jilleJr/adventofcode-2021-go/internal/common"

type Map map[Cave][]Cave

func (m Map) CountPaths() int {
	if common.Part2 {
		return m.countPathsRecPart2(CaveStart, Path{CaveStart}, false)
	} else {
		return m.countPathsRecPart1(CaveStart, Path{CaveStart})
	}
}

func (m Map) countPathsRecPart1(from Cave, pathTaken Path) int {
	if from == CaveEnd {
		log.Debug().WithStringer("path", pathTaken).Message("")
		return 1
	}
	var paths int
	for _, to := range m[from] {
		if to.IsSmall() && containsCave(pathTaken, to) {
			continue
		}
		//log.Debug().WithStringf("path", "%5s -> %-5s", from, to).Message("")
		paths += m.countPathsRecPart1(to, append(pathTaken, to))
	}
	return paths
}

func (m Map) countPathsRecPart2(from Cave, pathTaken Path, revisitSmall bool) int {
	if from == CaveEnd {
		log.Debug().WithStringer("path", pathTaken).Message("")
		return 1
	}
	var paths int
	for _, to := range m[from] {
		if to == CaveStart {
			continue
		}
		if to.IsSmall() && containsCave(pathTaken, to) {
			if !revisitSmall {
				paths += m.countPathsRecPart2(to, append(pathTaken, to), true)
			}
			continue
		}
		//log.Debug().WithStringf("path", "%5s -> %-5s", from, to).Message("")
		paths += m.countPathsRecPart2(to, append(pathTaken, to), revisitSmall)
	}
	return paths
}

func containsCave(slice []Cave, elem Cave) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
