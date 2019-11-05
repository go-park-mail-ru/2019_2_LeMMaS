package logger

import (
	"github.com/labstack/echo/middleware"
	"log"
	"os"
)

func SetLoggerConfig() middleware.LoggerConfig {
	f, err := os.OpenFile("agario.log",
		os.O_RDWR | os.O_CREATE | os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()

	return middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339}", "id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","file":"${long_file}"` + "\n",
		CustomTimeFormat: "2000-01-01 15:01:02",
		Output: f,
	}
}
