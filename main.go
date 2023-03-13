package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// 消息处理
func handleSend(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.POST("/send", handleSend)

	return e
}

func router02() http.Handler {
	return websocket.Handler(handleWS)
}

func handleWS(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		// websock连接失败处理
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

	}
}

func main() {
	server01 := &http.Server{
		Addr:         ":1234",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":1010",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
