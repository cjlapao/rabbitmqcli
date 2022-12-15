package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go-rabbitmq/client"
	"github.com/cjlapao/common-go-rabbitmq/constants"
	"github.com/cjlapao/common-go-rabbitmq/sender"
	rabbit_sender "github.com/cjlapao/common-go-rabbitmq/sender"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper/http_helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/rabbitmqcli/entities"
	"github.com/gorilla/mux"
)

func SendRQMPersistentMessageController(w http.ResponseWriter, r *http.Request) {
	var message entities.GenericRabbitMqMessageRequest
	logger := log.Get()
	http_helper.MapRequestBody(r, &message)
	vars := mux.Vars(r)

	connectionString := execution_context.Get().Configuration.GetString(constants.RABBITMQ_CONNECTION_STRING_NAME)
	if connectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.NewApiErrorResponse(entities.InvalidConnectionString, "Connection string cannot be empty or null, please check your environment variables", "")
		response.Log()
		json.NewEncoder(w).Encode(response)
		return
	}

	if client.Get() == nil {
		client.New(connectionString)
	}

	queueName := vars["queueName"]

	msg := entities.GenericRabbitMqMessage{
		MsgCorrelationId: message.CorrelationId,
		MsgName:          message.Name,
		MsgDomain:        message.Domain,
		MsgVersion:       message.Version,
		MsgContentType:   message.ContentType,
		MsgBody:          message.Data,
	}

	var sendMessage sender.SenderMessage

	if message.Transient {
		if message.RoutingKey != "" {
			if message.Callback != "" {
				sendMessage = sender.NewTransientExchangeMessageWithCallback(queueName, message.RoutingKey, message.Callback, msg)
			} else {
				sendMessage = sender.NewTransientExchangeMessage(queueName, message.RoutingKey, msg)
			}
		} else {
			if message.Callback != "" {
				sendMessage = sender.NewTransientQueueMessageWithCallback(queueName, message.Callback, msg)
			} else {
				sendMessage = sender.NewTransientQueueMessage(queueName, msg)
			}
		}
	} else {
		if message.RoutingKey != "" {
			if message.Callback != "" {
				sendMessage = sender.NewExchangeMessageWithCallback(queueName, message.RoutingKey, message.Callback, msg)
			} else {
				sendMessage = sender.NewExchangeMessage(queueName, message.RoutingKey, msg)
			}
		} else {
			if message.Callback != "" {
				sendMessage = sender.NewQueueMessageWithCallback(queueName, message.Callback, msg)
			} else {
				sendMessage = sender.NewQueueMessage(queueName, msg)
			}
		}
	}
	sender := rabbit_sender.New()
	if err := sender.Send(sendMessage); err != nil {
		logger.Exception(err, "sending message")
	}

	json.NewEncoder(w).Encode(message)
}
