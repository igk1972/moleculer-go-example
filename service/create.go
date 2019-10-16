package service

import (
	"../utils"

	moleculer "github.com/roytan883/moleculer-go"
	"github.com/roytan883/moleculer-go/protocol"
)

func Create(name string) moleculer.Service {
	service := moleculer.Service{
		ServiceName: name,
		//
		Actions: map[string]moleculer.RequestHandler{
			//
			"actionA": func(req *protocol.MsRequest) (interface{}, error) {
				utils.Log.Info("run actionA, req.Params = ", req.Params)
				data := map[string]interface{}{
					"res1": "AAA",
					"res2": 123,
				}
				return data, nil // return nil, errors.New("test return error in actionA")
			},
			//
			"actionB": func(req *protocol.MsRequest) (interface{}, error) {
				utils.Log.Info("run actionB, req.Params = ", req.Params)
				data := map[string]interface{}{
					"res1": "BBB",
					"res2": 456,
				}
				return data, nil // return nil, errors.New("test return error in actionB")
			},
			//
			"bench": func(req *protocol.MsRequest) (interface{}, error) {
				// utils.Log.Info("run actionB, req.Params = ", req.Params)
				data := map[string]interface{}{
					"res1": "CCC",
					"res2": 789,
				}
				return data, nil // return nil, errors.New("test return error in actionB")
			},
			//
		},
		//
		Events: map[string]moleculer.EventHandler{
			//
			"user.create": func(req *protocol.MsEvent) {
				utils.Log.Info("run onEventUserCreate, req.Data = ", req.Data)
			},
			//
			"user.delete": func(req *protocol.MsEvent) {
				utils.Log.Info("run onEventUserDelete, req.Data = ", req.Data)
			},
			//
		},
		//
	}
	//
	return service
}
