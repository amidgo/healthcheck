package httppinger

import (
	"net/http"

	"github.com/amidgo/healthcheck"
)

func Handler(h healthcheck.Handler, headerKey string) http.Handler {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		services := make([]string, 0)

		if headerKey != "" {
			services = r.Header.Values(headerKey)
		}

		err := h.Handle(r.Context(), services...)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return handler
}
