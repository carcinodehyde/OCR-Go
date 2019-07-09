package ocr

import (
	"net/http"
	"strconv"
	"strings"

	"os/exec"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/chilts/sid"
	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"
	conf "github.com/carcinodehyde/OCR-Go/config"
	cons "github.com/carcinodehyde/OCR-Go/constant"
	"github.com/carcinodehyde/OCR-Go/helpers"
	"github.com/carcinodehyde/OCR-Go/logging"
)

var log = logging.MustGetLogger(cons.LOG_MODULE)

/*Parse uploaded image
Return: json formatted parsed text of the image
*/
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

	client := gosseract.NewClient()
	enhanceType := c.Query("enhancement")

	if len(enhanceType) > 0 {
		img, err := imgio.Open(dst + ".jpg")
		if err != nil {
			log.Errorf(loggingID, "error open: %s", err.Error())
			c.JSON(http.StatusInternalServerError, cons.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{"Something went wrong."}, nil))
			return
		}

		switch enhanceType {
		case cons.KTP:
			err = helpers.EnhanceKTP(dst, img)
			client.SetLanguage("OCR+IND")
			break
		case cons.GELAP:
			err = helpers.EnhanceGelap(dst, img)
			client.SetLanguage("IND")
			break
		case cons.TERANG:
			err = helpers.EnhanceTerang(dst, img)
			client.SetLanguage("IND")
			break
		case cons.AUTO:
			ret, err := exec.Command(`identify`, `-format`, `"%[mean]"`, dst+".jpg").Output()
			if err != nil {
				log.Errorf(loggingID, "Error identify image: %s", err.Error())
			}
			strB := string(ret)
			strB = strings.Trim(strB, "\"")
			strB = strings.Split(strB, ".")[0]

			intB, err := strconv.Atoi(strB)

			if intB < 55000 {
				err = helpers.EnhanceKTP(dst, img)
				client.SetLanguage("OCR+IND")
			} else if intB < 55800 {
				err = helpers.EnhanceTerang(dst, img)
				client.SetLanguage("IND")
			} else if intB < 65000 {
				client.SetLanguage("IND")
				client.SetImage(dst + ".jpg")
			} else if intB > 65000 {
				err = helpers.EnhanceGelap(dst, img)
				client.SetLanguage("IND")
			}
			client.SetLanguage("IND")
		}

		if err != nil {
			log.Errorf(loggingID, "error enhancement: %s", err.Error())
			c.JSON(http.StatusInternalServerError, cons.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{"Something went wrong."}, nil))
			return
		}

		defer helpers.RemoveFile(dst + "_enhanced.jpg")
		client.SetImage(dst + "_enhanced.jpg")
	} else {
		client.SetLanguage("IND")
		client.SetImage(dst + ".jpg")
	}

	text, err := client.Text()

	if err != nil {
		log.Errorf(loggingID, "error tesseract reading: %s", err.Error())
		c.JSON(http.StatusInternalServerError, cons.NewGenericResponse(http.StatusInternalServerError, cons.ERR, []string{"Something went wrong."}, nil))
		return
	}

	splited := strings.Split(text, "\n")
	defer client.Close()

	c.JSON(http.StatusOK, cons.NewGenericResponse(http.StatusOK, cons.OK, []string{"Submit success"}, splited))
}
