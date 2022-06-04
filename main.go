package main

import (
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"myweb/models"
	"myweb/system"
)

func main() {
	configFilePath := flag.String("C", "conf/conf.yaml", "change the config file path")
	logConfigPath := flag.String("L", "./conf/seelog.xml", "change the log config file path")
	flag.Parse()

	logger, err := log.LoggerFromConfigAsFile(*logConfigPath) //配置xml文件目录
	if err != nil {
		_ = fmt.Errorf("parse seelog.xml error : %s\n", err)
	}
	log.ReplaceLogger(logger)
	defer log.Flush()
	logger.Info("seelog test begin")

	if err := system.LoadConfiguration(*configFilePath); err != nil {
		log.Critical("err parsing config log file", err)
		return
	}

	db, err := models.InitDB()
	defer db.Close()
	logger.Debug("seelog test end")

	/*	route:= gin.Default()
		route.StaticFS()*/

}
