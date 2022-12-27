package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/cjlapao/common-go/helper/http_helper"
	"github.com/cjlapao/rabbitmqcli/entities"
	"github.com/cjlapao/rabbitmqcli/environment"
)

func SetConfigController(w http.ResponseWriter, r *http.Request) {
	var config entities.ConfigRequest

	err := http_helper.MapRequestBody(r, &config)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.NewApiErrorResponse(entities.InvalidConnectionString, "Connection string cannot be empty or null, please check your environment variables", "")
		response.Log()
		json.NewEncoder(w).Encode(response)
		return
	}

	if config.ConnectionString == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := entities.NewApiErrorResponse(entities.BadRequest, "Connection string cannot be empty or null, please check your environment variables", "")
		response.Log()
		json.NewEncoder(w).Encode(response)
		return
	}

	env := environment.Get()
	env.SetConnectionString(config.ConnectionString)

	w.WriteHeader(http.StatusAccepted)
}

func GetConfigController(w http.ResponseWriter, r *http.Request) {
	env := environment.Get()
	config := entities.ConfigRequest{
		ConnectionString: env.ConnectionString(),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(config)
}
