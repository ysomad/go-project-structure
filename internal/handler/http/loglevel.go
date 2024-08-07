package http

import (
	"log/slog"

	"github.com/ysomad/go-project-structure/internal/gen/server/model"
	"github.com/ysomad/go-project-structure/internal/gen/server/restapi/operations/logging"
	"github.com/ysomad/go-project-structure/internal/slogx"
)

func GetLogLevel(logging.GetLogLevelParams) logging.GetLogLevelResponder {
	return logging.NewGetLogLevelOK().WithPayload(&model.LogLevel{
		Level: slogx.LevelVar.String(),
	})
}

func UpdateLogLevel(p logging.UpdateLogLevelParams, _ any) logging.UpdateLogLevelResponder {
	slogx.LevelVar.Set(slogx.ParseLevel(p.Body.Level))
	slog.InfoContext(p.HTTPRequest.Context(), "updated log level", "new_level", p.Body.Level)
	return logging.NewUpdateLogLevelNoContent()
}
