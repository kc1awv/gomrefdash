package dashboard

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kc1awv/gomrefdash/backend/pkg/dashboard/config"
)

type Dashboard struct {
	Config           *config.Config
	Router           *gin.Engine
	Reflector        *Reflector
	Version          string
	CountryCallsigns *CountryCallsigns
}

func NewDashboard(config *config.Config, version string) (*Dashboard, error) {
	// reflector setup
	reflector, err := NewReflectorFromFile(config.MrefFile, config.MrefPidFile)
	if err != nil {
		return nil, fmt.Errorf("unable to create reflector: %s", err)
	}

	// country callsigns setup
	log.Println("loading country callsigns from file")
	countryCallsigns, err := NewCountryCallsignsFromFile(config.CallsignCountryFile)
	if err != nil {
		return nil, fmt.Errorf("unable to create country callsigns: %s", err)
	}
	log.Println(countryCallsigns)

	d := Dashboard{Version: version, Config: config, Reflector: reflector, CountryCallsigns: countryCallsigns}
	// router
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	subPath := router.Group(config.SubPath)
	{
		subPath.GET("/status", d.showStatus) //dashboard status metadata
		subPath.GET("/json/reflector", d.showReflectorJSON)
		subPath.GET("/json/stations", d.showStationDataJSON)
		subPath.GET("/json/status", d.showStatus) //new spot for status
		subPath.GET("/json/links", d.showLinksDataJSON)
		subPath.GET("/json/modulesinuse", d.showModulesInUseJSON)
		subPath.GET("/json/peers", d.showPeers)
		subPath.GET("/json/metadata", d.showMetadata)
		subPath.GET("/json/countries", d.showCountries)
		subPath.Static("/assets", "frontend/spa/assets")
		subPath.Static("/icons", "frontend/spa/icons")
		subPath.StaticFile("/index.html", "frontend/spa/index.html")
		subPath.StaticFile("/", "frontend/spa/index.html")
		subPath.StaticFile("/favicon.ico", "frontend/spa/favicon.ico")

	}
	d.Router = router

	return &d, nil
}
