package main

import (
	"fmt"
	"log"

	"github.com/DeanThompson/ginpprof"
	"github.com/rollbar/rollbar-go"
	conf "gitlab.com/digiverse/gosseractcv/config"
	"gitlab.com/digiverse/gosseractcv/logging"
	"gitlab.com/digiverse/gosseractcv/routers"
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
