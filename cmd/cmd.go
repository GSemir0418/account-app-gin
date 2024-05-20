package cmd

import (
	"account-app-gin/internal/database"
	"log"
	"os/exec"

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
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
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
	dbCmd.AddCommand(mgrCmd)

	rootCmd.Execute()
}
