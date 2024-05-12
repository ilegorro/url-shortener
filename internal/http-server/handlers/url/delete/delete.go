package delete

import (
	"log/slog"
	"net/http"
	"url-shortener/internal/http-server/handlers"
	"url-shortener/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// URLDeleter is an interface for deleting url by alias
//
//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=URLDeleter
type URLDeleter interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDeleter URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, handlers.Error("incorrect request"))

			return
		}

		log.Info("request", slog.String("alias", alias))

		err := urlDeleter.DeleteURL(alias)
		if err != nil {
			log.Error("failed to delete url", logger.Err(err))
			render.JSON(w, r, handlers.Error("internal error"))

			return
		}

		render.JSON(w, r, handlers.OK())
	}
}
