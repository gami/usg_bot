package entity

// Team represents Team in touranemnt.
type Team struct {
	Category    string
	Block       string
	BlockNumber int
	Name        string
	LeaderName  string
}

// ToRow is a function to make row from *Team
func (t *Team) ToRow() []string {
	return nil
}
