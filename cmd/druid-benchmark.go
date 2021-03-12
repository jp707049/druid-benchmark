package main

import (
	"fmt"
	"github.com/jp707049/druid-benchmark/pkg/handlers"
	"github.com/jp707049/druid-benchmark/pkg/parser"
	"github.com/jp707049/druid-benchmark/pkg/reports"
)

func main() {
	args := parser.ParseArgs()

	queries := handlers.GetQueryList(args)

	client := handlers.CreateClient()

	results := handlers.AsyncHTTP(args, queries, client)

	if args.ReportType == "csvRaw" {
		reports.ReportCSV(args, results)
	} else {
		fmt.Println("Please choose a valid report type: raw or aggr")
	}

}