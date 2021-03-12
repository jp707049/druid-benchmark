package parser

import (
	"flag"
	"fmt"
	"os"
)

// Arguments is used to store the contents of arguments
// passed via command-line
type Arguments struct {
	Url      string
	Path     string
	Username string
	Password string
	// Number of times all the queries in specified path executed concurrently
	// Example: If you have 2 queries in path and concurrency = 2.
	// Total queries executed will be 4
	Concurrency int
	ReportType string
	ReportPath string
}

// ParseArgs parses the command line arguments
func ParseArgs() Arguments {
	url := flag.String("url", " ", "Enter the broker query endpoint")
	path := flag.String("path", " ", "Enter the path where queries are present")
	username := flag.String("username", " ", "Enter the username if Basic auth is enabled")
	password := flag.String("password", " ", "Enter the password if Basic auth is enabled")
	concurrency := flag.Int("concurrency", 1, "Number of times each query in specified path gets executed concurrently")
	reportType := flag.String("report", "csvRaw", "Report: csvRaw or csvAggregate")
	reportPath := flag.String("output", " ", "Path where the report file will be created")

	flag.Parse()

	if *url == " " || *path == " " {
		fmt.Println("Please add following arguments!!!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	args := Arguments{
		Url:  *url,
		Path: *path,
		Username: *username,
		Password: *password,
		Concurrency: *concurrency,
		ReportType: *reportType,
		ReportPath: *reportPath,
	}

	return args
}
