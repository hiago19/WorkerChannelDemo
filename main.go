package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerID, x)
		time.Sleep(time.Second)
	}

}

func main() {
	canal := make(chan int)
	qtdWorker := 5

	for i := range qtdWorker {
		go worker(i, canal)
	}

	for i := range 10 {
		canal <- i
	}
}
