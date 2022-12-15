package entities

import (
	"encoding/json"

	cryptorand "github.com/cjlapao/common-go-cryptorand"
)

type GenericRabbitMqMessage struct {
	MsgCorrelationId string                 `json:"-"`
	MsgName          string                 `json:"-"`
	MsgDomain        string                 `json:"-"`
	MsgVersion       string                 `json:"-"`
	MsgContentType   string                 `json:"-"`
	MsgBody          map[string]interface{} `json:"-"`
}

func (m GenericRabbitMqMessage) CorrelationID() string {
	if m.MsgCorrelationId != "" {
		return m.MsgCorrelationId
	}

	return cryptorand.GetRandomString(45)
}

func (m GenericRabbitMqMessage) Domain() string {
	if m.MsgDomain != "" {
		return m.MsgDomain
	}
	return "Tester"
}

func (m GenericRabbitMqMessage) Name() string {
	if m.MsgName != "" {
		return m.MsgName
	}
	return "HelloWorld"
}

func (m GenericRabbitMqMessage) Version() string {
	if m.MsgVersion != "" {
		return m.MsgVersion
	}
	return "1.0"
}

func (m GenericRabbitMqMessage) ContentType() string {
	if m.MsgContentType != "" {
		return m.MsgContentType
	}
	return "application/json"
}

func (m GenericRabbitMqMessage) Body() []byte {
	j, err := json.Marshal(m.MsgBody)
	if err != nil {
		return nil
	}

	return j
}
