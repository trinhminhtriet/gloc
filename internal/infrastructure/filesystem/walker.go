package filesystem

import (
	"gloc/internal/domain/entities"
	"gloc/internal/interfaces/options"
)

type Walker interface {
	Walk(paths []string, langs map[string]*entities.Language) (map[string]*entities.Language, error)
}

type fileWalker struct {
	opts *options.ClocOptions
}

func NewFileWalker(opts *options.ClocOptions) Walker {
	return &fileWalker{opts: opts}
}

func (w *fileWalker) Walk(paths []string, langs map[string]*entities.Language) (map[string]*entities.Language, error) {
	// Moved from getAllFiles
}
