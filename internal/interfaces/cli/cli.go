package cli

import "log"

func Run() {
	opts := options.NewClocOptions()
	langs := entities.NewDefinedLanguages()
	parser := parser.NewLineParser(opts, bufferpool.NewPool())
	walker := filesystem.NewFileWalker(opts)
	analyzer := services.NewAnalyzer(langs, parser, walker, opts)
	service := application.NewClocService(analyzer, &output.TextFormatter{})
	if err := service.Run([]string{"."}); err != nil {
		log.Fatal(err)
	}
}
