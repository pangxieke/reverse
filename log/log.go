package log

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"reverse/config"
)

var (
	Log      *logrus.Logger
	ErrorLog *logrus.Logger
	logFile  = "./access.log"
)

func Init() (err error) {
	if config.Server.LogFile != "" {
		logFile = config.Server.LogFile
	}
	initLog()
	return
}

func Info(args ...interface{}) {
	Log.Info(args)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return Log.WithFields(fields)
}

func initLog() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	Log.SetOutput(file)
}

type ResponseWithRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (rec *ResponseWithRecorder) WriteHeader(statusCode int) {
	rec.ResponseWriter.WriteHeader(statusCode)
	rec.statusCode = statusCode
}

func (rec *ResponseWithRecorder) Write(d []byte) (n int, err error) {
	n, err = rec.ResponseWriter.Write(d)
	if err != nil {
		return
	}
	rec.body.Write(d)

	return
}
