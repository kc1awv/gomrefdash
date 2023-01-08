package main

import (
	"flag"
	"log"
	"os"

	"github.com/kc1awv/gomrefdash/backend/pkg/dashboard"
	"github.com/kc1awv/gomrefdash/backend/pkg/dashboard/config"
)

var dver = "development"

func main() {
	log.Printf("Starting %s %s", os.Args[0], dver)

	configFilename := "gomrefdash.toml"
	// Defining a string flag
	configFilePath := flag.String("config", configFilename, "Filepath of the config file to read")

	// Call flag.Parse() to parse the command-line flags
	flag.Parse()

	searchPaths := []string{*configFilePath}

	userConfigDir, err := os.UserConfigDir()

	if err == nil {
		searchPaths = append(searchPaths, userConfigDir+string(os.PathSeparator)+configFilename)
	}

	cfg, err := config.NewConfigFromFile(searchPaths)
	if err != nil {
		log.Fatalf("unable to get config: %s", err)
	}
	dashboard, err := dashboard.NewDashboard(cfg, dver)
	if err != nil {
		log.Fatalf("unable to create dashboard: %s", err)
	}
	log.Fatal(dashboard.Router.Run(dashboard.Config.HostPort))
}
