package application

import (
	"gloc/internal/domain"
	"gloc/internal/interfaces"
)

type ClocService struct {
	analyzer domain.Analyzer
	output   interfaces.OutputFormatter
}

func NewClocService(analyzer domain.Analyzer, output interfaces.OutputFormatter) *ClocService {
	return &ClocService{analyzer: analyzer, output: output}
}

func (s *ClocService) Run(paths []string) error {
	result, err := s.analyzer.AnalyzePaths(paths)
	if err != nil {
		return err
	}
	return s.output.Format(result)
}
