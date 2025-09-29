package user

import (
	"net/http"

	middleware "github.com/enghasib/laundry_service/rest/middlewares"
)

func (h *UserHandler) UserRoutes(mux *http.ServeMux, mngr *middleware.MiddlewareManager) *http.ServeMux {
	mux.Handle("/users/register", http.HandlerFunc(h.CreateUser))
	return mux
}
