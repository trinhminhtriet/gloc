package output

import (
	"bytes"
	"os"
	"testing"

	"gloc/internal/domain/entities"
)

func TestXMLFormatter_Format(t *testing.T) {
	// Setup
	result := &entities.Result{
		Total: &entities.Language{
			Name:     "TOTAL",
			Total:    2,
			Code:     100,
			Comments: 10,
			Blanks:   20,
		},
		Languages: map[string]*entities.Language{
			"Go": {
				Name:     "Go",
				Files:    []string{"main.go"},
				Code:     80,
				Comments: 8,
				Blanks:   15,
			},
			"Python": {
				Name:     "Python",
				Files:    []string{"script.py"},
				Code:     20,
				Comments: 2,
				Blanks:   5,
			},
		},
	}
	formatter := &XMLFormatter{}

	// Redirect stdout for testing
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run
	err := formatter.Format(result)
	w.Close()
	os.Stdout = oldStdout
	if err != nil {
		t.Fatalf("Format failed: %v", err)
	}

	// Read output
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// Verify
	expected := bytes.Contains([]byte(output), []byte(`<language name="Go"`)) &&
		bytes.Contains([]byte(output), []byte(`<language name="Python"`)) &&
		bytes.Contains([]byte(output), []byte(`<total sum_files="2"`)) &&
		bytes.Contains([]byte(output), []byte(`code="100"`))
	if !expected {
		t.Errorf("Output mismatch:\nGot:\n%s\nExpected to contain Go, Python, total, and code=100", output)
	}
}
