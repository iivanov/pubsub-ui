package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"net/http"
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
		r.Use(cors.Default())
		r.Static("/ui", "./frontend/dist/")
		r.Static("/assets", "./frontend/dist/assets")
		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/ui")
		})

		r.GET("/api/topic", topic.HandleGetTopicList)
		r.POST("/api/topic", topic.HandleCreateTopic)
		r.DELETE("/api/topic/:topicName", topic.HandleDeleteTopic)
		r.POST("/api/topic/:topicName/subscription", subscription.HandleCreateSubscription)
		r.DELETE("/api/topic/subscription/:subscriptionName", subscription.HandleDeleteSubscription)
		r.POST("/api/topic/:topicName/message/publish", topic.HandlePublishMessage)
		r.GET("/api/subscription/:subscriptionName/messages", subscription.HandleGetMessages)
		if err := r.Run(":8780"); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
