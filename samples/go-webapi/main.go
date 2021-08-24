package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-webapi/controller"
	"go-webapi/middleware"
)

type config struct {
	Addr string `json:"addr"`
	Seed string `json:"seed"`
}

func newConfig() *config {
	return &config{
		Addr: "localhost:2322",
	}
}

func main() {
	conf := newConfig()
	readFlags(conf)

	engine := gin.Default()
	engine.Use(cors.Default())

	engine.Use(middleware.Mnemonic(conf.Seed))

	apiRouteGroup := engine.Group("/api/v1")

	controller.RegisterSignTransactionRoutes(apiRouteGroup.Group("/sign_transaction"))

	log.Fatal(engine.Run(conf.Addr))
}

func readFlags(conf *config) {
	c := flag.String("config", "go-webapi.config.json", "Configuration file")

	flag.Parse()

	if err := readJSONFile(*c, &conf); err != nil {
		panic(fmt.Sprintf("failed to read configuration: %s", err))
	}
}

func readJSONFile(name string, data interface{}) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(data)
}
