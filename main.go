package main

import (
	"fengshen-svr/biz"
	"fengshen-svr/cfg"
	"fengshen-svr/common/constants"
	"fengshen-svr/global"
	"fengshen-svr/middlewares"
	"fengshen-svr/routers"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initEnv()
	initConfig()
	initLogger()
	initDB()
	initInstances()
	initGinServer()
}

func initGinServer() {
	log.Infof("starting server...")
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Use(middlewares.HandleError)
	routers.InitGinGroups(r)
	r.Run(":" + strconv.Itoa(cfg.Config.Server.Port))
}

func initEnv() {
	env := strings.TrimSpace(os.Getenv(constants.ENV))
	env = strings.ToLower(env)
	fmt.Printf("current env is [%s]\n", env)
	if len(env) == 0 {
		env = constants.DEV
		fmt.Printf("now env will be set as default value [%s]\n", env)
	}
	global.InitEnv(env)
}

func initConfig() {
	env := global.GetEnv()
	filename := fmt.Sprintf("configs/%s.toml", env)
	fmt.Printf("use config file %s \n", filename)
	err := configor.Load(&cfg.Config, filename)
	if err != nil {
		fmt.Printf("init config failed err=%v \n", err)
		panic(err)
	}
}

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	level, err := log.ParseLevel(cfg.Config.Log.Level)
	if err != nil {
		fmt.Printf("init logger failed err=%v \n", err)
		panic(err)
	}
	log.SetLevel(level)
}

func initDB() {
	// source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	source := cfg.Config.DB.Url
	if strings.Contains(source, "%s") {
		user := os.Getenv(constants.MYSQL_USERNAME)
		pwd := os.Getenv(constants.MYSQL_PASSWORD)
		addr := os.Getenv(constants.MYSQL_ADDRESS)
		dataBase := os.Getenv(constants.MYSQL_DATABASE)
		if dataBase == "" {
			dataBase = "golang_demo"
		}
		source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	}

	fmt.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		fmt.Printf("gorm open db err=%v \n", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("db.DB() err=%v \n", err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(cfg.Config.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Config.DB.MaxOpenConns)
	global.InitDB(db)
}

func initInstances() {
	err := biz.InitAllBiz()
	if err != nil {
		log.Errorf("InitAllBiz err=%v", err)
		panic(err)
	}
}
