package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

var reflector *Reflector // don't really like global variables, but don't see a way around this atm.

// this will set some global things
func initReflector() {
  //attempt to get the mrefd.xml location from MREFDFILE
  mrefFilePath := os.Getenv("MREFDFILE")
  if mrefFilePath == "" {
    mrefFilePath = "/var/log/mrefd.xml"
    log.Printf("MREFDFILE was not set, so using %s", mrefFilePath)
  }
  r, err := NewReflectorFromFile(mrefFilePath)
  if err != nil {
    log.Fatal("unable to create reflector: %w", err)
  }
  reflector = r // set the global variable
  
}

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("error loading .env file: %s", err)
  }
  initReflector()
  gin.SetMode(gin.ReleaseMode)
  router = gin.Default()
  router.LoadHTMLGlob("templates/*")
  router.StaticFile("/favicon.ico", "static/favicon.ico")
  router.GET("/", showIndexPage)
  router.GET("/links/", showLinksPage)
  router.GET("/peers/", showPeersPage)
  router.GET("/status", showStatus) //dashboard status metadata
  router.GET("/reflectordata", showReflectorJSON)
  
  log.Fatal(router.Run())

}