package gloc

import "regexp"

type ClocOptions struct {
	Debug          bool
	SkipDuplicated bool
	ExcludeExts    map[string]struct{}
	IncludeLangs   map[string]struct{}
	ReNotMatch     *regexp.Regexp
	ReMatch        *regexp.Regexp
	ReNotMatchDir  *regexp.Regexp
	ReMatchDir     *regexp.Regexp
	Fullpath       bool

	OnCode    func(line string)
	OnBlank   func(line string)
	OnComment func(line string)
}

func NewClocOptions() *ClocOptions {
	return &ClocOptions{
		Debug:          false,
		SkipDuplicated: false,
		ExcludeExts:    make(map[string]struct{}),
		IncludeLangs:   make(map[string]struct{}),
	}
}
