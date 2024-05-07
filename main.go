package main

import (
	"context"
	"flag"
	"log"
	"main/conf"
	"main/rest"
	"main/tele"
	"os"
	"os/signal"
	"time"
)

var WaitInputForCommand string

func main() {

	assertStrong(conf.Init(), "conf.Init()")
	assertStrong(tele.Init(), "tele.Init()")
	assertStrong(rest.Init(), "rest.Init()")

	configObserveSignals()
	os.Exit(0)

}

func assertStrong(err error, name string) {
	if err != nil {
		log.Panicf("main()."+name+" error: %v", err)
		os.Exit(1)
	}
}

func configObserveSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	_, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	log.Println("shutting down")
}
