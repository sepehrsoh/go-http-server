package middleware

import "net/http"

type IMiddleware interface {
	Auth(next http.Handler) http.Handler
}
