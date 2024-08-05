package http

import (
	"net/http"

	"github.com/ysomad/go-project-structure/internal/gen/server/model"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/health"
)

func Ping(health.PingParams) health.PingResponder {
	return health.NewPingOK().WithPayload(&model.PingResponse{
		Status: http.StatusText(http.StatusOK),
	})
}
