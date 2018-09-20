package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-macaron/macaron"
	logrus "github.com/sirupsen/logrus"
)

var LogTimeFormat = "2006-01-02 15:04:05"
var ColorLog = false

func InitLogger() {
	logPath := "logs/app.log"
	logfile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		logrus.Fatalln("open logfile failed")
	}
	//defer logfile.Close()
	logrus.SetOutput(logfile)
	logrus.SetLevel(logrus.DebugLevel)
}

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() macaron.Handler {
	return func(ctx *macaron.Context) {
		start := time.Now()

		logPath := "logs/exectime.log"
		logfile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln("open logfile failed")
		}
		logger := logrus.New()
		logger.SetOutput(logfile)
		logger.Printf("%s: Started %s %s for %s", time.Now().Format(LogTimeFormat), ctx.Req.Method, ctx.Req.RequestURI, ctx.RemoteAddr())
		rw := ctx.Resp.(macaron.ResponseWriter)
		ctx.Next()

		content := fmt.Sprintf("%s: Completed %s %v %s in %v", time.Now().Format(LogTimeFormat), ctx.Req.RequestURI, rw.Status(), http.StatusText(rw.Status()), time.Since(start))
		//if ColorLog {
		//	switch rw.Status() {
		//	case 200, 201, 202:
		//		content = fmt.Sprintf("\033[1;32m%s\033[0m", content)
		//	case 301, 302:
		//		content = fmt.Sprintf("\033[1;37m%s\033[0m", content)
		//	case 304:
		//		content = fmt.Sprintf("\033[1;33m%s\033[0m", content)
		//	case 401, 403:
		//		content = fmt.Sprintf("\033[4;31m%s\033[0m", content)
		//	case 404:
		//		content = fmt.Sprintf("\033[1;31m%s\033[0m", content)
		//	case 500:
		//		content = fmt.Sprintf("\033[1;36m%s\033[0m", content)
		//	}
		//}
		logger.Println(content)
	}
}

func SetExectimeLog() macaron.Handler {
	return func(c *macaron.Context) {
		logPath := "logs/exectime.log"
		logfile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln("open logfile failed")
		}
		c.Map(logfile)
		logger := log.New(logfile, "[DEBUG]", log.LstdFlags|log.Llongfile)
		c.Map(logger)

	}
}
