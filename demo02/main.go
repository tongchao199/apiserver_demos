package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/tongchao199/apiserver_demos/demo02/config"
	"github.com/tongchao199/apiserver_demos/demo02/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Opts struct {
	cfg     string
	logFile string
	logSize int32
}

func main() {
	opts := &Opts{}

	pflag.StringVarP(&opts.cfg, "config", "c", "", "apiserver config file path.")
	pflag.StringVarP(&opts.logFile, "logFile", "f", "server.log", "apiserver log file.")
	pflag.Int32VarP(&opts.logSize, "logSize", "s", 500, "apiserver log size.")

	pflag.Parse()

	if err := Serve(opts); err != nil {
		log.Error(err)
	}
}

func Serve(opts *Opts) error {
	//Init logger
	if opts.logFile != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   opts.logFile,
			MaxSize:    (int)(opts.logSize),
			MaxBackups: 10,
			MaxAge:     30,   // days
			Compress:   true, // enable compress
			LocalTime:  true, //
		})
	}

	log.Infof("opts: %v", opts)

	// init config
	if err := config.Init(opts.cfg); err != nil {
		panic(err)
	}

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
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
