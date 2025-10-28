package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	demoHash_md5()
	demoHash_sha256()
	demoHashKey("mykey","Hello world, go!")
}


func demoHash_md5() {
	fmt.Println("--------Demo encoding hash using md5--------")
	message := "Hello world, go!"
	fmt.Println("plaintext:")
	fmt.Println(message)
	h := md5.New()
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)
}

func demoHash_sha256() {
	fmt.Println("--------Demo encoding hash using sha256--------")
	message := "Hello world, go!"
	fmt.Println("plaintext:")
	fmt.Println(message)
	h := sha256.New()
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)
}

func demoHashKey(key,message string) {
	fmt.Println("--------Demo encoding hash with key: HMAC and sha256--------")
	fmt.Println("key:")
	fmt.Println(key)
	fmt.Println("plaintext:")
	fmt.Println(message)
	hmacKey := []byte(key)
	h := hmac.New(sha256.New, hmacKey)
	h.Write([]byte(message))
	hash_message := hex.EncodeToString(h.Sum(nil))
	fmt.Println("hashing message:")
	fmt.Println(hash_message)
}