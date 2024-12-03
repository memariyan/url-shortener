package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"url-shortner/internal/config"
	"url-shortner/internal/database"
	"url-shortner/internal/http"
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
		config.ReadConfig()
		database.ConnectDB(&config.Application.MySQL)
		database.ConnectRedis(&config.Application.Redis)

		// run the server
		port, _ := rootCmd.Flags().GetInt("port")
		e := http.NewServer()
		err := e.Start(":" + strconv.Itoa(port))
		e.Logger.Fatal(err)
	},
}

func init() {
	rootCmd.PersistentFlags().IntP("port", "p", config.Application.Server.Port, "the port of server")
}
