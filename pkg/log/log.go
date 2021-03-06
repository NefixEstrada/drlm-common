// SPDX-License-Identifier: AGPL-3.0-only

package log

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is the configuration of the log package
type Config struct {
	// Level is the level that the logger is going to log
	Level string

	// File is the file where the logger is going to write the logs
	File string
}

// Init initializes the log package
func Init(cfg Config) {
	level, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)

	pathMap := lfshook.PathMap{
		logrus.TraceLevel: cfg.File,
		logrus.DebugLevel: cfg.File,
		logrus.InfoLevel:  cfg.File,
		logrus.WarnLevel:  cfg.File,
		logrus.ErrorLevel: cfg.File,
		logrus.FatalLevel: cfg.File,
		logrus.PanicLevel: cfg.File,
	}

	logrus.AddHook(lfshook.NewHook(
		pathMap,
		&logrus.TextFormatter{},
	))
}

// SetDefaults sets the default configurations for Viper
func SetDefaults(app string) {
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.file", app+".log")
}
