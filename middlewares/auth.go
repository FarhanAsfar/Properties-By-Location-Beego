package middlewares

import (
	"encoding/json"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func ValidateAPIKey(ctx *context.Context) {
	expectedKey, _ := web.AppConfig.String("x_api_key")

	requestedKey := ctx.Input.Header("X-Api-Key")

	if requestedKey == "" || requestedKey != expectedKey {
		ctx.Output.SetStatus(401)
		ctx.Output.Header("Content-Type", "application/json")

		payload, _ := json.Marshal(map[string]string{"error": "unauthorized"})
		ctx.WriteString(string(payload))
	}
}
