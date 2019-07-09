package api

import (
	conf "github.com/carcinodehyde/OCR-Go/config"
	cons "github.com/carcinodehyde/OCR-Go/constant"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	startTime := conf.Configuration.Server.StartTime.In(loc)
	runMode := conf.Configuration.Server.RunMode

	res := map[string]string{
		"start_time": startTime.Format("[02 January 2006] 15:04:05 MST"),
		"message":    "Version " + cons.AppVersion + " run on " + runMode + " mode",
	}
	c.JSON(http.StatusOK, res)
}
