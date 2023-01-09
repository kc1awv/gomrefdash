package dashboard

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Country struct {
	Country          string   `json:"country"`
	CountryCode      string   `json:"countrycode"`
	CallsignPrefixes []string `json:"-"`
}

type CountryCallsigns struct {
	Countries []Country
}

func NewCountryCallsignsFromFile(filename string) (*CountryCallsigns, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var countries []Country
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")
		country := fields[0]
		countryCode := strings.ToLower(fields[1])
		callsignPrefixes := strings.Split(fields[2], "-")
		countries = append(countries, Country{country, countryCode, callsignPrefixes})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &CountryCallsigns{Countries: countries}, nil
}

// GetCountryFromCallsign returns the Country struct for the given callsign and error if not found.
func (cc *CountryCallsigns) GetCountryFromCallsign(callsign string) (Country, error) {
	for _, country := range cc.Countries {
		for _, callsignPrefix := range country.CallsignPrefixes {
			if strings.HasPrefix(callsign, callsignPrefix) {
				return country, nil
			}
		}
	}
	return Country{}, fmt.Errorf("country for callsign not found")
}
