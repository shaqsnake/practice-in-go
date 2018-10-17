package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
		// time.Sleep(1 * time.Second)
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
