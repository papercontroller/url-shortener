package remove

import (
	"errors"
	"log/slog"
	"main/internal/lib/api/response"
	"main/internal/storage"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type URLDelete interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDelete URLDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.Delete.New"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, response.Error("invalid request"))

			return
		}

		err := urlDelete.DeleteURL(alias)
		if errors.Is(err, storage.ErrURLNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, response.Error("not found"))

			return
		}
		if err != nil {
			log.Error("failed to get url")

			render.JSON(w, r, response.Error("internal error"))

			return
		}

		log.Info("url deleted", slog.String("alias", alias))

		render.JSON(w, r, response.OK())
		w.WriteHeader(http.StatusOK)
	}
}
