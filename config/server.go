package config

import (
	"time"
)

type ServerModel struct {
	StartTime    time.Time
	RunMode      string
	Port         string
	Url          string
	TempFilesDir string
	FeUrl        string
}
