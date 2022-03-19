package providers

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sepehrsoh/go-http-server/config"
	"sepehrsoh/go-http-server/src/modules/panel"
	"time"

	"os"
	"os/signal"
)

type HttpServer struct {
	router       *mux.Router
	panelConfigs config.PanelServerConfiguration
	authConfig   config.AuthenticationConfiguration
}

func NewHttpServer(handler panel.IPanelServer, panelConfigs config.PanelServerConfiguration,
	authConfig config.AuthenticationConfiguration) IServer {
	router := getRouter(handler)

	return &HttpServer{
		router:       router,
		panelConfigs: panelConfigs,
		authConfig:   authConfig,
	}
}

func (s *HttpServer) Run() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "gracefully waiting for live connections")
	flag.Parse()
	var logPath = "./panel_server.log"
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.panelConfigs.Host, s.panelConfigs.HTTPPort),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CombinedLoggingHandler(f, csrf.Protect([]byte(s.authConfig.TokenSecret), csrf.Secure(false))(s.router)),
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Info("Panel Server Start")
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait until the timeout deadline.
	go func() {
		if err = srv.Shutdown(ctx); err != nil {
			log.Error(err)
		}
	}()
	<-ctx.Done()

	log.Info("Shutting down")
}

func getRouter(panelServer panel.IPanelServer) *mux.Router {
	router := mux.NewRouter()

	router.Use(panelServer.Auth)

	router.NotFoundHandler = http.HandlerFunc(panelServer.NotFound)

	router.HandleFunc("/ping", panelServer.Ping).Methods(http.MethodGet)

	return router
}
