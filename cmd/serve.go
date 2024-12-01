package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"url-shortner/application/config"
	"url-shortner/application/database"
	"url-shortner/application/http"
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
		// connect to database
		database.Connect(&config.Application.MySQL)

		// run the server
		port, _ := rootCmd.Flags().GetInt("port")
		e := http.New()
		err := e.Start(":" + strconv.Itoa(port))
		e.Logger.Fatal(err)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().IntP("port", "p", config.Application.Server.Port, "the port of server")
}
