package service

import (
	"../utils"

	moleculer "github.com/roytan883/moleculer-go"
)

var Broker *moleculer.ServiceBroker

func Start(config *moleculer.ServiceBrokerConfig) (err error) {
	if broker, err := moleculer.NewServiceBroker(config); err != nil {
		utils.Log.Fatalf("NewServiceBroker err: %v", err)
	} else {
		Broker = broker
		if err := broker.Start(); err != nil {
			return err
		} else {
			utils.Log.Info("broker.Start")
		}
	}
	return
}

func Stop() (err error) {
	if Broker != nil {
		utils.Log.Info("broker.Stop")
		return Broker.Stop()
	}
	return
}

func Run(config *moleculer.ServiceBrokerConfig) (err error) {
	if err := Start(config); err != nil {
		return err
	}
	utils.WaitExit(Broker)
	return
}
