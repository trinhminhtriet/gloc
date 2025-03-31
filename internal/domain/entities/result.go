package entities

type Result struct {
	Total         *Language
	Files         map[string]*ClocFile
	Languages     map[string]*Language
	MaxPathLength int
}
