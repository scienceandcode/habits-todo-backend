package common

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func GetEnv(env string) string {
	value, isSet := os.LookupEnv(env)

	if !isSet {
		log.Panicf("environment variable not set: %s", env)
	}

	return value
}

func WaitOsInterruption() {
	var waitGroup sync.WaitGroup

	osInterrupt := make(chan os.Signal, 1)
	signal.Notify(osInterrupt, os.Interrupt)

	syscallSigterm := make(chan os.Signal, 1)
	signal.Notify(syscallSigterm, syscall.SIGTERM)

	waitGroup.Add(1)

	go func() {
		<-osInterrupt
		defer waitGroup.Done()
	}()

	go func() {
		<-syscallSigterm
		defer waitGroup.Done()
	}()

	waitGroup.Wait()
}
