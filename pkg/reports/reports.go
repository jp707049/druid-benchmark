package reports

import (
	"github.com/jp707049/druid-benchmark/pkg/handlers"
	"github.com/jp707049/druid-benchmark/pkg/parser"
)

// ReportCSV write the result to a CSV file
func ReportCSV(args parser.Arguments, results []handlers.Result) {
	// Create a CSV file
	file := handlers.CreateFile(args)
	defer file.Close()

	// Write results to CSV file
	handlers.WriterCSV(file, results)
}