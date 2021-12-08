package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type (
	Host struct {
		Fiber *fiber.App
	}
)

func main() {
	app := fiber.New()
	path := os.Getenv("JORMSTATICPATH")
	// Hosts
	hosts := map[string]*Host{}
	sites, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range sites {
		n := f.Name()
		fmt.Println("domain:", n)
		sub := fiber.New()
		// sub.Use(middleware.Logger())
		// sub.Use(middleware.Recover())

		sub.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}))

		sub.Use(compress.New(compress.Config{
			Level: compress.LevelBestSpeed, // 1
		}))

		sub.Use(cache.New(cache.Config{
			Next: func(c *fiber.Ctx) bool {
				return c.Query("refresh") == "true"
			},
			Expiration:   30 * time.Minute,
			CacheControl: true,
		}))

		sub.Static("/", path+n)
		hosts[n] = &Host{sub}
	}
	// Server
	app.Use(func(c *fiber.Ctx) error {
		host := hosts[c.Hostname()]
		if host == nil {
			c.SendStatus(fiber.StatusNotFound)
		} else {
			host.Fiber.Handler()
		}
		return nil
	})
	log.Fatal(app.Listen(":" + os.Getenv("JORMSTATICPORT")))
}
