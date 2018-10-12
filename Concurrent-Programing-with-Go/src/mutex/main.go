package main

import (
	"fmt"
	"runtime"
	// "sync"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	f, _ := os.Create("./my.log")
	f.Close()

	logCh := make(chan string, 50)

	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./my.log", os.O_APPEND, os.ModeAppend)

				logTime := time.Now().Format(time.RFC3339Nano)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()
	// mutex := new(sync.Mutex)
	mutex := make(chan bool, 1)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			// mutex.Lock()
			mutex <- true
			go func() {
				// fmt.Printf("%d + %d = %d\n", i, j, i+j)
				// mutex.Unlock()
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}

	fmt.Scanln()
}
