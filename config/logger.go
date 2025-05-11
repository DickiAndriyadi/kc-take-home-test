package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type LogError struct {
	Message  string                 `json:"message"`
	Request  map[string]interface{} `json:"request"`
	FilePath string                 `json:"file_path"`
}

func Log() *logrus.Entry {
	logger := logrus.New()
	log := logrus.NewEntry(logger)
	log.Logger.SetFormatter(&logrus.JSONFormatter{})
	log.Logger.Out = os.Stdout

	return log
}

func GetErrorFileLine() string {
	_, filePath, lineNumber, _ := runtime.Caller(1)
	absFilePath, _ := filepath.Abs(filePath)
	return fmt.Sprintf("%s:%v", absFilePath, lineNumber)
}

func PrintErrorLog(err error, filePath string, request map[string]interface{}) {
	Log().Error(ErrorFormatter(LogError{
		Message:  err.Error(),
		Request:  request,
		FilePath: filePath,
	}))
}

func ErrorFormatter(logError LogError) error {
	errorByte, _ := json.Marshal(logError)
	return fmt.Errorf("%s", string(errorByte))
}
