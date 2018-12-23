package table

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// ReadCSV ... read a CSV file with automated type detection.
// Automated type detection tries to parse the first 20 lines in the file.
func ReadCSV(filename string) (Table, error) {
	tmp := Table{}

	// Read max 20 lines of strings and detect column contents.
	t, err := ioutil.ReadFile(filename)
	DetectTypes(string(t))
	if err != nil {
		return tmp, err
	}
	return tmp, nil
}

// DetectTypes ... find types in a csv and returns a columns specification.
func DetectTypes(csvtext string) {

	tmp := Table{}
	reader := csv.NewReader(strings.NewReader(csvtext))
	columns, err := reader.Read()
	if err == io.EOF {
		log.Println("Unexpected end of file, missing column headers")
	}
	tmp.Columns = make([]Col, 0, len(columns))
	tmp.Values = make([][]Val, len(columns), len(columns))
	for i := range columns {
		tmp.Columns = append(tmp.Columns, Col{StringCol, columns[i]})
		tmp.Values[i] = make([]Val, 0, 0)
	}
	// Add values to each column
	for i := 0; i < 20; i++ {
		datarow, err := reader.Read()
		if err == io.EOF {
			break
		}
		for j := range datarow {
			tmp.Values[j] = append(tmp.Values[j], datarow[j])
		}
	}
	tetsForInteger := func(s []Val) bool {
		return false
	}
	for i := range tmp.Columns {
		if tetsForInteger(tmp.Values[i]) {
			fmt.Println("parsed integers")
		}
	}
}
