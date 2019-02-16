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

	return cfg.Client(context.Background())
}

func sheetRange(tab string, from string, to string) string {
	return fmt.Sprintf("%v!%v:%v", tab, from, to)
}

// Rows is a function to fetch data in range
func (s *Sheet) Rows(tab string, from string, to string) ([][]string, error) {
	resp, err := s.client.Spreadsheets.Values.Get(s.SheetID, sheetRange(tab, from, to)).Do()
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
func (*Sheet) Search(tab string, key string, row string) []string {
	return nil
}

// Append is a function to add a row to the last row of sheet.
func (*Sheet) Append(tab string, data []string) [][]string {
	return nil
}
