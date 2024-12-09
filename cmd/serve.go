package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"url-shortner/internal/config"
	"url-shortner/internal/database"
	"url-shortner/internal/http"
	"url-shortner/internal/tracing"
	"url-shortner/internal/worker"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "url shortener server cli execution using cobra",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run",
	Long:  `Run the server on the defined port`,
	Run: func(cmd *cobra.Command, args []string) {
		startApplication()
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("port", "p", config.App().Server.Port, "the port of server")
}

func startApplication() {
	database.ConnectDB(&config.App().MySQL)
	database.ConnectRedis(&config.App().Redis)
	tracing.Start()
	worker.SaveWorker().Start()
	defer worker.SaveWorker().Stop()

	// run the server
	port, _ := rootCmd.Flags().GetInt("port")
	e := http.NewHttpServer()
	err := e.Start(":" + strconv.Itoa(port))
	e.Logger.Fatal(err)
}
