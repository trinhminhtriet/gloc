package services

import (
	"gloc/internal/domain/entities"
	"os"
)

type Analyzer interface {
	AnalyzeFile(filename string, lang *entities.Language) (*entities.ClocFile, error)
	AnalyzePaths(paths []string) (*entities.Result, error)
}

type analyzer struct {
	langs   map[string]*entities.Language
	parser  Parser
	walker  Walker
	options *options.ClocOptions
}

func NewAnalyzer(langs map[string]*entities.Language, parser Parser, walker Walker, opts *options.ClocOptions) Analyzer {
	return &analyzer{
		langs:   langs,
		parser:  parser,
		walker:  walker,
		options: opts,
	}
}

func (a *analyzer) AnalyzeFile(filename string, lang *entities.Language) (*entities.ClocFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return &entities.ClocFile{Name: filename, Lang: lang.Name}, nil
	}
	defer file.Close()
	return a.parser.Parse(filename, lang, file)
}

func (a *analyzer) AnalyzePaths(paths []string) (*entities.Result, error) {
	// Implementation moved from Processor.Analyze
	total := entities.NewLanguage("TOTAL", nil, nil)
	languages, err := a.walker.Walk(paths, a.langs)
	if err != nil {
		return nil, err
	}
	// Rest of the logic remains similar but uses injected dependencies
}
