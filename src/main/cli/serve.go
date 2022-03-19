package cli

import (
	"github.com/spf13/cobra"
	"log"
	"sepehrsoh/go-http-server/config"
	"sepehrsoh/go-http-server/src/service/wiring"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving http-server, simple go http-server.",
	Long:  `Serving http-server, simple go http-server.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	conf, err := config.GetConfigFromYaml()

	if err != nil {
		panic(err)
	}

	if conf.AuthenticationConfiguration.TokenSecret == "" {
		log.Fatal("There is no token secret in config file\n")
	}

	wiring.HttpWire = wiring.NewWiring(conf)
	panelServer := wiring.HttpWire.GetPanelServer()
	panelServer.Run()
}
