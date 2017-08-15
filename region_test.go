package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomcoo(t *testing.T) {
	t.Log("should return a random float32 within a range")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10000; i++ {
		min := r.Intn(10)
		max := r.Intn(10) + 10
		if res := Randomcoo(min, max, r); res > float32(max) || res < float32(min) {
			t.Errorf("Expected to be within the range")
		}
	}
}

func TestWhoamiEU(t *testing.T) {
	t.Log("should return the name of the region")
	region := EURegion{}
	if whoami := region.Whoami(); whoami != "EU" {
		t.Errorf("Expected to be EU")
	}
}

func TestWhoamiUS(t *testing.T) {
	t.Log("should return the name of the region")
	region := USRegion{}
	if whoami := region.Whoami(); whoami != "US" {
		t.Errorf("Expected to be US")
	}
}

func TestCoordinateEU(t *testing.T) {
	t.Log("should return a point inside Europe")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	region := EURegion{}
	for i := 0; i < 10000; i++ {
		var lat, lon = region.Coordinate(r)
		if int(lat) < 45 || int(lat) > 53 || int(lon) < 6 || int(lon) > 29 {
			t.Errorf("Expected to be in EU")
		}
	}
}

func TestCoordinateUS(t *testing.T) {
	t.Log("should return a point inside Europe")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	region := USRegion{}
	for i := 0; i < 10000; i++ {
		var lat, lon = region.Coordinate(r)
		if int(lat) < 32 || int(lat) > 48 || int(lon) < -117 || int(lon) > -81 {
			t.Errorf("Expected to be in US")
		}
	}
}
