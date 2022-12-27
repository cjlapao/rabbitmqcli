package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go-rabbitmq/client"
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/helper/http_helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/rabbitmqcli/entities"
	"github.com/cjlapao/rabbitmqcli/environment"
	"github.com/gorilla/mux"
)

type BaseController struct {
	Request        *http.Request
	Writer         http.ResponseWriter
	Context        *execution_context.Context
	Environment    *environment.Environment
	RabbitMqClient *client.RabbitMQClient
	Logger         *log.Logger
	Variables      map[string]string
}

func NewBaseController(r *http.Request, w http.ResponseWriter) (*BaseController, *entities.ApiErrorResponse) {
	base := BaseController{
		Context:     execution_context.Get(),
		Environment: environment.Get(),
		Logger:      log.Get(),
		Request:     r,
		Writer:      w,
	}

	base.Variables = mux.Vars(r)

	connectionString := base.Environment.ConnectionString()
	if connectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.NewApiErrorResponse(entities.InvalidConnectionString, "Connection string cannot be empty or null, please check your environment variables", "")
		response.Log()
		json.NewEncoder(w).Encode(response)
		return nil, &response
	}
	base.Logger.Info(connectionString)

	client := client.New(connectionString)

	if client == nil {
		response := base.HandleError(entities.InvalidClient, "There was an error trying to get a valid client to connect to the rabbitmq server")
		return nil, &response
	}

	return &base, nil
}

func (base *BaseController) HandleError(code entities.ApiResponseErrorCode, msg string) entities.ApiErrorResponse {
	return base.HandleUrlError(code, msg, "")
}

func (base *BaseController) HandleUrlError(code entities.ApiResponseErrorCode, msg string, url string) entities.ApiErrorResponse {
	base.Writer.WriteHeader(http.StatusBadRequest)
	response := entities.NewApiErrorResponse(code, msg, url)
	response.Log()
	json.NewEncoder(base.Writer).Encode(response)

	return response
}

func (base *BaseController) MapBody(dest interface{}) error {
	return http_helper.MapRequestBody(base.Request, dest)
}
