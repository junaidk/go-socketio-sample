package main

import (
	"bufio"
	"fmt"
	"github.com/zhouhui8915/go-socket.io-client"
	"log"
	"os"
)

func main() {

	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	opts.Query["user"] = "user"
	opts.Query["pwd"] = "pass"
	opts.Query["name"] = "junaid"

	uri := "http://localhost:8080"

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
		return
	}

	client.On("error", func() {
		log.Printf("on error\n")
	})
	client.On("connection", func(msg string) {
		log.Printf("on connect: %v\n", msg)

	})
	client.On("message", func(msg string) {
		log.Printf("on message: %v\n", msg)
	})

	client.On("reply", func(msg string) {
		log.Printf("on reply: %v\n", msg)
	})

	client.On("disconnection", func() {
		log.Printf("on disconnect\n")
	})

	out := client.Emit("username", os.Args[1])

	fmt.Println(out)

	reader := bufio.NewReader(os.Stdin)
	for {
		data, _, _ := reader.ReadLine()
		command := string(data)

		client.Emit("msg", "chat")
		log.Printf("send message:%v\n", command)
	}
}
