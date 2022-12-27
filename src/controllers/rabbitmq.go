package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go-rabbitmq/sender"
	rabbit_sender "github.com/cjlapao/common-go-rabbitmq/sender"
	"github.com/cjlapao/rabbitmqcli/entities"
)

func SendRQMPersistentQueueMessageController(w http.ResponseWriter, r *http.Request) {
	var message entities.GenericRabbitMqMessageRequest
	base, err := NewBaseController(r, w)
	if err != nil {
		return
	}
	base.MapBody(&message)

	queueName := base.Variables["queueName"]

	msg := entities.GenericRabbitMqMessage{
		MsgCorrelationId: message.CorrelationId,
		MsgName:          message.Name,
		MsgDomain:        message.Domain,
		MsgVersion:       message.Version,
		MsgContentType:   message.ContentType,
		MsgBody:          message.Data,
	}

	var sendMessage sender.Message

	if message.Transient {
		if message.Callback != "" {
			sendMessage = sender.NewTransientQueueMessageWithCallback(queueName, message.Callback, msg)
		} else {
			sendMessage = sender.NewTransientQueueMessage(queueName, msg)
		}
	} else {
		if message.Callback != "" {
			sendMessage = sender.NewQueueMessageWithCallback(queueName, message.Callback, msg)
		} else {
			sendMessage = sender.NewQueueMessage(queueName, msg)
		}
	}

	sender := rabbit_sender.New()
	if err := sender.Send(sendMessage); err != nil {
		base.Logger.Exception(err, "sending message")
	}

	json.NewEncoder(w).Encode(message)
}

func SendRQMPersistentExchangeMessageController(w http.ResponseWriter, r *http.Request) {
	var message entities.GenericRabbitMqMessageRequest
	base, err := NewBaseController(r, w)
	if err != nil {
		return
	}
	base.MapBody(&message)

	queueName := base.Variables["exchangeName"]

	msg := entities.GenericRabbitMqMessage{
		MsgCorrelationId: message.CorrelationId,
		MsgName:          message.Name,
		MsgDomain:        message.Domain,
		MsgVersion:       message.Version,
		MsgContentType:   message.ContentType,
		MsgBody:          message.Data,
	}

	var sendMessage sender.Message

	if message.Transient {
		if message.Callback != "" {
			sendMessage = sender.NewTransientExchangeMessageWithCallback(queueName, message.RoutingKey, message.Callback, msg)
		} else {
			sendMessage = sender.NewTransientExchangeMessage(queueName, message.RoutingKey, msg)
		}
	} else {
		if message.Callback != "" {
			sendMessage = sender.NewExchangeMessageWithCallback(queueName, message.RoutingKey, message.Callback, msg)
		} else {
			sendMessage = sender.NewExchangeMessage(queueName, message.RoutingKey, msg)
		}
	}
	sender := rabbit_sender.New()
	if err := sender.Send(sendMessage); err != nil {
		base.Logger.Exception(err, "sending message")
	}

	json.NewEncoder(w).Encode(message)
}
