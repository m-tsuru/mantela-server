package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tkytel/mantela-adder/lib"
)

func main() {
	// Load the configuration file
	config, err := lib.GetConfig("./config.toml")
	if err != nil {
		panic("Error loading configuration: " + err.Error())
	}

	// Starting up the Web Server
	app := fiber.New()
	if config.Serve.Static {
		app.Static("/", config.Serve.StaticDir)
	}
	app.Get("/.well-known/mantela.json", func(c *fiber.Ctx) error {
		return lib.HandleMantela(c, config.Mantela.Source, config.Mantela.Diff)
	})

	app.Listen(":" + strconv.Itoa(config.Serve.Port))
}
