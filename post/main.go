package main

import (
	"post/handler"
	"post/subscriber"

	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	post "post/proto/post"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.post"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// db := config.NewGormDatabase()

	// Register Handler
	post.RegisterPostHandler(service.Server(), new(handler.Post))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.post", service.Server(), new(subscriber.Post))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
