package controllers

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pubsub-ui/src/pub_sub"
	"strings"
)

func NewTopic(pubSub *pub_sub.PubSub) *Topic {
	return &Topic{
		pubSub: pubSub,
	}
}

type Topic struct {
	pubSub *pub_sub.PubSub
}

func (mp *Topic) HandleIndex(c *gin.Context) {
	topics, err := mp.pubSub.TopicsList(c)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": err.Error()})
		return
	}

	type SubscriptionView struct {
		ID                    string
		Name                  string
		PublishEndpoint       string
		EnableMessageOrdering bool
		ExactlyOnceDelivery   bool
		AckDeadlineSeconds    int
	}

	type TopicForView struct {
		Name          string
		ID            string
		Subscriptions []SubscriptionView
	}

	topicsForView := make([]TopicForView, len(topics))

	for i, topic := range topics {
		topicsForView[i].Name = strings.Replace(topic.ID(), ".", "_", -1)
		topicsForView[i].ID = topic.ID()

		subscriptions, subsErr := mp.pubSub.SubscriptionsList(c, topic.ID())
		if subsErr != nil {
			c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": err.Error()})
			return
		}
		topicsForView[i].Subscriptions = make([]SubscriptionView, len(subscriptions))
		for j, subscription := range subscriptions {
			topicsForView[i].Subscriptions[j].Name = subscription.ID()
			topicsForView[i].Subscriptions[j].ID = subscription.ID()

			subConfig, confErr := subscription.Config(c)
			if confErr != nil {
				log.Printf("Error getting subscription config: %v", confErr)
			}
			topicsForView[i].Subscriptions[j].PublishEndpoint = subConfig.PushConfig.Endpoint
			topicsForView[i].Subscriptions[j].EnableMessageOrdering = subConfig.EnableMessageOrdering
			topicsForView[i].Subscriptions[j].ExactlyOnceDelivery = subConfig.EnableExactlyOnceDelivery
			topicsForView[i].Subscriptions[j].AckDeadlineSeconds = int(subConfig.AckDeadline.Seconds())
		}
	}
	c.HTML(http.StatusOK, "index.gohtml", gin.H{"title": "Main Page", "topics": topicsForView})
}

func (mp *Topic) HandlePublishMessage(c *gin.Context) {
	topicName := c.Param("topicName")
	if topicName == "" {
		c.HTML(http.StatusBadRequest, "error.gohtml", gin.H{"error": "topic name is empty"})
		return
	}

	type Message struct {
		Message    string `form:"message" required:"true"`
		Attributes string `form:"attributes"`
	}

	var msg Message
	if err := c.ShouldBind(&msg); err != nil {
		c.HTML(http.StatusBadRequest, "error.gohtml", gin.H{"error": err.Error()})
		return
	}

	attributes := make(map[string]string)
	if msg.Attributes != "" {
		if err := json.Unmarshal([]byte(msg.Attributes), &attributes); err != nil {
			c.HTML(http.StatusBadRequest, "error.gohtml", gin.H{"error": err.Error()})
			return
		}
	}

	publishErr := mp.pubSub.PublishMessage(
		c,
		topicName,
		&pubsub.Message{
			Data:       []byte(msg.Message),
			Attributes: attributes,
		},
	)
	if publishErr != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": publishErr.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (mp *Topic) HandleCreateTopic(c *gin.Context) {
	topicName := c.PostForm("name")
	if topicName == "" {
		c.HTML(http.StatusBadRequest, "error.gohtml", gin.H{"error": "topic name is empty"})
		return
	}
	if err := mp.pubSub.CreateTopic(c, topicName); err != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (mp *Topic) HandleDeleteTopic(c *gin.Context) {
	topicName := c.Param("topicName")
	if topicName == "" {
		c.HTML(http.StatusBadRequest, "error.gohtml", gin.H{"error": "topic name is empty"})
		return
	}
	if err := mp.pubSub.DeleteTopic(c, topicName); err != nil {
		c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
