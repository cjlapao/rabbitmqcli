package environment

import (
	"github.com/cjlapao/common-go-rabbitmq/constants"
	"github.com/cjlapao/common-go/execution_context"
)

var globalEnvironment *Environment

type Environment struct {
	context          *execution_context.Context
	connectionString string
}

func new() *Environment {
	return &Environment{
		context: execution_context.Get(),
	}
}

func Get() *Environment {
	if globalEnvironment != nil {
		return globalEnvironment
	}

	globalEnvironment = new()

	return globalEnvironment
}

func (env *Environment) ConnectionString() string {
	if env.connectionString == "" {
		env.connectionString = env.context.Configuration.GetString(constants.RABBITMQ_CONNECTION_STRING_NAME)
	}

	return env.connectionString
}

func (env *Environment) SetConnectionString(value string) {
	env.connectionString = value
}
