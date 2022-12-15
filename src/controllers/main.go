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

	listener.AddController(SendRQMPersistentMessageController, "/queue/{queueName}", "POST")
	listener.Start()
}
