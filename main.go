package main

import (
	"fmt"
	"log"

	"github.com/DeanThompson/ginpprof"
	conf "github.com/carcinodehyde/OCR-Go/config"
	"github.com/carcinodehyde/OCR-Go/logging"
	"github.com/carcinodehyde/OCR-Go/routers"
	"github.com/rollbar/rollbar-go"
)

func main() {
	app := routers.GetRouter()
	ginpprof.Wrap(app)
	err := app.Run(":" + conf.Configuration.Server.Port)
	if err != nil {
		log.Panicf(logging.INTERNAL, "can't start the app: %s", err.Error())
		rollbar.Error(logging.INTERNAL + fmt.Sprintf("can't start the app: %s", err.Error()))
	}
}
