package middleware

import (
	"net/http"
)

type MiddleWare struct{}

func NewMiddleWare() *MiddleWare {
	return &MiddleWare{}
}

func (s *MiddleWare) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//	TODO complete this part base on your project
		next.ServeHTTP(writer, request)
	})
}
