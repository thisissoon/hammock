package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	client := flag.String("client", "", "Client ID")
	key := flag.String("key", "", "Client Secret")
	data := flag.String("data", "", "Request Body")

	flag.Parse()

	mac := hmac.New(sha256.New, []byte(*key))
	mac.Write([]byte(*data))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	fmt.Println(signature)
	header := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", *client, signature)))

	fmt.Println(header)
}
