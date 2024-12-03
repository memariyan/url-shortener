package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"url-shortner/internal/config"
	"url-shortner/internal/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "db init",
	Long:  `Initialize the database according to the sql script`,
	Run: func(cmd *cobra.Command, args []string) {
		config.ReadConfig()
		database.ConnectDB(&config.Application.MySQL)
		sqlQuery, err := os.ReadFile("./scripts/init_db.sql")
		if err != nil {
			panic(err)
		}

		database.MySQL.Exec(string(sqlQuery))
		log.Infoln("database initialized successfully!")
	},
}
