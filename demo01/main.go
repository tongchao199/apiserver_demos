package main

import (
	"errors"
	"flag"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tongchao199/apiserver_demos/demo01/router"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Opts struct {
	listen  string
	logFile string
	logSize int
}

func main() {
	opts := &Opts{}

	flag.StringVar(&opts.listen, "listen", ":30000", "server listen port")
	flag.StringVar(&opts.logFile, "logFile", "server.log", "log file path")
	flag.IntVar(&opts.logSize, "logSize", 512, "log file size, the unit is MB(MegaByte)")

	flag.Parse()

	if err := Serve(opts); err != nil {
		log.Error(err)
	}
}

func Serve(opts *Opts) error {
	//Init logger
	if opts.logFile != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   opts.logFile,
			MaxSize:    opts.logSize,
			MaxBackups: 10,
			MaxAge:     30,   // days
			Compress:   true, // enable compress
			LocalTime:  true, //
		})
	}

	log.Infof("opts: %v", opts)

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
	go func(opts *Opts) {
		if err := pingServer(opts); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Infof("The router has been deployed successfully.")
	}(opts)

	log.Infof("Start to listening the incoming requests on http address: %s", ":8080")
	log.Infof(http.ListenAndServe(opts.listen, g).Error())

	return nil
}

// pingServer pings the http server to make sure the router is working.
func pingServer(opts *Opts) error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1" + opts.listen + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
