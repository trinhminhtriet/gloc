package parser

import (
	"io"

	"gloc/internal/domain/entities"
	"gloc/internal/interfaces/options"
)

type Parser interface {
	Parse(filename string, lang *entities.Language, reader io.Reader) (*entities.ClocFile, error)
}

type lineParser struct {
	opts *options.ClocOptions
	pool *bufferpool.Pool
}

func NewLineParser(opts *options.ClocOptions, pool *bufferpool.Pool) Parser {
	return &lineParser{opts: opts, pool: pool}
}

func (p *lineParser) Parse(filename string, lang *entities.Language, reader io.Reader) (*entities.ClocFile, error) {
	// Moved from AnalyzeReader
}
