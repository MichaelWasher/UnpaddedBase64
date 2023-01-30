package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

var VERSION string = "0.001"
var APP_NAME string = "ubase64d"

func print_help(){
	fmt.Printf("%s: version %s\n", APP_NAME, VERSION)
	fmt.Printf("%s <unpadded-base64-content>\n", APP_NAME)
}	
	
func main() {
	if len(os.Args) != 2{
		print_help()
		fmt.Println("Please ensure that only a single value is provided to the tool.")
		os.Exit(0)
	}
	input_str := os.Args[1]
	decoded, err := base64.RawURLEncoding.DecodeString(input_str)

	if err != nil {
		fmt.Println("decode error:", err)
		fmt.Println("please ensure valid unpadded base64 content is provided.")
		return
	}

	fmt.Println(string(decoded))
}
