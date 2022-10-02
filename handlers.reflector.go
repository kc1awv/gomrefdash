package main

import (
	"net/http"
	"os"
	"strings"

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

// showReflectorJSON returns the entire reflector stuct in JSON.
func showReflectorJSON(c *gin.Context) {
  reflector.Lock.Lock()
  defer reflector.Lock.Unlock()
  c.JSON(200, reflector)
}

func showStationDataJSON(c *gin.Context) {
  reflector.Lock.Lock()
  defer reflector.Lock.Unlock()
  type stationData struct {
    Callsign string `json:"callsign"`
    CallsignSuffix string `json:"callsignsuffix"`
    LinkPeer string `json:"linkpeer"` 
    OnModule string `json:"onmodule"`
    LastHeard string `json:"lastheard"`
  }
  var data struct {
    Stations []stationData `json:"stations"`
  }
  for _, station := range reflector.ReflectorData.Stations {
    callsignSplit := strings.Split(station.Callsign, " ")
    if len(callsignSplit) < 2 {
      continue
    }
    data.Stations = append(data.Stations, stationData{
      Callsign: callsignSplit[0],
      CallsignSuffix: callsignSplit[1],
      LinkPeer: station.ViaPeer,
      OnModule: station.OnModule,
      LastHeard: station.LastHeardTime,
    })
  }
  c.JSON(200, data)
}

func showPeers(c *gin.Context) {
  reflector.Lock.Lock()
  defer reflector.Lock.Unlock()
  c.JSON(200, reflector.ReflectorData.Peers)
}

func showLinksDataJSON(c *gin.Context) {
  c.JSON(200, reflector.ReflectorData.Nodes)
}

func showModulesInUseJSON(c *gin.Context) {
  type moduleInfo struct {
    Name string `json:"name"`
    Callsigns []string `json:"callsigns"`
  }

  var moduleData []moduleInfo

  modules := reflector.GetModules()
  for k, v := range modules {
    moduleData = append(moduleData, moduleInfo{Name: k, Callsigns: v})
  }
  c.JSON(200, moduleData)
}

func showIndexPage(c *gin.Context) {
  r := reflector

  info      := r.GetInfo()
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
  email := os.Getenv("EMAIL")
  c.HTML(
    http.StatusOK,
    "links.html",
    gin.H{
      "title":     "Links",
      "version":   dver,
      "info":      info,
      "email":     email,
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
  email := os.Getenv("EMAIL")
  c.HTML(
    http.StatusOK,
    "peers.html",
    gin.H{
      "title":     "Peers",
      "version":   dver,
      "info":      info,
      "email":     email,
      "peers":     peers,
      "ipv4":      ipv4,
      "ipv6":      ipv6,
    },
  )
}
