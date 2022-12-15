package entities

type GenericRabbitMqMessageRequest struct {
	ContentType   string                 `json:"contentType"`
	CorrelationId string                 `json:"correlationId"`
	Domain        string                 `json:"domain"`
	Name          string                 `json:"name"`
	Version       string                 `json:"version"`
	Transient     bool                   `json:"transient"`
	Callback      string                 `json:"callback"`
	RoutingKey    string                 `json:"routingKey"`
	Data          map[string]interface{} `json:"data"`
}
