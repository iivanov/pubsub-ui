package src

import (
	"context"
	"go.uber.org/dig"
	"pubsub-ui/config"
	"pubsub-ui/src/controllers"
	"pubsub-ui/src/pub_sub"
)

var ContainerRegistry = []interface{}{
	pub_sub.NewPubSub,
	config.NewConfig,
	controllers.NewTopic,
	controllers.NewSubscription,
	func(config config.Config) pub_sub.Config { return config.PubSub },
}

func GetContainer(ctx context.Context) *dig.Container {
	container := dig.New()

	if err := container.Provide(func() context.Context { return ctx }); err != nil {
		panic(err)
	}

	for _, v := range ContainerRegistry {
		if err := container.Provide(v); err != nil {
			panic(err)
		}
	}

	return container
}
