package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"url-shortner/application/config"
	"url-shortner/application/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "db init",
	Long:  `Initialize the database according to the sql script`,
	Run: func(cmd *cobra.Command, args []string) {

		// connect to database
		database.Connect(&config.Application.MySQL)

		sqlQuery, err := os.ReadFile("./scripts/init_db.sql")
		if err != nil {
			panic(err)
		}
		database.DB.Exec(string(sqlQuery))
		log.Infoln("database initialized successfully!")
	},
}
