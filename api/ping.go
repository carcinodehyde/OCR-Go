package api

import (
	conf "gitlab.com/digiverse/gosseractcv/config"
	cons "gitlab.com/digiverse/gosseractcv/constant"

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
		"message":    "Version " + cons.APP_VERSION + " run on " + runMode + " mode",
	}
	c.JSON(http.StatusOK, res)
}