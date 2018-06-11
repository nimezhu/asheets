package asheets

import (
	"fmt"
	"log"

	sheets "google.golang.org/api/sheets/v4"
)

/* readSheet is 1-index
 *  TODO: Add More Interfaces Such as translate A,B,C,D or Sheet Header
 *  TODO 101 : Return Column Names
 */
func readSheet(id string, srv *sheets.Service, spreadsheetId string, nameIdx int, valueIdxs []int) ([]string, map[string]interface{}) {
	if len(valueIdxs) == 1 {
		a := make(map[string]interface{})
		m := _readSheetToStringMap(id, srv, spreadsheetId, nameIdx, valueIdxs[0])
		for k, v := range m {
			a[k] = v
		}
		return nil, a
	} else {
		a := make(map[string]interface{})
		h, m := _readSheet(id, srv, spreadsheetId, nameIdx, valueIdxs) //TODO handle header
		for k, v := range m {
			a[k] = v
		}
		return h, a
	}
	return nil, nil
}

/*
 * TODO:
 */
func _readSheet(id string, srv *sheets.Service, spreadsheetId string, nameIdx int, valueIdxs []int) ([]string, map[string][]string) {
	readRange := id + "!A1:ZZ" //TODO: READ FIRST LINE AS NAMES
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	header := make([]string, len(valueIdxs))
	r := make(map[string][]string)
	l := len(valueIdxs)
	if len(resp.Values) > 0 {
		for i0, row := range resp.Values {
			if i0 == 0 {
				//handle header
				for i, k := range valueIdxs {
					header[i] = row[k-1].(string)
				}
			} else {
				s := make([]string, l)
				for i, k := range valueIdxs {
					s[i] = row[k-1].(string)
				}
				r[row[nameIdx-1].(string)] = s
			}
		}
	} else {
		fmt.Print("No data found.")
	}
	return header, r
}

/* readSheetToStringMap is 1-index */
func _readSheetToStringMap(id string, srv *sheets.Service, spreadsheetId string, nameIdx int, valueIdx int) map[string]string {
	readRange := id + "!A2:ZZ"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	fmt.Println("Reading ", id)
	r := make(map[string]string)
	if len(resp.Values) > 0 {
		for _, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			r[row[nameIdx-1].(string)] = row[valueIdx-1].(string)
		}
	} else {
		fmt.Print("No data found.")
	}
	return r
}

/* func readcol */
func ReadCol(id string, srv *sheets.Service, sheetid string, col string) []string {
	readRange := id + "!" + col + "1:" + col
	resp, err := srv.Spreadsheets.Values.Get(sheetid, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	fmt.Println("Reading ", id)
	if len(resp.Values) > 0 {
		r := make([]string, len(resp.Values))
		for i, row := range resp.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			r[i] = row[0].(string)
		}
		return r
	} else {
		return []string{}
	}
}

/*TODO THIS add more interface */
type IndexEntry struct {
	Genome string
	Id     string
	Type   string
	Nc     int
	Vc     []int
}

/*
func readIndex(srv *sheets.Service, spreadsheetId string) []IndexEntry {
	id := "Index"
	readRange := id + "!A2:E"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	a := make([]IndexEntry, len(resp.Values))
	if len(resp.Values) > 0 {
		for i, row := range resp.Values {
			ns := row[3].(string)
			nc, err := strconv.Atoi(ns)
			//TODO
			if err != nil {
				nc = colNameToNumber(ns)
			}
			vs := strings.Split(row[4].(string), ",")
			vc := make([]int, len(vs))
			for i, v := range vs {
				vc[i], err = strconv.Atoi(v)
				if err != nil {
					vc[i] = colNameToNumber(v)
				}
			}
			a[i] = IndexEntry{row[0].(string), row[1].(string), row[2].(string), nc, vc}
		}
	} else {
		fmt.Print("No data found.")
	}
	return a
}
*/
