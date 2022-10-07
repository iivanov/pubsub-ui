package pub_sub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/iterator"
	"time"
)

type Config struct {
	ProjectId string `yaml:"projectId" required:"true" env:"PUBSUB_PROJECT_ID"`
}

type CreateSubscriptionFormData struct {
	TopicName        string `form:"topicName" binding:"required"`
	SubscriptionName string `form:"subscriptionName" binding:"required"`
}

func NewPubSub(ctx context.Context, config Config) (*PubSub, error) {
	client, clientErr := pubsub.NewClient(ctx, config.ProjectId)
	if clientErr != nil {
		return nil, clientErr
	}
	return &PubSub{
		conf:   config,
		client: client,
	}, nil
}

func (p *PubSub) DeleteSubscription(ctx context.Context, name string) error {
	return p.client.Subscription(name).Delete(ctx)
}

func (p *PubSub) DeleteTopic(ctx context.Context, name string) error {
	t := p.client.Topic(name)
	if t == nil {
		return errors.New("topic not found: " + name)
	}
	return t.Delete(ctx)
}

func (p *PubSub) GetMessages(ctx context.Context, subscriptionName string, messageCount int) ([]*pubsub.Message, error) {
	consumeCtx, cancel := context.WithTimeout(ctx, time.Duration(500)*time.Millisecond)
	defer cancel()
	client, clientErr := pubsub.NewClient(consumeCtx, p.conf.ProjectId)
	if clientErr != nil {
		return nil, clientErr
	}
	defer client.Close()
	if messageCount < 1 {
		return nil, errors.New("messageCount must be greater than 0")
	}
	sub := client.Subscription(subscriptionName)

	sub.ReceiveSettings.MaxOutstandingMessages = messageCount + 1
	result := make([]*pubsub.Message, 0, messageCount)
	err := sub.Receive(consumeCtx, func(ctx context.Context, msg *pubsub.Message) {
		result = append(result, msg)
		msg.Ack()
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PubSub) PublishMessage(ctx context.Context, topicName string, message *pubsub.Message) error {
	t := p.client.Topic(topicName)
	if t == nil {
		return errors.New("topic not found: " + topicName)
	}
	_, err := t.Publish(
		ctx,
		message,
	).Get(ctx)
	return err
}

func (p *PubSub) CreateTopic(ctx context.Context, name string) error {
	_, err := p.client.CreateTopic(ctx, name)
	return err
}

func (p *PubSub) CreateSubscription(ctx context.Context, subscriptionName, topicName string, config pubsub.SubscriptionConfig) error {
	config.Topic = p.client.Topic(topicName)
	_, err := p.client.CreateSubscription(ctx, subscriptionName, config)

	return err
}

func (p *PubSub) SubscriptionsList(ctx context.Context, topicID string) ([]*pubsub.Subscription, error) {
	var subs []*pubsub.Subscription

	it := p.client.Topic(topicID).Subscriptions(ctx)
	for {
		sub, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Next: %v", err)
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

func (p *PubSub) TopicsList(ctx context.Context) ([]*pubsub.Topic, error) {
	var topics []*pubsub.Topic

	it := p.client.Topics(ctx)
	for {
		topic, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Next: %v", err)
		}
		topics = append(topics, topic)
	}

	return topics, nil
}

type PubSub struct {
	conf   Config
	client *pubsub.Client
}
