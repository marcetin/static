package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
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
		sub.Use(middleware.Logger())
		sub.Use(middleware.Recover())
		sub.Static("/", path+n)
		hosts[n] = &Host{sub}
	}
	// Server
	app.Use(func(c *fiber.Ctx) {
		host := hosts[c.Hostname()]
		if host == nil {
			c.SendStatus(fiber.StatusNotFound)
		} else {
			host.Fiber.Handler()(c.Fasthttp)
		}
	})
	log.Fatal(app.Listen(":" + os.Getenv("JORMSTATICPORT")))
}
