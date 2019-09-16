package main

import (
	"context"
	"log"
	"time"
)

const (
	duration = 2 * time.Second
)

func main() {
	log.Println("start")

	ctx, cancel := context.WithCancel(context.Background())

	cnt := 0
B:
	for {
		select {
		case t := <-time.After(duration):
			log.Printf("after %v second\n", t)
			cnt++
			if cnt == 4 {
				cancel()
			}
		case <-ctx.Done():
			log.Println("canceled..")
			break B
		}
	}

	log.Println("end")
}
