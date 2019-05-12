package sheet

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"

	"usg_bot/config"
	"usg_bot/entity"
)

// Sheet represents SpreadSheet as Repository.
type Sheet struct {
	SheetID string
	client  *sheets.Service
}

// NewSheet is a function to initialize Sheet Client
func NewSheet(sheetID string) *Sheet {
	client, err := sheets.New(getClient())
	if err != nil {
		panic(err)
	}
	return &Sheet{
		SheetID: sheetID,
		client:  client,
	}
}

func getClient() *http.Client {
	file := config.GetConfig().GoogleJWTPath
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}

	cfg, err := google.JWTConfigFromJSON(
		[]byte(data), sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to read secret: %v", err)
	}

	return cfg.Client(context.Background())
}

// Rows is a function to fetch data in range
func (s *Sheet) Rows(rng *entity.Range) ([][]string, error) {
	resp, err := s.client.Spreadsheets.Values.Get(s.SheetID, rng.String()).Do()
	if err != nil {
		return nil, err
	}

	if len(resp.Values) == 0 {
		return make([][]string, 0), nil
	}

	ret := make([][]string, len(resp.Values))
	for _, row := range resp.Values {
		rs := make([]string, len(row))
		for _, cell := range row {
			rs = append(rs, fmt.Sprintf("%v", cell))
		}
		ret = append(ret, rs)
	}

	return ret, nil
}

// Search is a function to search data where the key exists in row
// key must exist in first row in given range.
func (s *Sheet) Search(rng *entity.Range, key string, needle string) ([]string, error) {
	rows, err := s.Rows(rng)
	if err != nil {
		return nil, err
	}
	keyIdx := -1
	for i, row := range rows {
		for k, str := range row {
			if keyIdx == -1 && str == key {
				keyIdx = i
			}
			if k == i && str == needle {
				// found
				return row, nil
			}
		}
	}
	if keyIdx == -1 {

	}
	return nil, nil
}

// Append is a function to add a row to the last row of sheet.
func (*Sheet) Append(tab string, row []string) [][]string {
	return nil
}
