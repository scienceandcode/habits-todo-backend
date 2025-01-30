package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"net/mail"
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

func EncryptAES(plainText string) string {
	key := GetEnv("APP_ENCRYPTION_CYPHER_TEXT")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Error while creating new cipher: %v", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalf("Error while creating new cipher: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

	return base64.URLEncoding.EncodeToString(cipherText)
}

func DecryptAES(cipherText string) (string, error) {
	key := GetEnv("APP_ENCRYPTION_CYPHER_TEXT")
	cipherTextBytes, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(cipherTextBytes) < aes.BlockSize {
		return "", err
	}

	iv := cipherTextBytes[:aes.BlockSize]
	cipherTextBytes = cipherTextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return string(cipherTextBytes), nil
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
