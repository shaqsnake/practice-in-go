package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Quote
type Quote struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

func main() {
	// Proxy settings
	proxy, _ := url.Parse("http://10.74.149.9:3128")
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	var wg sync.WaitGroup

	for _, symbol := range stockSymbols {
		wg.Add(1)
		go func(symbol string) {
			resp, err := client.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + symbol)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			quote := new(Quote)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
			wg.Done()
		}(symbol)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("Execution time: ", elapsed)

}
