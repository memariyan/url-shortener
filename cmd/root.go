package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
