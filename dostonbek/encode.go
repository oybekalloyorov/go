package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	message := "hello,go (*w3hu%#"
	demoBase64(message)
	demoHex(message)
}

func demoBase64(message string) {
	fmt.Println("--------Demo encoding base64--------")
	fmt.Println("plaintext:")
	fmt.Println(message)
	encoding := base64.StdEncoding.EncodeToString([]byte(message))
	fmt.Println("base64 msg:")
	fmt.Println(encoding)
	decoding, _ := base64.StdEncoding.DecodeString(encoding)
	fmt.Println("decoding base64 msg:")
	fmt.Println(string(decoding))
}

func demoHex(message string) {
	fmt.Println("--------Demo encoding Hex--------")
	fmt.Println("plaintext:")
	fmt.Println(message)
	encoding := hex.EncodeToString([]byte(message))
	fmt.Println("Hex msg:")
	fmt.Println(encoding)
	decoding, _ := hex.DecodeString(encoding)
	fmt.Println("decoding Hex msg:")
	fmt.Println(string(decoding))
}