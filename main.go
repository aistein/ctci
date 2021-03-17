package main

import (
	"github.com/aistein/ctci/ch1"
)

// ChapterName prints the name of the CTCI chapter corresponding to this package
type ChapterName func()

var implementedChapters = []ChapterName{
	ch1.ChapterName,
}

func main() {
	for _, printer := range implementedChapters {
		printer()
	}
}
