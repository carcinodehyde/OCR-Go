package helpers

import (
	cons "gitlab.com/digiverse/gosseractcv/constant"
	"gitlab.com/digiverse/gosseractcv/logging"

	"os"
)

var log = logging.MustGetLogger(cons.LOG_MODULE)

func RemoveFile(filePath string) {

	err := os.Remove(filePath)
	if err != nil {
		log.Errorf(logging.INTERNAL, "can't remove temp file: %s", err.Error())
	}
}
