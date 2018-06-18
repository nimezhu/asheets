package asheets

import (
	"fmt"

	sheets "google.golang.org/api/sheets/v4"
)

type TableMap struct {
	ColIds   []string
	RowIds   []string
	RowValue map[string][]string
}

/*ReadSheet */
func ReadSheet(title string, srv *sheets.Service, spreadsheetId string, idCol string) (*TableMap, error) {
	idIndex := ColNameToNumber(idCol) - 1
	readRange := title + "!A1:ZZ"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		//log.Fatalf("Unable to retrieve data from sheet. %v", err)
		return nil, err
	}
	r := make(map[string][]string)
	var header []string
	var rowid []string
	var l int
	if len(resp.Values) > 0 {
		rowid = make([]string, len(resp.Values)-1)
		for i0, row := range resp.Values {
			if i0 == 0 {
				l = len(row)
				header = make([]string, l)
				for i, _ := range row {
					header[i] = row[i].(string)
				}
			} else {
				s := make([]string, l)
				for i, _ := range row {
					s[i] = row[i].(string)
				}
				rowid[i0-1] = row[idIndex].(string)
				r[row[idIndex].(string)] = s
			}
		}
	} else {
		fmt.Print("No data found.")
	}
	m := &TableMap{header, rowid, r}
	return m, nil
}
