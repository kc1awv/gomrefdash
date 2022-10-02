package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Reflector struct {
  Lock sync.Mutex // used during cache refreshes or when requested by client
  LastUpdateCheckTime time.Time // the last time update was checked
  ReflectorData ReflectorData
  ReflectorFilePath string // contains the FilePath of reflector file. mrefd.xml
}

type ReflectorData struct {
  FileTime time.Time // populated after data is loaded (this isn't in the xml file)
  Callsign string    `xml:"CALLSIGN,attr"`
  Version  string    `xml:"VERSION"`
  Peers    []struct {
    Callsign      string   `xml:"CALLSIGN"`
    IPAddress     string   `xml:"IP"`
    LinkedModule  string   `xml:"LINKEDMODULE"`
    ConnectTime   string   `xml:"CONNECTTIME"`
    LastHeardTime string   `xml:"LASTHEARDTIME"`
  } `xml:"PEERS>PEER"`
  Nodes    []struct {
    Callsign      string   `xml:"CALLSIGN"`
    IPAddress     string   `xml:"IP"`
    LinkedModule  string   `xml:"LINKEDMODULE"`
    Protocol      string   `xml:"PROTOCOL"`
    ConnectTime   string   `xml:"CONNECTTIME"`
    LastHeardTime string   `xml:"LASTHEARDTIME"`
  } `xml:"NODES>NODE"`
  Stations []struct {
    Callsign      string   `xml:"CALLSIGN"`
    ViaNode       string   `xml:"VIANODE"`
    OnModule      string   `xml:"ONMODULE"`
    ViaPeer       string   `xml:"VIAPEER"`
    LastHeardTime string   `xml:"LASTHEARDTIME"`
  } `xml:"STATIONS>STATION"`
}

func NewReflectorDataFromFile(filePath string) (*ReflectorData, error) { 
  stat, err := os.Stat(filePath)
  if err != nil {
    return nil, fmt.Errorf("unable to stat file %s: %v", filePath, err)
  }
  b, err := os.ReadFile(filePath)
  if err != nil {
    return nil, fmt.Errorf("unable to read file %s: %v", filePath, err)
  }
  var data ReflectorData
  err = xml.Unmarshal(b, &data)
  if err != nil {
    return nil, fmt.Errorf("unable to unmarshal reflector data %s: %v", filePath, err)
  }
  // set the file date (used to compare later)
  data.FileTime = stat.ModTime()
  return &data, nil
}

// NewReflectorFromFile reads the xml file from filePath into the struct
// and sets the LastUpdateTime from the file's date.
func NewReflectorFromFile(filePath string) (*Reflector, error) {
  reflectorData, err := NewReflectorDataFromFile(filePath)
  if err != nil {
    return nil, fmt.Errorf("unable to obtain reflector data for %s: %v", filePath, err)
  }
  reflector := Reflector{ReflectorFilePath: filePath, LastUpdateCheckTime: time.Now(), ReflectorData: *reflectorData}
  return &reflector, nil
}

// Refreshes the data from xml file only if needed.
func (r *Reflector) refreshIfNeeded() {
  r.Lock.Lock()
  defer r.Lock.Unlock()
  r.LastUpdateCheckTime = time.Now()
  stat, err := os.Stat(r.ReflectorFilePath)
  if err != nil {
    log.Printf("unable to stat %s, %s", r.ReflectorFilePath, err)
    return
  }
  
  if !stat.ModTime().Equal(r.ReflectorData.FileTime) {
    log.Printf("refreshing mrefd data modtime: %v, last reflector data file time: %v", stat.ModTime(), r.ReflectorData.FileTime)
    d, err := NewReflectorDataFromFile(r.ReflectorFilePath)
    if err != nil {
      log.Printf("unable to refresh reflector data: %s", err)
      return
    }
    r.ReflectorData = *d
  }

}

func (r *Reflector) GetInfo() map[string][]string {
  //r.refreshIfNeeded()
  r.Lock.Lock()
  defer r.Lock.Unlock()
  data := r.ReflectorData
  info := make(map[string][]string)
  info[data.Callsign] = append(info[data.Callsign], data.Version)
  return info
}

func (r *Reflector) GetModules() map[string][]string {
  //r.refreshIfNeeded()
  r.Lock.Lock()
  defer r.Lock.Unlock()
  data := r.ReflectorData
  
  modules := make(map[string][]string)
  for _, node := range data.Nodes {
    modules[node.LinkedModule] = append(modules[node.LinkedModule], node.Callsign)
  }
  return modules
}

func (r *Reflector) GetNodes() map[string][]string {
  //r.refreshIfNeeded()
  r.Lock.Lock()
  defer r.Lock.Unlock()
  data := r.ReflectorData
  nodes := make(map[string][]string)
  for _, node := range data.Nodes {
    nodes[node.Callsign] = append(nodes[node.Callsign], node.IPAddress, node.LinkedModule, node.Protocol, node.ConnectTime, node.LastHeardTime)
  }
  return nodes
}

func (r *Reflector) GetPeers() map[string][]string {
  //r.refreshIfNeeded()
  r.Lock.Lock()
  defer r.Lock.Unlock()
  data := r.ReflectorData
  peers := make(map[string][]string)
  for _, peer := range data.Peers {
    peers[peer.Callsign] = append(peers[peer.Callsign], peer.IPAddress, peer.LinkedModule, peer.ConnectTime, peer.LastHeardTime)
  }
  return peers
}

func (r *Reflector) GetStations() map[int][]string {
  //r.refreshIfNeeded()
  r.Lock.Lock()
  defer r.Lock.Unlock()
  data := r.ReflectorData
  stations := make(map[int][]string)
  for i, station := range data.Stations {
    c := strings.Fields(station.Callsign)
    if len(c) < 2 {
      c = append(c, " ")
    }
    callsign, module := c[0], c[1]
    stations[i] = append(stations[i], callsign, module, station.ViaNode, station.OnModule, station.ViaPeer, station.LastHeardTime)
  }
  return stations
}
