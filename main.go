package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //global var for using a waitGroup

func main() {

	var url string
	var threads int
	var token string

	flag.StringVar(&url, "u", "", "request for pushing")
	flag.IntVar(&threads, "t", 10, "amount of threads")
	flag.StringVar(&token, "h", "", "Authorization token")

	flag.Parse()

	start := time.Now()

	fmt.Println("Processing...")

	for i := 0; i < threads; i++ {

		go makeRequest(url, token)
		wg.Add(1) // tell 1 gourutine checking a status of another gourutine

	}

	wg.Wait() // this guy waiting when a guy above tell us that everything done
	elapsed := time.Since(start)

	fmt.Println("------------------------------------")

	fmt.Println("Target:", url)
	fmt.Println("Amount of requests:", threads)
	fmt.Println("Program execuded for:", elapsed)

}

func makeRequest(url string, token string) {

	defer wg.Done() //this guy will send info for (wg.Add(1)) about status each gourutine

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	res, _ := client.Do(req)

	req.Header = http.Header {
		
		"Authorization": {token},
		"Content": {"application/json"},

	}
	if err != nil {
		log.Fatalln()
	} 

	fmt.Println("Response code of request: ", res.StatusCode)
}
