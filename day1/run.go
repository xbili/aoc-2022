package day1

import (
	"fmt"
	"os"
)

func Run(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ecr := NewElvesCaloriesReader(f)
	fmt.Println(FindRichestElves(3, ecr))
}
