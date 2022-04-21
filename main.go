package main

import (
	"fmt"
	"go.uber.org/automaxprocs/maxprocs"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var build = "develop"

func main() {
	// Set the correct number of threads for the service
	// based on what is available either by the machine or quotas.
	if _, err := maxprocs.Set(); err != nil {
		fmt.Printf("maxprocs: %w", err)
		os.Exit(1)
	}

	g := runtime.GOMAXPROCS(0)

	log.Printf("starting service build[%s] with CPU[%d]", build, g)
	defer log.Println("service ended")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	log.Println("stopping service")
}
