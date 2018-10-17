package main

import (
	"sync"
	"time"
	"encoding/csv"
	"os"
	"strconv"
	"fmt"
)

func main() {
	start := time.Now()

	extractChannel := make(chan *Order)
	transformChannel := make(chan *Order)
	doneChannel := make(chan bool)

	go extract(extractChannel)
	go transform(extractChannel, transformChannel)
	go load(transformChannel, doneChannel)

	<-doneChannel

	fmt.Println(time.Since(start))
}

type Product struct {
	PartNumber string
	UnitCost float64
	UnitPrice float64
}

type Order struct {
	CustomerNumber int
	PartNumber string
	Quantity int

	UnitCost float64
	UnitPrice float64
}

func extract(extractChannel chan *Order) {
	f, _ := os.Open("./orders.txt")
	defer f.Close()

	r := csv.NewReader(f)
	for record, err := r.Read(); err == nil; record, err = r.Read() {
		order := new(Order)
		order.CustomerNumber, _ = strconv.Atoi(record[0])
		order.PartNumber = record[1]
		order.Quantity, _ = strconv.Atoi(record[2])
		extractChannel <- order
	}

	close(extractChannel)
}

func transform(extractChannel, transformChannel chan *Order) {
	f, _ := os.Open("./productList.txt")
	defer f.Close()

	r := csv.NewReader(f)
	records, _ := r.ReadAll()
	productList := make(map[string]*Product)
	for _, record := range records {
		product := new(Product)
		product.PartNumber = record[0]
		product.UnitCost, _ = strconv.ParseFloat(record[1], 64)
		product.UnitPrice, _ = strconv.ParseFloat(record[2], 64)
		productList[product.PartNumber] = product
	}

	var wg sync.WaitGroup
	for order := range extractChannel {
		// time.Sleep(3 * time.Millisecond)
		wg.Add(1)
		go func(order *Order) {
			order.UnitCost = productList[order.PartNumber].UnitCost
			order.UnitPrice = productList[order.PartNumber].UnitPrice	
			transformChannel <- order
			wg.Done()
		}(order)
	}
	wg.Wait()

	close(transformChannel)
}

func load(transfromChannel chan *Order, doneChannel chan bool) {
	f, _ := os.Create("./dest.txt")
	defer f.Close()

	fmt.Fprintf(f, "%20s%15s%12s%12s%15s%15s\n", "Part Number", "Quantity", "Unit Cost", "Unit Price", "Total Cost", "Total Price")

	var wg sync.WaitGroup
	for order := range transfromChannel {
		// time.Sleep(1 * time.Millisecond)
		wg.Add(1)
		go func(order *Order) {
			fmt.Fprintf(f, "%20s %15d %12.2f %12.2f %15.2f %15.2f\n", order.PartNumber, order.Quantity, order.UnitCost, order.UnitPrice, order.UnitCost * float64(order.Quantity), order.UnitPrice * float64(order.Quantity))
			wg.Done()
		}(order)
	}
	wg.Wait()
		
	doneChannel <- true
}