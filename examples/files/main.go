package main

import (
	"fmt"

	"github.com/trinhminhtriet/gloc"
)

func main() {
	languages := gloc.NewDefinedLanguages()
	options := gloc.NewClocOptions()
	paths := []string{
		".",
	}

	processor := gloc.NewProcessor(languages, options)
	result, err := processor.Analyze(paths)
	if err != nil {
		fmt.Printf("gloc fail. error: %v\n", err)
		return
	}

	for _, item := range result.Files {
		fmt.Println(item)
	}
	fmt.Println(result.Total)
	fmt.Printf("%+v", result)
}
