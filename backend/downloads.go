package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Downloads struct {
	downloads []Download
}

type Download struct {
	App_id        string
	Latitude      float32
	Longitude     float32
	Downloaded_at time.Time
}

//Find return an array of Download, given an interval
func (d Downloads) Find(start string, end string) []Download {
	const layout = "2006-01-02"
	startAsTime, errStart := time.Parse(layout, start)
	endAsTime, errEnd := time.Parse(layout, end)

	if errStart != nil {
		startAsTime = time.Date(1970, 0, 0, 0, 0, 0, 0, time.UTC)
	}

	if errEnd != nil {
		endAsTime = time.Now().UTC()
	}

	return d.Filter(func(item Download) bool {
		return item.HasBeenDownloadeWithin(startAsTime, endAsTime)
	})
}

//HasBeenDownloadeWithin given an interval of time and a download , returns true if the download has been downloaded in a such interval, false otherwise
func (d Download) HasBeenDownloadeWithin(start time.Time, end time.Time) bool {
	return d.HasBeenDownloadeBefore(end) && d.HasBeenDownloadeAfter(start)
}

func (d Download) HasBeenDownloadeBefore(t time.Time) bool {
	return d.Downloaded_at.Before(t) || d.Downloaded_at.Equal(t)
}

func (d Download) HasBeenDownloadeAfter(t time.Time) bool {
	return d.Downloaded_at.After(t) || d.Downloaded_at.Equal(t)
}

func (d Downloads) Filter(f func(Download) bool) []Download {
	choosen := make([]Download, 0)
	for _, v := range d.downloads {
		if f(v) {
			choosen = append(choosen, v)
		}
	}
	return choosen
}

//Init initialize the set of downloads with 1000 items.
//The moment of each download is within this year
func (d *Downloads) Init() *Downloads {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 1000; i++ {
		d.add(generateDownload(i, r))
	}
	return d
}

func (d *Downloads) add(download Download) {
	d.downloads = append(d.downloads, download)
}

//return a Download struct with random lat, lon and downloaded_at (Locale.UTC)
func generateDownload(id int, r *rand.Rand) Download {
	var lat, lon = randcoordinate(r)
	var downloadedAt = randate(r)
	return Download{strconv.Itoa(id + 1), lat, lon, downloadedAt}
}

func randcoordinate(r *rand.Rand) (float32, float32) {
	var randomRegion = ChooseARegion(r)
	return randomRegion.Coordinate(r)
}

//return a random Time within this Year (from now back)
func randate(r *rand.Rand) time.Time {
	var now = time.Now().UTC()
	min := now.AddDate(-1, 0, 0).Unix()
	max := now.Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
}
