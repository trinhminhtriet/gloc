package entities

type ClocFile struct {
	Code     int32
	Comments int32
	Blanks   int32
	Name     string
	Lang     string
}

type ClocFiles []ClocFile

// Sorting methods (SortByName, SortByComments, etc.) moved here
