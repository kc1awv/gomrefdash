package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Dashboard struct {
	Config    *Config
	Router    *gin.Engine
	Reflector *Reflector
}

func NewDashboard(config *Config) (*Dashboard, error) {
	// reflector setup
	reflector, err := NewReflectorFromFile(config.MrefFile)
	if err != nil {
		return nil, fmt.Errorf("unable to create reflector: %s", err)
	}

	d := Dashboard{Config: config, Reflector: reflector}
	// router
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	subPath := router.Group(config.SubPath)
	{
		subPath.StaticFile("/favicon.ico", "static/favicon.ico")
		subPath.Static("/static", "static")
		subPath.GET("/", d.showIndexPage)
		subPath.GET("/links", d.showLinksPage)
		subPath.GET("/peers", d.showPeersPage)
		subPath.GET("/status", d.showStatus) //dashboard status metadata
		subPath.GET("/json/reflector", d.showReflectorJSON)
		subPath.GET("/json/stations", d.showStationDataJSON)
		subPath.GET("/json/links", d.showLinksDataJSON)
		subPath.GET("/json/modulesinuse", d.showModulesInUseJSON)
		subPath.GET("/json/peers", d.showPeers)
	}
	d.Router = router

	return &d, nil
}
