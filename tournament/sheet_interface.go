package tournament

import (
	"usg_bot/entity"
)

// Sheet represents SpreadSheet Repository.
type Sheet interface {
	// Range struct may be better to be added.
	Rows(rng entity.Range) ([][]string, error)

	Search(rng entity.Range, key string, row string) ([]string, error)
	Append(tab string, data []string) [][]string
}
