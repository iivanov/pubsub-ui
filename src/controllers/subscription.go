package controllers

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pubsub-ui/src/pub_sub"
	"time"
)

func NewSubscription(pubSub *pub_sub.PubSub) *Subscription {
	return &Subscription{
		pubSub: pubSub,
	}
}

type Subscription struct {
	pubSub *pub_sub.PubSub
}

type SubscriptionFormData struct {
	SubscriptionName          string `json:"name" binding:"required"`
	AckDeadline               int    `json:"ackDeadline" binding:"required"`
	EnableExactlyOnceDelivery bool   `json:"enableExactlyOnceDelivery"`
	EnableMessageOrdering     bool   `json:"enableMessageOrdering"`
	PublishEndpoint           string `json:"publishEndpoint"`
}

func (mp *Subscription) HandleDeleteSubscription(c *gin.Context) {
	subscriptionName := c.Param("subscriptionName")
	if err := mp.pubSub.DeleteSubscription(c, subscriptionName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (mp *Subscription) HandleGetMessages(c *gin.Context) {
	subscriptionName := c.Param("subscriptionName")
	if subscriptionName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subscription name is required"})
		return
	}

	messages, err := mp.pubSub.GetMessages(c, subscriptionName, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type MessageView struct {
		Message    string
		Attributes string
	}

	messagesView := make([]MessageView, len(messages))
	for i, message := range messages {
		messagesView[i].Message = string(message.Data)
		atr, attrErr := json.Marshal(message.Attributes)
		if attrErr != nil {
			log.Println("Error marshaling attributes: ", attrErr)
		}
		messagesView[i].Attributes = string(atr)
	}

	c.JSON(http.StatusOK, gin.H{"messages": messagesView, "subscriptionName": subscriptionName})
}

func (mp *Subscription) HandleCreateSubscription(c *gin.Context) {
	topicName := c.Param("topicName")
	if topicName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "topic name is empty"})
		return
	}
	var data SubscriptionFormData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config := pubsub.SubscriptionConfig{
		AckDeadline: time.Duration(data.AckDeadline) * time.Second,
		PushConfig: pubsub.PushConfig{
			Endpoint: data.PublishEndpoint,
		},
		EnableExactlyOnceDelivery: data.EnableExactlyOnceDelivery,
		EnableMessageOrdering:     data.EnableMessageOrdering,
	}

	if err := mp.pubSub.CreateSubscription(c, data.SubscriptionName, topicName, config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
