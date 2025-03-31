package output

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"gloc/internal/domain/entities"
)

// JSONFormatter formats analysis results as JSON.
type JSONFormatter struct{}

// JSONLanguages represents the JSON structure for languages.
type JSONLanguages struct {
	Languages []JSONLanguage `json:"languages"`
	Total     JSONTotal      `json:"total"`
}

// JSONLanguage represents a single language in JSON.
type JSONLanguage struct {
	Name       string `json:"name"`
	FilesCount int32  `json:"files"`
	Code       int32  `json:"code"`
	Comments   int32  `json:"comment"`
	Blanks     int32  `json:"blank"`
}

// JSONTotal represents the total statistics in JSON.
type JSONTotal struct {
	FilesCount int32 `json:"files"`
	Code       int32 `json:"code"`
	Comments   int32 `json:"comment"`
	Blanks     int32 `json:"blank"`
}

// Format outputs the analysis results in JSON format.
func (f *JSONFormatter) Format(result *entities.Result) error {
	if result == nil {
		return fmt.Errorf("result is nil")
	}

	// Prepare sorted languages
	languages := make([]*entities.Language, 0, len(result.Languages))
	for _, lang := range result.Languages {
		if len(lang.Files) > 0 { // Only include languages with files
			languages = append(languages, lang)
		}
	}
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Code > languages[j].Code // Sort by code lines descending
	})

	// Build JSON structure
	jsonLangs := JSONLanguages{
		Languages: make([]JSONLanguage, len(languages)),
		Total: JSONTotal{
			FilesCount: result.Total.Total,
			Code:       result.Total.Code,
			Comments:   result.Total.Comments,
			Blanks:     result.Total.Blanks,
		},
	}

	for i, lang := range languages {
		jsonLangs.Languages[i] = JSONLanguage{
			Name:       lang.Name,
			FilesCount: int32(len(lang.Files)),
			Code:       lang.Code,
			Comments:   lang.Comments,
			Blanks:     lang.Blanks,
		}
	}

	// Marshal to JSON
	output, err := json.MarshalIndent(jsonLangs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Write to stdout
	_, err = fmt.Fprintf(os.Stdout, "%s\n", output)
	return err
}
