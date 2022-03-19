package panel

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Get struct {
}

func NewPanelGet() *Get {
	return &Get{}
}

func (g *Get) Ping(w http.ResponseWriter, r *http.Request) {
	logrus.Info("Receive message to Ping")
	response(w, r, "ping.html", &PageResponse{
		Title: "ping",
		Parameters: struct {
			Pong string
		}{Pong: "Pong"},
		Errors: nil,
	})
}
