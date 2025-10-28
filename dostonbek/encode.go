package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	message := "hello,go (*w3hu%#"
	demoBase64(message)
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