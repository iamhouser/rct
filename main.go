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

	flag.StringVar(&url, "u", "", "request for pushing")
	flag.IntVar(&threads, "t", 10, "amount of threads")

	flag.Parse()

	start := time.Now()

	fmt.Println("Processing...")

	for i := 0; i < threads; i++ {

		go makeRequest(url)
		wg.Add(1) // tell 1 gourutine checking a status of another gourutine

	}

	wg.Wait() // this guy waiting when a guy above tell us that everything done
	elapsed := time.Since(start)

	fmt.Println("------------------------------------")

	fmt.Println("Target:", url)
	fmt.Println("Amount of requests:", threads)
	fmt.Println("Program execuded for:", elapsed)

}

func makeRequest(url string) {

	defer wg.Done() //this guy will send info for (wg.Add(1)) about status each gourutine

	resp, err := http.Get(url)
	resp.Header.Set("Authorisation", "Bearer")
	if err != nil {
		log.Fatalln()
	} 

	fmt.Println("Response code of request: ", resp.StatusCode)
}
