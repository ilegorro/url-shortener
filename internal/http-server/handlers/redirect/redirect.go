package redirect

import (
	"errors"
	"log/slog"
	"net/http"
	"url-shortener/internal/http-server/handlers"
	"url-shortener/internal/logger"
	"url-shortener/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// URLGetter is an interface for getting url by alias
//
//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --name=URLGetter
type URLGetter interface {
	GetURL(alias string) (string, error)
}

func New(log *slog.Logger, urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.redirect.New"

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

		log.Info("request", slog.Any("alias", alias))

		url, err := urlGetter.GetURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)
			render.JSON(w, r, handlers.Error("not found"))

			return
		}
		if err != nil {
			log.Error("failed to get url", logger.Err(err))
			render.JSON(w, r, handlers.Error("internal error"))

			return
		}

		log.Info("got url", slog.String("url", url))

		http.Redirect(w, r, url, http.StatusFound)
	}
}
