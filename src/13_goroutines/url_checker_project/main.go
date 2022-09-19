package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"  
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {	

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is Down!", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code : %d \n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			body, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1]
			file += ".txt"

			fmt.Printf("Writing response body to %s \n", file)

			err = ioutil.WriteFile(file, body, 0664)

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println(strings.Repeat("-", 50))

	wg.Done()	
}

func main() {

	var wg sync.WaitGroup

	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}

	wg.Wait()

	fmt.Println("DONE")
}
