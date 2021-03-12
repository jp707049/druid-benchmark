package handlers

import (
	"fmt"
	"github.com/jp707049/druid-benchmark/pkg/parser"
	"net/http"
	"os"
	"sync"
	"time"
)

type Result struct {
	// Query name is the file name. Ex: query.json
	QryName     string
	StatusCode  int
	// Total time taken by the query in Seconds
	ElapsedTime float64
}

// CreateRequest returns a new http.Request which can be executed by client
func CreateRequest(method string, args parser.Arguments, body *os.File) *http.Request {
	req, err := http.NewRequest(method, args.Url, body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if args.Username != " " && args.Password != " " {
		req.SetBasicAuth(args.Username, args.Password)
	}

	return req
}

// CreateClient returns an http.Client.
// It sets IdleConnTimeout and Timeout to 60 seconds
func CreateClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: 60 * time.Second,
		},
		Timeout: 60 * time.Second,
	}

	return client
}

// AsyncHTTP returns a slice of type Result.
// The function executes requests concurrently depending on the concurrency set
func AsyncHTTP(args parser.Arguments, queries []string, client *http.Client) []Result {
	ch := make(chan Result)
	var qry string
	var wg sync.WaitGroup

	for i := 0; i < args.Concurrency; i++ {
		for _, qry = range queries {
			body := ReadQuery(args.Path + qry)
			req := CreateRequest("POST", args, body)
			wg.Add(1)
			go ExecuteQuery(qry, client, req, ch, &wg)
		}
	}

	// close the channel in the background
	go func() {
		wg.Wait()
		close(ch)
	}()

	var results []Result

	// read from channel as they come in until its closed
	// append each result type to the slice of result
	for res := range ch {
		results = append(results, res)
	}

	return results
}

// ExecuteQuery is responsible to make an HTTP request and
// stores the result in a channel ch
func ExecuteQuery(qry string, client *http.Client, req *http.Request, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	st := time.Now()
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if response.StatusCode == 401 {
		fmt.Println(response.Status)
		os.Exit(1)
	}
	elapsed := time.Since(st).Seconds()
	ch <- Result{qry, response.StatusCode, elapsed}
}
