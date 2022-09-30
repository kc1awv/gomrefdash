package main

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "os"
  "strings"
)

type Reflector struct {
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

func getReflector() Reflector {
  // xmlFile, err := os.Open("mrefd.xml")
  xmlFile, err := os.Open("/var/log/mrefd.xml")
  if err != nil {
      fmt.Println(err)
  }
  defer xmlFile.Close()
  byteValue, _ := ioutil.ReadAll(xmlFile)
  var data Reflector
  xml.Unmarshal(byteValue, &data)
  return data
}

func getInfo() map[string][]string {
  data := getReflector()
  info := make(map[string][]string)
  info[data.Callsign] = append(info[data.Callsign], data.Version)
  return info
}

func getModules() map[string][]string {
  data := getReflector()
  modules := make(map[string][]string)
  for _, node := range data.Nodes {
    modules[node.LinkedModule] = append(modules[node.LinkedModule], node.Callsign)
  }
  return modules
}

func getNodes() map[string][]string {
  data := getReflector()
  nodes := make(map[string][]string)
  for _, node := range data.Nodes {
    nodes[node.Callsign] = append(nodes[node.Callsign], node.IPAddress, node.LinkedModule, node.Protocol, node.ConnectTime, node.LastHeardTime)
  }
  return nodes
}

func getPeers() map[string][]string {
  data := getReflector()
  peers := make(map[string][]string)
  for _, peer := range data.Peers {
    peers[peer.Callsign] = append(peers[peer.Callsign], peer.IPAddress, peer.LinkedModule, peer.ConnectTime, peer.LastHeardTime)
  }
  return peers
}

func getStations() map[int][]string {
  data := getReflector()
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
