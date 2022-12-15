package startup

import (
	"github.com/cjlapao/common-go/execution_context"
	"github.com/cjlapao/common-go/security/encryption"
	"github.com/cjlapao/common-go/service_provider"
	controller "github.com/cjlapao/rabbitmqcli/controllers"
)

var providers = service_provider.Get()

func Init() {
	ctx := execution_context.Get()
	ctx.WithDefaultAuthorization()
	ctx.Authorization.WithAudience("carloslapao.com")
	kv := ctx.Authorization.KeyVault
	kv.WithBase64HmacKey("HMAC", providers.Configuration.GetString("JWT_HMAC_PRIVATE_KEY"), encryption.Bit256)
	kv.SetDefaultKey("RSA512_4096")

	controller.Init()
}
