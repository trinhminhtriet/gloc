package output

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gloc/internal/domain/entities"
)

// OutputFormatter defines the interface for formatting analysis results.
type OutputFormatter interface {
	Format(result *entities.Result) error
}

// TextFormatter formats analysis results as a text table.
type TextFormatter struct{}

// Format outputs the analysis results in a text table format.
func (f *TextFormatter) Format(result *entities.Result) error {
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

	// Calculate column widths
	const (
		langHeader    = "Language"
		filesHeader   = "files"
		blankHeader   = "blank"
		commentHeader = "comment"
		codeHeader    = "code"
	)
	maxLangWidth := len(langHeader)
	for _, lang := range languages {
		if len(lang.Name) > maxLangWidth {
			maxLangWidth = len(lang.Name)
		}
	}

	// Build the output
	var builder strings.Builder
	builder.WriteString(strings.Repeat("-", 80) + "\n")
	builder.WriteString(fmt.Sprintf("%-*s %10s %12s %12s %12s\n",
		maxLangWidth, langHeader, filesHeader, blankHeader, commentHeader, codeHeader))
	builder.WriteString(strings.Repeat("-", 80) + "\n")

	// Add language rows
	for _, lang := range languages {
		builder.WriteString(fmt.Sprintf("%-*s %10d %12d %12d %12d\n",
			maxLangWidth, lang.Name, len(lang.Files), lang.Blanks, lang.Comments, lang.Code))
	}

	// Add total row
	builder.WriteString(strings.Repeat("-", 80) + "\n")
	builder.WriteString(fmt.Sprintf("%-*s %10d %12d %12d %12d\n",
		maxLangWidth, "TOTAL", result.Total.Total, result.Total.Blanks, result.Total.Comments, result.Total.Code))
	builder.WriteString(strings.Repeat("-", 80) + "\n")

	// Write to stdout
	_, err := fmt.Fprint(os.Stdout, builder.String())
	return err
}
