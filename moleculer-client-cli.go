package main

import (
	"flag"
	"os"
	"strings"
	"time"

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
		NodeID:   "moleculer-go-client-cli",
		LogLevel: moleculer.ErrorLevel,
	}
	service.Start(config)
	//
	go time.AfterFunc(time.Second*1, func() {
		//
		utils.Log.Info("broker.Call demoService.actionA start")
		if res, err := service.Broker.Call("demoService.actionA", map[string]interface{}{
			"arg1": "aaa",
			"arg2": 123,
		}, nil); err != nil {
			utils.Log.Info("broker.Call demoService.actionA end, err: ", err)
		} else {
			utils.Log.Info("broker.Call demoService.actionA end, res: ", res)
		}
		//
		utils.Log.Info("broker.Call demoService.actionB start")
		if res, err := service.Broker.Call("demoService.actionB", map[string]interface{}{
			"arg1": "bbb",
			"arg2": 456,
		}, nil); err != nil {
			utils.Log.Info("broker.Call demoService.actionB end, err: ", err)
		} else {
			utils.Log.Info("broker.Call demoService.actionB end, res: ", res)
		}
		//
		utils.Log.Info("broker.Emit user.create start")
		if err := service.Broker.Emit("user.create", map[string]interface{}{
			"user":   "userA",
			"status": "create",
		}); err != nil {
			utils.Log.Info("broker.Emit user.create end, err: ", err)
		} else {
			utils.Log.Info("broker.Emit user.create end")
		}
		//
		utils.Log.Info("broker.Broadcast user.delete start")
		if err := service.Broker.Broadcast("user.delete", map[string]interface{}{
			"user":   "userB",
			"status": "delete",
		}); err != nil {
			utils.Log.Info("broker.Broadcast user.delete end, err: ", err)
		} else {
			utils.Log.Info("broker.Broadcast user.delete end")
		}
		//
		service.Stop()
		time.Sleep(time.Second * 1)
		os.Exit(0)
	})
	//
	utils.WaitExit(service.Broker)

}
