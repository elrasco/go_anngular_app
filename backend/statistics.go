package main

import (
	"time"

	"github.com/AsGz/geo/georeverse"
)

type Statistics struct {
	ByCountry map[string]int
	ByTime    map[string]int
}

func (d Downloads) getStatistics(start string, end string) Statistics {
	dataPath := "./data/polygons.properties"
	g, err := georeverse.NewCountryReverser(dataPath)
	if err != nil {
		panic(err)
	}

	s := newStatitistics()
	for _, download := range d.Find(start, end) {
		s.setByCountry(download, g)
		s.setByTime(download)
	}
	return s
}

func (s *Statistics) setByCountry(d Download, g *georeverse.CountryReverser) {
	if country := g.GetCountryCode(float64(d.Longitude), float64(d.Latitude)); len(country) > 0 {
		s.ByCountry[country] = s.ByCountry[country] + 1
	} else {
		s.ByCountry["UNKNOWN"] = s.ByCountry[country] + 1
	}
}

//groups the times of the downloads by "MOR", "AFT", "EVE", "NIG"
//NOTE: we need a service to calculate the offset from UTC given a point (lat, lon)
func (s *Statistics) setByTime(d Download) {
	if isMorning(d.Downloaded_at) {
		s.ByTime["MOR"] = s.ByTime["MOR"] + 1
	}
	if isAfternoon(d.Downloaded_at) {
		s.ByTime["AFT"] = s.ByTime["AFT"] + 1
	}
	if isEvening(d.Downloaded_at) {
		s.ByTime["EVE"] = s.ByTime["EVE"] + 1
	}
	if isNight(d.Downloaded_at) {
		s.ByTime["NIG"] = s.ByTime["NIG"] + 1
	}
}

func isMorning(t time.Time) bool {
	start := time.Date(t.Year(), t.Month(), t.Day(), 5, 0, 0, 0, time.UTC)
	end := time.Date(t.Year(), t.Month(), t.Day(), 11, 59, 0, 0, time.UTC)
	return isAfterOrEqual(t, start) && t.UTC().Before(end)
}

func isAfternoon(t time.Time) bool {
	start := time.Date(t.Year(), t.Month(), t.Day(), 12, 0, 0, 0, time.UTC)
	end := time.Date(t.Year(), t.Month(), t.Day(), 17, 59, 0, 0, time.UTC)
	return isAfterOrEqual(t, start) && t.UTC().Before(end)
}

func isEvening(t time.Time) bool {
	start := time.Date(t.Year(), t.Month(), t.Day(), 18, 0, 0, 0, time.UTC)
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 0, 0, time.UTC)
	return isAfterOrEqual(t, start) && t.UTC().Before(end)
}

func isNight(t time.Time) bool {
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
	end := time.Date(t.Year(), t.Month(), t.Day(), 5, 59, 0, 0, time.UTC)
	return isAfterOrEqual(t, start) && t.UTC().Before(end)
}

func isAfterOrEqual(t1 time.Time, t2 time.Time) bool {
	return t1.After(t2) || t1.Equal(t2)
}

func newStatitistics() Statistics {
	byCountry := map[string]int{}
	byTime := map[string]int{}
	return Statistics{byCountry, byTime}
}
