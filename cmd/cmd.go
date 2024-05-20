package cmd

import (
	"account-app-gin/internal/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "app",
	}
	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()

			r.GET("/", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "Hello, World!",
				})
			})

			r.Run()
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

	database.ConnectDB()

	rootCmd.AddCommand(srvCmd, dbCmd)
	dbCmd.AddCommand(mgrCmd)

	rootCmd.Execute()
}
