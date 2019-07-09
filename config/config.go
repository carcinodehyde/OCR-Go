package config

import (
	cons "github.com/carcinodehyde/OCR-Go/constant"
	"github.com/carcinodehyde/OCR-Go/logging"

	"os"
	"strconv"
	"time"
)

type ConfigurationModel struct {
	Server       ServerModel
	RollbarToken string
}

var (
	log           = logging.MustGetLogger(cons.LOG_MODULE)
	Configuration = &ConfigurationModel{}
)

func init() {
	internal := logging.INTERNAL
	prefixEnv := cons.PREFIX_ENV

	server := ServerModel{
		StartTime:    time.Now(),
		RunMode:      os.Getenv(prefixEnv + "RUN_MODE"),
		Port:         os.Getenv(prefixEnv + "PORT"),
		Url:          os.Getenv(prefixEnv + "SERVER_URL"),
		TempFilesDir: os.Getenv(prefixEnv + "TEMP_FILES_DIR"),
		FeUrl:        os.Getenv(prefixEnv + "FE_URL"),
	}

	rollbarToken := os.Getenv(prefixEnv + "ROLLBAR_TOKEN")

	Configuration = &ConfigurationModel{
		Server:       server,
		RollbarToken: rollbarToken,
	}

	if len(Configuration.Server.Port) < 1 {
		log.Panicf(internal, "Failed get port number")
	}

	if _, err := strconv.Atoi(Configuration.Server.Port); err != nil {
		log.Panicf(internal, "Seems port number isn't integer")
	}
}
