package cmd

import (
	"account-app-gin/internal/database"
	"account-app-gin/router"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "app",
	}
	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			r := router.New()
			r.Run(":8080")
		},
	}
	dbCmd := &cobra.Command{
		Use: "db",
	}
	mgrCmd := &cobra.Command{
		Use: "migrate:create",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}
	clearDBCmd := &cobra.Command{
		Use: "clear",
		Run: func(cmd *cobra.Command, args []string) {
			database.TruncateTables(nil, []string{"users", "items", "tags", "validation_codes"})
		},
	}
	testCmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			if err := exec.Command(
				"go", "test", "./...",
			).Run(); err != nil {
				log.Fatalln(err)
			}
		},
	}

	database.ConnectDB()

	rootCmd.AddCommand(srvCmd, dbCmd, testCmd)
	dbCmd.AddCommand(mgrCmd, clearDBCmd)

	rootCmd.Execute()
}
