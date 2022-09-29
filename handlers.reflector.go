package main

import (
  "net/http"
	"os"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  info     := getInfo()
	modules  := getModules()
	stations := getStations()
	ipv4     := os.Getenv("IPV4")
	ipv6     := os.Getenv("IPV6")
  c.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "title":     "Home Page",
      "info":      info,
			"modules":   modules,
			"stations":  stations,
			"ipv4":      ipv4,
			"ipv6":      ipv6,
    },
  )
}

func showLinksPage(c *gin.Context) {
  nodes := getNodes()
	ipv4  := os.Getenv("IPV4")
	ipv6  := os.Getenv("IPV6")
  c.HTML(
    http.StatusOK,
    "links.html",
    gin.H{
      "title": "Links",
      "nodes": nodes,
			"ipv4":  ipv4,
			"ipv6":  ipv6,
    },
  )
}

func showPeersPage(c *gin.Context) {
  peers := getPeers()
	ipv4  := os.Getenv("IPV4")
	ipv6  := os.Getenv("IPV6")
  c.HTML(
    http.StatusOK,
    "peers.html",
    gin.H{
      "title": "Peers",
      "peers": peers,
			"ipv4":  ipv4,
			"ipv6":  ipv6,
    },
  )
}
