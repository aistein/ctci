package main

import (
	ch1 "github.com/aistein/ctci/chapter_1_arrays_and_strings"
)

type ChapterName func()
const implementedChapters = []ChapterName{
	ch1.ChapterName,
}

func main() {
	for _, printer := range implementedChapters {
		printer()
	}
}