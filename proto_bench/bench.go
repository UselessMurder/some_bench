package main

import (
	"time"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	msg "./tutorial"
	"fmt"
)

func main() {
	content, _ := ioutil.ReadFile("./bin")
	message := &msg.Msg {
		Name: "Hello",
		Id: 0,
		Content: content,
	}
	start := time.Now()
	data, _ := proto.Marshal(message)

	newMessage := &msg.Msg{}

	proto.Unmarshal(data, newMessage)
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

