package panel

import (
	"net/http"
	"sepehrsoh/go-http-server/src/middleware"
)

type IPanelServer interface {
	NotFound(w http.ResponseWriter, r *http.Request)
	IPanelGet
	IPanelPut
	IPanelPost
	IPanelDelete
	middleware.IMiddleware
}

type IPanelGet interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type IPanelDelete interface {
}

type IPanelPost interface {
}

type IPanelPut interface {
}
