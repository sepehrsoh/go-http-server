package cli

import (
	"github.com/spf13/cobra"
)

var (
	configFile string

	rootCmd = &cobra.Command{
		Use:   "http-server",
		Short: "simple http server using go",
		Long:  `simple http server using go`,
	}
)

func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config", "config file")
	rootCmd.PersistentFlags().StringP("author", "a", "sepehrsoh", "sepehrsoh.ir")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
}
