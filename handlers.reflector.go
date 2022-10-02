package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func showStatus(c *gin.Context) { 
  r := reflector
  r.refreshIfNeeded()
  status := struct{
    LastUpdateUnixTime int64 `json:"lastupdate"`
    LastDataUpdateUnixTime int64 `json:"lastmrefdupdate"`
  }{
    LastUpdateUnixTime: r.LastUpdateCheckTime.Unix(),
    LastDataUpdateUnixTime: r.ReflectorData.FileTime.Unix(),
  }

  c.JSON(200, status)
}

func showIndexPage(c *gin.Context) {
  r := reflector

  info      := r.GetInfo()
  modules   := r.GetModules()
  stations  := r.GetStations()
  ipv4      := os.Getenv("IPV4")
  ipv6      := os.Getenv("IPV6")
  refresh   := os.Getenv("REFRESH")
  email     := os.Getenv("EMAIL")
  c.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "title":     "Last Heard",
      "version":   dver,
      "info":      info,
      "email":     email,
      "modules":   modules,
      "stations":  stations,
      "ipv4":      ipv4,
      "ipv6":      ipv6,
      "refresh":   refresh,
    },
  )
}

func showLinksPage(c *gin.Context) {
  r := reflector
  info  := r.GetInfo()
  nodes := r.GetNodes()
  ipv4  := os.Getenv("IPV4")
  ipv6  := os.Getenv("IPV6")
  c.HTML(
    http.StatusOK,
    "links.html",
    gin.H{
      "title":     "Links",
      "version":   dver,
      "info":      info,
      "nodes":     nodes,
      "ipv4":      ipv4,
      "ipv6":      ipv6,
    },
  )
}

func showPeersPage(c *gin.Context) {
  r := reflector
  info  := r.GetInfo()
  peers := r.GetPeers()
  ipv4  := os.Getenv("IPV4")
  ipv6  := os.Getenv("IPV6")
  c.HTML(
    http.StatusOK,
    "peers.html",
    gin.H{
      "title":     "Peers",
      "version":   dver,
      "info":      info,
      "peers":     peers,
      "ipv4":      ipv4,
      "ipv6":      ipv6,
    },
  )
}
