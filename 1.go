package main

import (
    b64 "encoding/base64"
	// "encoding/hex"
    "fmt"
)

func main() {

    data := "54686520736f6c7574696f6e2069733a2074686174776173746f6f65617379"

    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()
}