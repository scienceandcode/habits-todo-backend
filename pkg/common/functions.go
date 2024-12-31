package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

func GenerateHMACUsingSHA256(str, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func VerifyHMAC(str, key, receivedHMAC string) bool {
	calculatedHMAC := GenerateHMACUsingSHA256(str, key)
	return calculatedHMAC == receivedHMAC
}
