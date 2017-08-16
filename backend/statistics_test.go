package main

import (
	"testing"
	"time"
)

func TestIsMorning(t *testing.T) {
	now := time.Now()
	t1 := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.UTC)

	if morning := isMorning(t1); !morning {
		t.Errorf("%s is expected to be in morning", t1)
	}
}

func TestIsAfternoon(t *testing.T) {
	now := time.Now()
	t1 := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)

	if afternoon := isAfternoon(t1); !afternoon {
		t.Errorf("%s is expected to be in afternoon", t1)
	}
}

func TestIsEvening(t *testing.T) {
	now := time.Now()
	t1 := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.UTC)

	if evening := isEvening(t1); !evening {
		t.Errorf("%s is expected to be in evening", t1)
	}
}

func TestIsNight(t *testing.T) {
	now := time.Now()
	t1 := time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.UTC)

	if night := isNight(t1); !night {
		t.Errorf("%s is expected to be in night", t1)
	}
}

func TestShouldBelongsToATime(t *testing.T) {
	t.Log("should belong to a time interval")
	now := time.Now()
	t1 := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, time.UTC)

	if belongs := isMorning(t1) || isAfternoon(t1) || isEvening(t1) || isNight(t1); !belongs {
		t.Errorf("%s is expected to be in the morning or in the afternoon or in the evening or in the night", t1)
	}
}
