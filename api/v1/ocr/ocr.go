package ocr

import (
	"net/http"
	"strings"

	conf "gitlab.com/digiverse/gosseractcv/config"
	cons "gitlab.com/digiverse/gosseractcv/constant"
	"gitlab.com/digiverse/gosseractcv/helpers"
	"gitlab.com/digiverse/gosseractcv/logging"
	"github.com/chilts/sid"
	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"
	"gocv.io/x/gocv"
)

var log = logging.MustGetLogger(cons.LOG_MODULE)

func Parse(c *gin.Context) {
	loggingID := logging.INTERNAL

	file, _ := c.FormFile("file")
	log.Infof(loggingID, "uploaded file with name: %s", file.Filename)

	dst := conf.Configuration.Server.TempFilesDir + sid.Id()
	if err := c.SaveUploadedFile(file, dst+".jpg"); err != nil {
		log.Errorf(loggingID, "error save: %s", err.Error())
		c.JSON(http.StatusInternalServerError, cons.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{"Something went wrong."}, nil))
		return
	}
	defer helpers.RemoveFile(dst + ".jpg")

	tmpGray := gocv.IMRead(dst+".jpg", gocv.IMReadGrayScale)
	if tmpGray.Empty() {
		log.Errorf(loggingID, "Opencv failed to read image: %s\n", dst+".jpg")
		c.JSON(http.StatusInternalServerError, cons.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{"Something went wrong."}, nil))
		return
	}

	tmpBin := gocv.NewMat()
	gocv.Threshold(tmpGray, &tmpBin, 101, 255, gocv.ThresholdBinary)

	if ok := gocv.IMWrite(dst+"_gray"+".jpg", tmpBin); !ok {
		log.Errorf(loggingID, "Opencv failed to write grayscale image: %s\n", dst+"_gray"+".jpg")
		c.JSON(http.StatusInternalServerError, cons.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{"Something went wrong."}, nil))
		return
	}

	client := gosseract.NewClient()
	client.SetLanguage("OCR")
	defer client.Close()
	client.SetImage(dst + "_gray" + ".jpg")
	text, _ := client.Text()

	splited := strings.Split(text, "\n")
	log.Infof(loggingID, "splitted text: %s", splited)

	c.JSON(http.StatusOK, cons.NewGenericResponse(http.StatusOK, cons.OK, []string{"Submit success"}, splited))
}
