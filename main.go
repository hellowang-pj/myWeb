package main

import (
	"flag"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"myweb/controllers"
	"myweb/models"
	"myweb/system"
	"net/http"
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
	fmt.Println(filepath.Join(getCurrentDirectory()))
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	//router.Static("/static", filepath.Join(getCurrentDirectory(),"./static"))
	router.StaticFS("/static", http.Dir("./static"))
	router.StaticFile("/favicon.ico", filepath.Join(getCurrentDirectory(), "./views/favicon.ico"))
	router.NoRoute(controllers.Handle404)
	router.GET("/ahaha", func(c *gin.Context) {
		c.String(http.StatusOK, "Who are you?")
	})

	router.Run(":8090")
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Critical(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
