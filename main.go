package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/therudite/api/common"
	"github.com/therudite/api/config"
	_ "github.com/therudite/api/docs"
	"github.com/therudite/api/errors"
	"github.com/therudite/api/resources"
	"github.com/therudite/api/routes"
)

type serverConfig struct {
	Port string
}

func interrupt(errc chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errc <- fmt.Errorf("%s", <-c)
}

// @APIVersion 1.0.0
// @APITitle Swagger API
// @APIDescription Swagger API
// @Contact niteshagarwal1.618@gmail.com
// @TermsOfServiceUrl http://agarn.in
// @License E_CORP
// @LicenseUrl http://agarn.in
func main() {

	var err error

	// instantiating config manager
	var configManager config.ConfigManager
	configManager, err = config.NewConfigManager()
	if err != nil {
		panic(err)
		panic("Error instantiating config manager")
	}

	// instantiating logger
	var loggerConfig = new(common.LoggerConfig)
	configManager.Load("logger", loggerConfig)
	common.InitializeLogger(loggerConfig)
	common.LogString("Logger loaded, starting server ...")

	// instantiating resource manager
	var resourcemanager resources.ResourceManagerInterface
	resourcemanager, err = resources.NewResourceManager(configManager)
	if err != nil {
		panic(err)
		panic(errors.ResourceInitializationError.Error())
	}
	defer resourcemanager.Close()

	// Server Config
	serverconfig := new(serverConfig)
	configManager.Load("server", serverconfig)

	// instantiating router
	engine := gin.Default()

	errc := make(chan error)
	// Listen interrupts
	go interrupt(errc)

	// instantiating server
	go func() {
		fmt.Print("Creating route")

		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		routes.CreateFeedbackRoutes(engine, resourcemanager)
		server := http.Server{
			Addr:    ":" + serverconfig.Port,
			Handler: engine,
		}
		server.SetKeepAlivesEnabled(false)
		errc <- server.ListenAndServe()
	}()
	common.LogJSON(map[string]string{"message": "exit"}, <-errc)
}
