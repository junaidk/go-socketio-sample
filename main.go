package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	sockMap := make(map[string]string)

	router := gin.New()
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		s.Emit("connection", "have "+s.ID())
		return nil
	})

	server.OnEvent("/", "username", func(s socketio.Conn, msg string) {
		fmt.Println("username:", msg)
		fmt.Println("sid:", s.ID())

		// join a room, each client will have its own room
		s.Join(msg)
		sockMap[msg] = s.ID()
		s.Emit("connection", "have "+s.ID())

	})

	//server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
	//	s.SetContext(msg)
	//	return "recv " + msg
	//})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	router.GET("/ping", func(c *gin.Context) {
		name := c.Query("name")
		fmt.Println(name)
		out := server.BroadcastToRoom("/", name, "message", "hello there")
		fmt.Println("braodcast:", out)
	})

	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	router.Run()
}
