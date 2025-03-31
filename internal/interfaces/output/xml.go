package output

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"

	"gloc/internal/domain/entities"
)

// XMLFormatter formats analysis results as XML.
type XMLFormatter struct{}

// XMLLanguages represents the XML structure for languages.
type XMLLanguages struct {
	XMLName   xml.Name      `xml:"languages"`
	Languages []XMLLanguage `xml:"language"`
	Total     XMLTotal      `xml:"total"`
}

// XMLLanguage represents a single language in XML.
type XMLLanguage struct {
	Name       string `xml:"name,attr"`
	FilesCount int32  `xml:"files_count,attr"`
	Code       int32  `xml:"code,attr"`
	Comments   int32  `xml:"comment,attr"`
	Blanks     int32  `xml:"blank,attr"`
}

// XMLTotal represents the total statistics in XML.
type XMLTotal struct {
	SumFiles int32 `xml:"sum_files,attr"`
	Code     int32 `xml:"code,attr"`
	Comments int32 `xml:"comment,attr"`
	Blanks   int32 `xml:"blank,attr"`
}

// Format outputs the analysis results in XML format.
func (f *XMLFormatter) Format(result *entities.Result) error {
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
		return languages[i].Name < languages[j].Name // Sort alphabetically by name
	})

	// Build XML structure
	xmlLangs := XMLLanguages{
		Languages: make([]XMLLanguage, len(languages)),
		Total: XMLTotal{
			SumFiles: result.Total.Total,
			Code:     result.Total.Code,
			Comments: result.Total.Comments,
			Blanks:   result.Total.Blanks,
		},
	}

	for i, lang := range languages {
		xmlLangs.Languages[i] = XMLLanguage{
			Name:       lang.Name,
			FilesCount: int32(len(lang.Files)),
			Code:       lang.Code,
			Comments:   lang.Comments,
			Blanks:     lang.Blanks,
		}
	}

	// Marshal to XML
	output, err := xml.MarshalIndent(xmlLangs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal XML: %w", err)
	}

	// Write XML header and content to stdout
	_, err = fmt.Fprintf(os.Stdout, "%s\n%s\n", xml.Header, string(output))
	return err
}
