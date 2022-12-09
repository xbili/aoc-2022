package day6

import (
	"fmt"
	"os"
)

func Run(filename string, windowSize int) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sr := NewStreamReader(f)
	md := NewMarkerDetector(windowSize)

	for {
		chunk, ok := sr.Read()
		if !ok {
			break
		}

		index, found := md.Read(chunk)
		if found {
			fmt.Println(index)
			break
		}
	}
}
