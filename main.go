package main

import (
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"myweb/controllers"
	"myweb/models"
	"myweb/system"
	"os"
	"path/filepath"
	"strings"
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
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.NoRoute(controllers.Handle404)

}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Critical(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
