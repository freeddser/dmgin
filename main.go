package main

import (
	"flag"
	"fmt"
	"github.com/freeddser/dmgin/repository"
	"github.com/freeddser/dmgin/router"
	"github.com/freeddser/rs-common/config"
	"github.com/freeddser/rs-common/util"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// init config
	configFile := flag.String("c", "", "Configuration File")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options\n")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := config.Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}

	// 禁用控制台颜色
	gin.DisableConsoleColor()

	// debug, release
	mode := config.MustGetString("server.mode")
	gin.SetMode(mode)

	// 创建记录日志的文件
	f, _ := os.Create(config.MustGetString("server.log_path"))
	gin.DefaultWriter = io.MultiWriter(f)

	// init Redis Cache
	InitRedisFactory()
	// init Mysql database
	//mysql
	if config.MustGetString("switch.mysql") == "on" {
		err := repository.InitMysqlDbFactory()
		if err != nil {
			log.Panic("Cannot connect to Mysqldatabase: ", err.Error())
			return
		}
	}

	//http server
	log.Println("HTTP Server Started in :" + strconv.Itoa(config.MustGetInt("Server.port")))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.MustGetInt("Server.port")), router.NewRouter()))

}

func InitRedisFactory() {
	util.InitRedis(config.MustGetString("cache.redis_main"))
}
