package main

import (
	"github.com/aistein/ctci/ch1"
)

type ChapterName func()
var implementedChapters = []ChapterName{
	ch1.ChapterName,
}

func main() {
	for _, printer := range implementedChapters {
		printer()
	}
}
