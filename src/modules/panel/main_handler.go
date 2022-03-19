package panel

import (
	"net/http"
	"sepehrsoh/go-http-server/src/middleware"
)

type Handler struct {
	authRepository middleware.IMiddleware
	IPanelGet
	IPanelPost
	IPanelDelete
	IPanelPut
	middleware.IMiddleware
}

func NewHandler(authRepository middleware.IMiddleware,
	getRepository IPanelGet, postRepository IPanelPost,
	deleteRepository IPanelDelete, putRepository IPanelPut,
	middleware middleware.IMiddleware) *Handler {
	return &Handler{
		authRepository: authRepository,
		IPanelGet:      getRepository,
		IPanelPost:     postRepository,
		IPanelDelete:   deleteRepository,
		IPanelPut:      putRepository,
		IMiddleware:    middleware,
	}
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	response(w, r, "404.html", &PageResponse{
		Title:  "Not Found!",
		Errors: nil,
	})
}
