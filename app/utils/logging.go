package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogStructur struct {
	Message string
	Types   string
	Folder  string
}

func Log(msg string, types string, folder string) {
	logger := logrus.New()
	folders := folder
	os.MkdirAll("./logs/"+folders, 0755)
	filename := "./logs/" + folders + "/" + time.Now().Format("02-01-2006") + ".log"
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0755)
	defer file.Close()
	logger.SetOutput(file)

	switch types {
	case "err":
		logger.Error(msg)
	case "info":
		logger.Info(msg)
	case "debug":
		logger.Debug(msg)
	}

}