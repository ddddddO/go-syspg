package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP)

	log.Println("wait..ctlC or ")

	for {
		select {
		case <-signals:
			log.Println("end")
			break
			/*		default:
					log.Println("d")
			*/
		}
	}
}
