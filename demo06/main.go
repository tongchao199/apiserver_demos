package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/tongchao199/apiserver_demos/demo06/config"
	"github.com/tongchao199/apiserver_demos/demo06/model"
	"github.com/tongchao199/apiserver_demos/demo06/router"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Opts struct {
	cfg string
}

func main() {
	opts := &Opts{}

	pflag.StringVarP(&opts.cfg, "config", "c", "", "apiserver config file path.")
	pflag.Parse()

	if err := Serve(opts); err != nil {
		log.Error("failed to start server", err)
	}
}

func Serve(opts *Opts) error {
	// init config and logger
	if err := config.Init(opts.cfg); err != nil {
		panic(err)
	}

	// Init db
	db := model.NewDatabase()
	defer db.Close()

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())

	return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
