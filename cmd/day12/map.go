package main

type Map map[Cave][]Cave

func (m Map) CountPaths() int {
	return m.countPathsRec(CaveStart, nil, Path{CaveStart})
}

func (m Map) countPathsRec(from Cave, blacklist []Cave, pathTaken Path) int {
	if from == CaveEnd {
		log.Debug().WithStringer("path", pathTaken).Message("")
		return 1
	} else if from == CaveStart || from.IsSmall() {
		blacklist = append(blacklist, from)
	}
	var paths int
	for _, to := range m[from] {
		if containsCave(blacklist, to) {
			continue
		}
		//log.Debug().WithStringf("path", "%5s -> %-5s", from, to).Message("")
		paths += m.countPathsRec(to, blacklist, append(pathTaken, to))
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
