package logger

import (
	"fmt"
	"io"
	"os"

	absPathFinder "github.com/logger/controller/abspath"
	log "github.com/sirupsen/logrus"
)

func InitializeLogging() {

	
	curPath := absPathFinder.GetCurrentAbPath()

	
	logFile := curPath + "/test2.log"
	log.Info(logFile)
	var file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
	   fmt.Println("Could Not Open Log File : " + err.Error())
	}

	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)
	
 
	//log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	log.SetLevel(log.DebugLevel)
 }