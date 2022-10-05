package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (d *Dashboard) showStatus(c *gin.Context) {
	r := d.Reflector
	r.refreshIfNeeded()
	status := struct {
		LastUpdateUnixTime     int64 `json:"lastupdate"`
		LastDataUpdateUnixTime int64 `json:"lastmrefdupdate"`
	}{
		LastUpdateUnixTime:     r.LastUpdateCheckTime.Unix(),
		LastDataUpdateUnixTime: r.ReflectorData.FileTime.Unix(),
	}

	c.JSON(200, status)
}

// showReflectorJSON returns the entire reflector stuct in JSON.
func (d *Dashboard) showReflectorJSON(c *gin.Context) {
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	c.JSON(200, d.Reflector)
}

func (d *Dashboard) showStationDataJSON(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	type stationData struct {
		Callsign       string `json:"callsign"`
		CallsignSuffix string `json:"callsignsuffix"`
		ViaNode        string `json:"vianode"`
		OnModule       string `json:"onmodule"`
		LastHeard      string `json:"lastheard"`
	}
	var data struct {
		Stations []stationData `json:"stations"`
	}
	for _, station := range d.Reflector.ReflectorData.Stations {
		callsignSplit := strings.Split(station.Callsign, " ")
		if len(callsignSplit) < 2 {
			continue
		}
		data.Stations = append(data.Stations, stationData{
			Callsign:       callsignSplit[0],
			CallsignSuffix: callsignSplit[1],
			ViaNode:        station.ViaNode,
			OnModule:       station.OnModule,
			LastHeard:      station.LastHeardTime,
		})
	}
	c.JSON(200, data)
}

func (d *Dashboard) showPeers(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	d.Reflector.Lock.Lock()
	defer d.Reflector.Lock.Unlock()
	c.JSON(200, d.Reflector.ReflectorData.Peers)
}

func (d *Dashboard) showLinksDataJSON(c *gin.Context) {
	d.Reflector.refreshIfNeeded()
	c.JSON(200, d.Reflector.ReflectorData.Nodes)
}

func (d *Dashboard) showModulesInUseJSON(c *gin.Context) {
	type moduleInfo struct {
		Name      string   `json:"name"`
		Callsigns []string `json:"callsigns"`
	}

	var moduleData []moduleInfo

	modules := d.Reflector.GetModules()
	for k, v := range modules {
		moduleData = append(moduleData, moduleInfo{Name: k, Callsigns: v})
	}
	c.JSON(200, moduleData)
}

func (d *Dashboard) showIndexPage(c *gin.Context) {
	r := d.Reflector
	info := r.GetInfo()
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Last Heard",
			"version": dver,
			"info":    info,
			"config":  d.Config,
		},
	)
}

func (d *Dashboard) showLinksPage(c *gin.Context) {
	r := d.Reflector
	info := r.GetInfo()
	c.HTML(
		http.StatusOK,
		"links.html",
		gin.H{
			"title":   "Links",
			"version": dver,
			"info":    info,
			"config":  d.Config,
		},
	)
}

func (d *Dashboard) showPeersPage(c *gin.Context) {
	r := d.Reflector
	info := r.GetInfo()
	c.HTML(
		http.StatusOK,
		"peers.html",
		gin.H{
			"title":   "Peers",
			"version": dver,
			"info":    info,
			"config":  d.Config,
		},
	)
}
