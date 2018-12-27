package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/WALL-E/http2mq/app"
)

func main() {
	conf := flag.String("conf", "http2mq.yaml", "http2mq configuration file")
	flag.Parse()

	application, err := app.NewApp(*conf)
	if err != nil {
		fmt.Printf("create app error %s", err.Error())
		os.Exit(2)
	}
	defer application.Close()

	application.Run()
}
