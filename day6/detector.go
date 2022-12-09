package day6

import "strings"

// MarkerDetector detects and return the position of the first character of a
// start marker.
type MarkerDetector struct {
	index    int
	capacity int
	window   []string
	seen     map[string]bool
}

func NewMarkerDetector(capacity int) *MarkerDetector {
	md := new(MarkerDetector)
	md.index = 0
	md.capacity = capacity
	md.window = []string{}
	md.seen = map[string]bool{}

	return md
}

// Reads a chunk of signal, and returns the start index of start of packet
// if we find one.
func (md *MarkerDetector) Read(c string) (startIndex int, markerFound bool) {
	split := strings.Split(c, "")

	for _, c := range split {
		// We have already seen this character in the current window, remove
		// characters from start of queue until we remove the repeated character
		for len(md.window) == md.capacity || md.seen[c] {
			last := md.window[0]
			delete(md.seen, last)
			md.window = md.window[1:]
		}

		md.window = append(md.window, c)
		md.seen[c] = true
		md.index++

		// We found the start of packet marker
		if len(md.seen) == md.capacity {
			return md.index, true
		}
	}

	return md.index, false
}
