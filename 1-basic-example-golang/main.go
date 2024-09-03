package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "HTTP Server",
		Usage: "Run a simple http server",
		Commands: []*cli.Command{
			{
				Name:  "run-http-server",
				Usage: "Start the http server",
				Action: func(c *cli.Context) error {
					r := gin.Default()
					r.GET("/", func(c *gin.Context) {
						c.JSON(200, gin.H{
							"message": "Hello World!!",
						})
					})
					r.Run(":8080")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
