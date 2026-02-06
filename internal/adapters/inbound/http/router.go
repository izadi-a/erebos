package http

import (
	"net/http"
	"strings"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	userHandler *UserHandler
}

func NewRouter(userHandler *UserHandler) http.Handler {
	return withMiddlewares(&Router{
		userHandler: userHandler,
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// Swagger UI
	if strings.HasPrefix(path, "/swagger/") {
		httpSwagger.WrapHandler(w, req)
		return
	}

	switch {
	// User
	case path == "/user" && method == http.MethodPost:
		r.userHandler.Create(w, req)
	case path == "/users" && method == http.MethodGet:
		r.userHandler.FindAll(w, req)
	case strings.HasPrefix(path, "/user/") && method == http.MethodGet:
		id := strings.TrimPrefix(path, "/user/")
		r.userHandler.FindByID(w, req, id)
	case strings.HasPrefix(path, "/user/") && method == http.MethodDelete:
		id := strings.TrimPrefix(path, "/user/")
		r.userHandler.Delete(w, req, id)

	default:
		http.NotFound(w, req)
	}
}
