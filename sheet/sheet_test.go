package sheet_test

import (
	"fmt"
	"testing"

	"usg_bot/sheet"
)

var (
	sheetID = "1cxUFPdLfWQZeB_348B9OybhtotN_d0z9jGnYy3fzhJI"
)

func Test_Rows(t *testing.T) {
	s := sheet.NewSheet(sheetID)
	rs, err := s.Rows("sample", "A1", "D5")

	if err != nil {
		t.Fatalf("failed to fetch sheets.Row %v", err)
	}

	if len(rs) == 0 {
		t.Error("empty sheet rows")
	}

	t.Log(fmt.Sprintf("rows: %v", rs))
}
