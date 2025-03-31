package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"gloc/internal/application"
	"gloc/internal/domain/entities"
	"gloc/internal/domain/services"
	"gloc/internal/infrastructure/filesystem"
	"gloc/internal/infrastructure/parser"
	"gloc/internal/interfaces/options"
	"gloc/internal/interfaces/output"
	"gloc/internal/pkg/bufferpool"

	"github.com/golangci/golangci-lint/pkg/config"
)

const (
	versionString = "gloc v0.1.0"
	usage         = `gloc - A blazing-fast LOC (Lines of Code) counter in Go

Usage:
  gloc [flags] [paths...]

Flags:
`
)

// Config holds CLI configuration parsed from flags.
type Config struct {
	Version        bool
	ShowLang       bool
	OutputFormat   string
	ExcludeExts    []string
	Debug          bool
	SkipDuplicated bool
	Paths          []string
}

func main() {
	// Parse configuration from CLI flags
	config, err := parseConfig()
	if err != nil {
		log.Fatalf("Failed to parse configuration: %v", err)
	}

	// Handle early exits
	if config.Version {
		fmt.Println(versionString)
		return
	}

	// Initialize options
	opts := initOptions(config)

	// Load defined languages
	definedLangs := entities.NewDefinedLanguages()

	// Handle show-lang flag
	if config.ShowLang {
		fmt.Println(definedLangs.GetFormattedString())
		return
	}

	// Initialize application with dependencies
	app, err := initApplication(opts, definedLangs.Langs)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Run the application
	if err := app.Run(config.Paths); err != nil {
		log.Fatalf("Error running gloc: %v", err)
	}
}

// parseConfig parses CLI flags into a Config struct.
func parseConfig() (*Config, error) {
	config := &Config{}

	// Define flags with custom usage
	flag.BoolVar(&config.Version, "version", false, "Print the version and exit")
	flag.BoolVar(&config.ShowLang, "show-lang", false, "Show supported languages and exit")
	flag.StringVar(&config.OutputFormat, "output", "text", "Output format (text, json, xml)")
	flag.StringVar(&config.ExcludeExts, "exclude-ext", "", "Comma-separated list of extensions to exclude")
	flag.BoolVar(&config.Debug, "debug", false, "Enable debug logging")
	flag.BoolVar(&config.SkipDuplicated, "skip-duplicated", false, "Skip duplicated files based on MD5 hash")

	// Set custom usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s", usage)
		flag.PrintDefaults()
	}

	flag.Parse()

	// Parse paths from remaining arguments
	config.Paths = flag.Args()
	if len(config.Paths) == 0 {
		config.Paths = []string{"."}
	}

	// Convert exclude-ext to slice
	if config.ExcludeExts != "" {
		config.ExcludeExts = splitTrim(config.ExcludeExts, ",")
	}

	return config, nil
}

// initOptions initializes ClocOptions from Config.
func initOptions(config *Config) *options.ClocOptions {
	opts := options.NewClocOptions()
	opts.Debug = config.Debug
	opts.SkipDuplicated = config.SkipDuplicated
	for _, ext := range config.ExcludeExts {
		opts.ExcludeExts[ext] = struct{}{}
	}
	return opts
}

// initApplication initializes the application with all dependencies.
func initApplication(opts *options.ClocOptions, langs map[string]*entities.Language) (*application.ClocService, error) {
	// Initialize buffer pool
	bufferPool := bufferpool.NewPool()

	// Initialize parser
	lineParser := parser.NewLineParser(opts, bufferPool)

	// Initialize file walker
	fileWalker := filesystem.NewFileWalker(opts)

	// Initialize analyzer
	analyzer := services.NewAnalyzer(langs, lineParser, fileWalker, opts)

	// Initialize output formatter
	formatter, err := newOutputFormatter(opts, strings.ToLower(strings.TrimSpace(config.OutputFormat)))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize output formatter: %w", err)
	}

	// Initialize cloc service
	return application.NewClocService(analyzer, formatter), nil
}

// newOutputFormatter creates an OutputFormatter based on the specified format.
func newOutputFormatter(opts *options.ClocOptions, format string) (output.OutputFormatter, error) {
	switch format {
	case "text":
		return &output.TextFormatter{}, nil
	case "json":
		return &output.JSONFormatter{}, nil
	case "xml":
		return &output.XMLFormatter{}, nil
	default:
		return nil, fmt.Errorf("unsupported output format: %s", format)
	}
}

// splitTrim splits a string by a separator and trims whitespace from each part.
func splitTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
