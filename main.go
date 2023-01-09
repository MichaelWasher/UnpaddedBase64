package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	input_str := os.Args[1]
	decoded, err := base64.RawURLEncoding.DecodeString(input_str)

	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println(string(decoded))
}
