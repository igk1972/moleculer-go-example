package main

import (
	"flag"
	"strings"

	"./service"
	"./utils"

	nats "github.com/nats-io/go-nats"
	moleculer "github.com/roytan883/moleculer-go"
)

func main() {
	urls := flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	flag.Parse()
	//
	hosts := strings.Split(*urls, ",")
	utils.Log.Printf("hosts '%v'", hosts)
	//
	config := &moleculer.ServiceBrokerConfig{
		NatsHost: hosts,
		LogLevel: moleculer.ErrorLevel,
		NodeID:   "moleculer-go-service",
		Services: map[string]moleculer.Service{
			"demoService": service.Create("demoService"),
		},
	}
	service.Run(config)
}
