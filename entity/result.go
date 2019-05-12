package entity

// Result represents Game Result.
type Result struct {
	Category   string
	Block      string
	GameNumber int
	Win        string
	Lose       string
	WinScore   int
	LoseScore  int
}

// ToRow is a function to make row from *Result
func (t *Result) ToRow() []string {
	return nil
}
