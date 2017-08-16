package main

import "math/rand"

type Region interface {
	//return a random position (latitude, longitude)
	Coordinate(r *rand.Rand) (float32, float32)
	Whoami() string
}

type USRegion struct{}

//return a random position (latitude, longitude) in US.
//given this 2 coordinates: (48N, 117W) and (32N, 81W), it's easy to draw a rectangle inside US
func (region USRegion) Coordinate(r *rand.Rand) (float32, float32) {
	return Randomcoo(32, 48, r), Randomcoo(-117, -81, r)
}

func (region USRegion) Whoami() string {
	return "US"
}

type EURegion struct{}

//return a random position (latitude, longitude) in EU.
//given this 2 coordinates: (53N, 6E) and (45N, 29E), it's easy to draw a rectangle inside EU
func (region EURegion) Coordinate(r *rand.Rand) (float32, float32) {
	return Randomcoo(45, 53, r), Randomcoo(6, 29, r)
}

func (region EURegion) Whoami() string {
	return "EU"
}

type ASIARegion struct{}

//return a random position (latitude, longitude) in ASIA.
//given this 2 coordinates: (27N, 57E) and (64N, 117E), it's easy to draw a rectangle inside EU
func (region ASIARegion) Coordinate(r *rand.Rand) (float32, float32) {
	return Randomcoo(27, 64, r), Randomcoo(57, 117, r)
}

func (region ASIARegion) Whoami() string {
	return "ASIA"
}

//return a random float between min and max plus a value between [0, 1)
func Randomcoo(min int, max int, r *rand.Rand) float32 {
	return float32(r.Intn(max-min)+min) + r.Float32()
}

//return a random region between US and Europe
func ChooseARegion(r *rand.Rand) Region {
	regions := []Region{USRegion{}, EURegion{}, ASIARegion{}}
	return regions[r.Intn(3)]
}
