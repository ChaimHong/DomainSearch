package main

import (
	"log"
	"time"
)

func main() {
	exit := make(chan int)
	chans := make(chan int, 3)
	for i := 0; i < 3; i++ {
		chans <- i
	}

	go func() {
	L:
		for {
			select {
			case v, ok := <-chans:
				if !ok {
					log.Println("end")
					exit <- 1
					break L
				}
				time.Sleep(1e9)
				log.Printf("chan %v", v)
			}
		}
	}()

	<-exit
}
