package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const (
	envURL   = "REQUEST_URL"
	envCount = "REQUEST_COUNT"
)

var count *int
var url *string
var client *http.Client

func main() {
	var wg sync.WaitGroup
	for i := 0; i < *count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			status, err := getURL(url)

			if err != nil {
				log.Fatal(err)
			}
			log.Output(0, "url: "+*url+",status: "+status)
		}()
	}

	wg.Wait()
}
func getURL(url *string) (string, error) {
	resp, err := client.Get(*url)
	if err != nil {
		return "", err
	}
	return resp.Status, err
}
func init() {
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	flag := flag.NewFlagSet("f", flag.PanicOnError)
	c, err := strconv.Atoi(os.Getenv(envCount))
	if err != nil {
		log.Output(0, "Info: "+err.Error())
	}
	count = flag.Int("count", c, "numbers to iterate")
	url = flag.String("url", os.Getenv(envURL), "url to make requests")
	flag.Parse(os.Args[1:])
}
