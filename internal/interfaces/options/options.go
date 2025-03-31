package options

type ClocOptions struct {
	Debug          bool
	SkipDuplicated bool
	ExcludeExts    map[string]struct{}
	// ... other fields
}

func NewClocOptions() *ClocOptions {
	return &ClocOptions{
		ExcludeExts: make(map[string]struct{}),
	}
}
