package tournament

// Sheet represents SpreadSheet Repository.
type Sheet interface {
	Rows(name string, from string, to string) ([][]string, error)

	Search(tab string, key string, row string) []string
	Append(tab string, data []string) [][]string
}
