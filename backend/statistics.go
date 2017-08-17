package main

import (
	"time"

	"github.com/AsGz/geo/georeverse"
)

type Statistics struct {
	ByCountry map[string]int
	ByTime    map[string]int
	BySeason  map[string]int
	ByWeekDay map[string]int
	ByMonth   map[string]int
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
		s.setByTime(download.Downloaded_at)
		s.setSeason(download.Downloaded_at)
		s.setByWeekDay(download.Downloaded_at)
		s.setByMonth(download.Downloaded_at)
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
func (s *Statistics) setByTime(t time.Time) {
	if isMorning := t.Hour() >= 6 && t.Hour() < 12; isMorning {
		s.ByTime["MOR"] = s.ByTime["MOR"] + 1
	}
	if isAfternoon := t.Hour() >= 12 && t.Hour() < 18; isAfternoon {
		s.ByTime["AFT"] = s.ByTime["AFT"] + 1
	}
	if isEvening := t.Hour() >= 18 && t.Hour() < 24; isEvening {
		s.ByTime["EVE"] = s.ByTime["EVE"] + 1
	}
	if isNight := t.Hour() >= 0 && t.Hour() < 6; isNight {
		s.ByTime["NIG"] = s.ByTime["NIG"] + 1
	}
}

//given a statistics and a date set the season the date belongs to
//the doesn't take in consideration
func (s *Statistics) setSeason(t time.Time) {
	switch t.Month() {
	case time.January:
		fallthrough
	case time.February:
		fallthrough
	case time.March:
		s.BySeason["WIN"] = s.BySeason["WIN"] + 1
	case time.April:
		fallthrough
	case time.May:
		fallthrough
	case time.June:
		s.BySeason["SPR"] = s.BySeason["SPR"] + 1
	case time.July:
		fallthrough
	case time.August:
		fallthrough
	case time.September:
		s.BySeason["SUM"] = s.BySeason["SUM"] + 1
	case time.October:
		fallthrough
	case time.November:
		fallthrough
	case time.December:
		s.BySeason["AUT"] = s.BySeason["AUT"] + 1
	}
}

func (s *Statistics) setByWeekDay(t time.Time) {
	s.ByWeekDay[t.Weekday().String()] = s.ByWeekDay[t.Weekday().String()] + 1
}

func (s *Statistics) setByMonth(t time.Time) {
	s.ByMonth[t.Month().String()] = s.ByMonth[t.Month().String()] + 1
}

func newStatitistics() Statistics {
	byCountry := map[string]int{}
	byTime := map[string]int{}
	bySeason := map[string]int{}
	byWeekDay := map[string]int{}
	byMonth := map[string]int{}
	return Statistics{byCountry, byTime, bySeason, byWeekDay, byMonth}
}
