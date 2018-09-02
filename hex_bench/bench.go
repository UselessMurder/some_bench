package main

import (
	"fmt"
	"time"
	"encoding/hex"
	"io/ioutil"
)


func main() {
	content, _ := ioutil.ReadFile("./bin")
	start := time.Now()
	str := hex.EncodeToString(content)
//	fmt.Println(str)
	content, _ = hex.DecodeString(str)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}
