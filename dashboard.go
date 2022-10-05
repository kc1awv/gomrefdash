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
	router.StaticFile("/favicon.ico", "static/favicon.ico")
	router.Static("/static", "static")
	router.GET("/", d.showIndexPage)
	router.GET("/links/", d.showLinksPage)
	router.GET("/peers/", d.showPeersPage)
	router.GET("/status", d.showStatus) //dashboard status metadata
	router.GET("/json/reflector", d.showReflectorJSON)
	router.GET("/json/stations", d.showStationDataJSON)
	router.GET("/json/links", d.showLinksDataJSON)
	router.GET("/json/modulesinuse", d.showModulesInUseJSON)
	router.GET("/json/peers", d.showPeers)
	d.Router = router

	return &d, nil
}
