# druid-benchmark
Tool built using Go to benchmark druid query latency

## Command Line Arguments
```
-concurrency int
    Number of times each query in specified path gets executed concurrently (default 1)
-output string
  	Path where the report file will be created (default " ")
-password string
  	Enter the password if Basic auth is enabled (default " ")
-path string
   	Enter the path where queries are present (default " ")
-report string
   	Report: csvRaw (default "csvRaw")
-url string
  	Enter the broker query endpoint (default " ")
-username string
  	Enter the username if Basic auth is enabled (default " ")
```

## Run
```
go run cmd/druid-benchmark.go <arguments_here>
```

## Results
- Tool generates a CSV file which can be used to analyze the query latency
- Following is the CSV structure

| QueryName     | Status Code   | Elapsed Time |
| ------------- | ------------- | ------------ |
| file1.json    | 200           | 500          |

- Query Name -> Is the actual `json` filename
- Elapsed Time -> Is in `milliseconds` 
