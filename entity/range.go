package entity

import "fmt"

// Range represents SpreadSheet Range
type Range struct {
	// Sheet name
	tab string
	// Start point of Range
	from string
	// End point of Range
	to string
}

func NewRange(tab string, from string, to string) *Range {
	return &Range{
		tab:  tab,
		from: from,
		to:   to,
	}
}

//Tab is a function to return Tab which this Range targets.
func (r *Range) Tab() string {
	return r.tab
}

// From is a function to return From of this Range.
func (r *Range) From() string {
	return r.from
}

// To is a function to return To of this Range.
func (r *Range) To() string {
	return r.to
}

func (r *Range) String() string {
	return fmt.Sprintf("%v!%v:%v", r.tab, r.from, r.to)
}
