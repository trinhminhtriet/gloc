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

	for _, lang := range result.Languages {
		fmt.Println(lang)
	}
	fmt.Println(result.Total)
	fmt.Printf("%+v", result)
}
