package main

func NewWindowedScanner(is IntScanner, windowSize int) IntScanner {
	return &windowedScanner{
		windows: make([]int, windowSize),
		is:      is,
	}
}

type windowedScanner struct {
	windows []int
	is      IntScanner
	scans   int
}

func (ws *windowedScanner) Scan() bool {
	for {
		if !ws.is.Scan() {
			return false
		}

		num := ws.is.Int()
		low := ws.scans - len(ws.windows) + 1
		if low < 0 {
			low = 0
		}
		ws.windows[ws.scans%len(ws.windows)] = 0
		for i := ws.scans; i >= low; i-- {
			ws.windows[i%len(ws.windows)] += num
		}
		log.Debug().WithInt("#", ws.scans).
			WithInt("num", num).
			WithStringf("windows", "%v", ws.windows).
			Messagef("%T", ws)

		ws.scans++
		if ws.scans >= len(ws.windows) {
			break
		}
	}
	return true
}

func (ws *windowedScanner) Int() int {
	res := ws.windows[ws.scans%len(ws.windows)]
	log.Debug().WithInt("out", res).Messagef("%T", ws)
	return res
}

func (ws *windowedScanner) Err() error {
	return ws.is.Err()
}
