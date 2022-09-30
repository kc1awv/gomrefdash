package main

import (
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
  "log"
)

var router *gin.Engine

func main() {

  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  
  router = gin.Default()

  router.LoadHTMLGlob("templates/*")
  router.StaticFile("/favicon.ico", "static/favicon.ico")
  router.GET("/", showIndexPage)
  router.GET("/links/", showLinksPage)
  router.GET("/peers/", showPeersPage)

  log.Fatal(router.Run())

}