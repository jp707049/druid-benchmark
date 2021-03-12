package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/jp707049/druid-benchmark/pkg/parser"
	"io/ioutil"
	"os"
	"strconv"
)

// GetQueryList Returns s slice of file names (queries) from given directory path
func GetQueryList(args parser.Arguments) []string {
	var queries []string

	files, err := ioutil.ReadDir(args.Path)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range files {
		queries = append(queries, f.Name())
	}

	return queries
}

// ReadQuery returns a new os.File type which will be used as a POST body
// in an HTTP request
func ReadQuery(absPath string) *os.File {
	body, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

// CreateFile returns a new file in path specified by cfg.ReportPath
func CreateFile(args parser.Arguments) *os.File {
	if args.ReportPath == " " {
		args.ReportPath = "result.csv"
	} else {
		args.ReportPath = args.ReportPath + "result.csv"
	}

	file, err := os.Create(args.ReportPath)
	if err != nil {
		fmt.Println(err)
	}

	return file
}

// WriterCSV writes the results into the CSV file
func WriterCSV(file *os.File, results []Result) {
	writer := csv.NewWriter(file)
	for _, r := range results {
		err := writer.Write([]string{r.QryName, strconv.Itoa(r.StatusCode),
			strconv.FormatFloat(r.ElapsedTime, 'f', 5, 64)})
		if err != nil {
			fmt.Println(err)
		}
	}
	defer writer.Flush()
}
