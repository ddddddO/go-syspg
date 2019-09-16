package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.Println("strat")

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		log.Println("sub finished")
		time.Sleep(2 * time.Second)
		cancel()
	}()

	<-ctx.Done()
	log.Println("end")
}
