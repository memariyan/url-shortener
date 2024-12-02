package cmd

import (
	"log"
)

func Execute() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error on running cmd!")
	}
}
