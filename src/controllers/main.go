package controllers

import (
	restapi "github.com/cjlapao/common-go-restapi"
)

var listener *restapi.HttpListener

func Init() {
	// userCtx := UserContext{}
	listener = restapi.GetHttpListener()
	listener.AddJsonContent().AddLogger().AddHealthCheck()
	// listener.WithAuthentication(userCtx)

	listener.AddController(SetConfigController, "/config", "POST")
	listener.AddController(GetConfigController, "/config", "GET")

	listener.AddController(SendRQMPersistentQueueMessageController, "/queue/{queueName}", "POST")
	listener.AddController(SendRQMPersistentExchangeMessageController, "/exchange/{exchangeName}", "POST")
	listener.Start()
}
