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
	App_id        string    `json:id`
	Latitude      float32   `json:latitude`
	Longitude     float32   `json:longitude`
	Downloaded_at time.Time `json:downloaded_at`
}

func (d Downloads) Find(start string, end string) []Download {
	const layout = "2006-01-02"
	startAsTime, errStart := time.Parse(layout, start)
	endAsTime, errEnd := time.Parse(layout, end)

	if errStart != nil {
		startAsTime = time.Now().AddDate(0, 0, -7)
	}

	if errEnd != nil {
		endAsTime = time.Now().AddDate(0, 0, 7)
	}

	return d.Filter(func(item Download) bool {
		return item.HasBeenDownloadeWithin(startAsTime, endAsTime)
	})
}

func (d Download) HasBeenDownloadeWithin(start time.Time, end time.Time) bool {
	return d.HasBeenDownloadeBefore(end) && d.HasBeenDownloadeAfter(start)
}

func (d Download) HasBeenDownloadeBefore(t time.Time) bool {
	return d.Downloaded_at.Before(t) || d.Downloaded_at.Equal(t)
}

func (d Download) HasBeenDownloadeAfter(t time.Time) bool {
	return d.Downloaded_at.After(t) || d.Downloaded_at.Equal(t)
}

func (d *Downloads) Init() *Downloads {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10000; i++ {
		var lat = float32(r.Intn(180) - 90)
		if lat < 0 {
			lat = lat + r.Float32()
		} else {
			lat = lat - r.Float32()
		}
		var lon = float32(r.Intn(359)) + r.Float32()
		var downloaded_at = randate(r)
		d.downloads = append(d.downloads, Download{strconv.Itoa(i + 1), lat, lon, downloaded_at})
	}
	return d
}

func (d Downloads) Filter(f func(Download) bool) []Download {
	vsf := make([]Download, 0)
	for _, v := range d.downloads {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func randate(r *rand.Rand) time.Time {
	min := time.Date(2017, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	var now = time.Now()

	max := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
}
