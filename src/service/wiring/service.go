package wiring

import (
	"sepehrsoh/go-http-server/src/middleware"
	"sepehrsoh/go-http-server/src/modules/panel"
	"sepehrsoh/go-http-server/src/service/providers"
)

func (w *Wire) GetPanelServer() providers.IServer {
	return providers.NewHttpServer(
		w.GetPanelServerHandler(),
		w.Configurations.PanelServerConfiguration,
		w.Configurations.AuthenticationConfiguration,
	)
}

func (w *Wire) GetPanelServerHandler() *panel.Handler {
	return panel.NewHandler(
		middleware.NewMiddleWare(),
		panel.NewPanelGet(),
		panel.NewPanelPost(),
		panel.NewPanelDelete(),
		panel.NewPanelPut(),
		middleware.NewMiddleWare())
}
