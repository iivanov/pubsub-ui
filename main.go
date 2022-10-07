package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pubsub-ui/src"
	"pubsub-ui/src/controllers"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint
		cancel()
		fmt.Println("shutting down")
		time.Sleep(time.Second)
		os.Exit(0)
	}()

	err := src.GetContainer(ctx).Invoke(func(topic *controllers.Topic, subscription *controllers.Subscription) error {
		r := gin.Default()
		r.Static("/css", "./templates/static/css")
		r.Static("/img", "./templates/static/img")
		r.Static("/fonts", "./templates/static/fonts")
		r.Static("/js", "./templates/static/js")

		r.LoadHTMLGlob("templates/**/*.gohtml")
		r.GET("/", topic.HandleIndex)
		r.POST("/topic", topic.HandleCreateTopic)
		r.POST("/topic/:topicName/delete", topic.HandleDeleteTopic)
		r.POST("/topic/:topicName/subscription", subscription.HandleCreateSubscription)
		r.POST("/topic/:topicName/subscription/:subscriptionName/delete", subscription.HandleDeleteSubscription)
		r.POST("/topic/:topicName/message/publish", topic.HandlePublishMessage)
		r.GET("/subscription/:subscriptionName/messages", subscription.HandleGetMessages)
		if err := r.Run(":8780"); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
