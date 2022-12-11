package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Dashboard struct {
	Config    *Config
	Router    *gin.Engine
	Reflector *Reflector
	Version   string
}

func NewDashboard(config *Config, version string) (*Dashboard, error) {
	// reflector setup
	reflector, err := NewReflectorFromFile(config.MrefFile)
	if err != nil {
		return nil, fmt.Errorf("unable to create reflector: %s", err)
	}

	d := Dashboard{Version: version, Config: config, Reflector: reflector}
	// router
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	subPath := router.Group(config.SubPath)
	{
		subPath.GET("/status", d.showStatus) //dashboard status metadata
		subPath.GET("/json/reflector", d.showReflectorJSON)
		subPath.GET("/json/stations", d.showStationDataJSON)
		subPath.GET("/json/links", d.showLinksDataJSON)
		subPath.GET("/json/modulesinuse", d.showModulesInUseJSON)
		subPath.GET("/json/peers", d.showPeers)
		subPath.GET("/json/metadata", d.showMetadata)
		subPath.Static("/assets", "frontend/spa/assets")
		subPath.Static("/icons", "frontend/spa/icons")
		subPath.StaticFile("/index.html", "frontend/spa/index.html")
		subPath.StaticFile("/", "frontend/spa/index.html")
		subPath.StaticFile("/favicon.ico", "frontend/spa/favicon.ico")

	}
	d.Router = router

	return &d, nil
}
