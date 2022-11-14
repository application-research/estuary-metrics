package main

import (
	"context"
	"fmt"
	"github.com/application-research/estuary-metrics/core"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/rest/app"
	_ "github.com/application-research/estuary-metrics/rest/docs"
	"github.com/droundy/goopt"
	"github.com/spf13/viper"
	"github.com/whyrusleeping/memo"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string

	// OsSignal signal used to shutdown
	OsSignal chan os.Signal
)

// @title Estuary Metrics API
// @version 0.0.1
// @description Estuary Metrics API
// @termsOfService

// @contact.name Outercore Engineering
// @contact.url http://me.com/terms.html
// @contact.email me@me.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host https://metrics-api.estuary.tech/api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @security[0].BearerAuth
// @BasePath /
func main() {
	OsSignal = make(chan os.Signal, 1)

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	dbHost, okHost := viper.Get("DB_HOST").(string)
	dbUser, okUser := viper.Get("DB_USER").(string)
	dbPass, okPass := viper.Get("DB_PASS").(string)
	dbName, okName := viper.Get("DB_NAME").(string)
	dbPort, okPort := viper.Get("DB_PORT").(string)
	if !okHost || !okUser || !okPass || !okName || !okPort {
		log.Fatalf("Error while reading database config")
	}

	// Define version information
	goopt.Version = fmt.Sprintf(
		`Application build information
				  Build date      : %s
				  Build number    : %s
				  Git commit      : %s
				  Runtime version : %s
				  Built on OS     : %s
				`, BuildDate, BuildNumber, LatestCommit, RuntimeVer, BuiltOnOs)
	goopt.Parse(nil)

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	dao.DB = db                   // database connection
	dao.Cacher = memo.NewCacher() // cache instance

	//	initialize the core
	dao.Metrics, _ = core.Init(dao.DB, dao.Cacher)

	dao.Logger = func(ctx context.Context, sql string) {
		fmt.Printf("SQL: %s\n", sql)
	}
	go app.GinServer()     // rest
	go app.InitHeartbeat() // heartbeat

	LoopForever()
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Printf("Entering infinite loop\n")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Printf("Exiting infinite loop received OsSignal\n")
}
