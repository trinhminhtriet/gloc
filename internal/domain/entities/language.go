package entities

import (
	"regexp"
	"sort"
)

type Language struct {
	Name              string
	LineComments      []string
	RegexLineComments []*regexp.Regexp
	MultiLines        [][]string
	Files             []string
	Code              int32
	Comments          int32
	Blanks            int32
	Total             int32
}

func NewLanguage(name string, lineComments []string, multiLines [][]string) *Language {
	return &Language{
		Name:         name,
		LineComments: lineComments,
		MultiLines:   multiLines,
	}
}

type Languages []Language

// Sorting methods (SortByName, SortByCode, etc.) moved here as methods on Languages
func (ls Languages) SortByCode() {
	sort.Slice(ls, func(i, j int) bool {
		return ls[i].Code > ls[j].Code
	})
}
