package main

import (
	"testing"
	"time"
)

func TestSetByTime(t *testing.T) {
	s := newStatitistics()

	t1 := time.Date(2015, time.April, 8, 6, 0, 0, 0, time.UTC)
	s.setByTime(t1)
	if morning := s.ByTime["MOR"] == 1; !morning {
		t.Errorf("%s is expected to be in morning", t1)
	}

	t1 = time.Date(2015, time.April, 8, 17, 59, 0, 0, time.UTC)
	s.setByTime(t1)
	if afternoon := s.ByTime["AFT"] == 1; !afternoon {
		t.Errorf("%s is expected to be in afternoon", t1)
	}

	t1 = time.Date(2015, time.April, 8, 18, 0, 0, 0, time.UTC)
	s.setByTime(t1)
	if evening := s.ByTime["EVE"] == 1; !evening {
		t.Errorf("%s is expected to be in evening", t1)
	}

	t1 = time.Date(2015, time.April, 8, 24, 0, 0, 0, time.UTC)
	s.setByTime(t1)
	if night := s.ByTime["NIG"] == 1; !night {
		t.Errorf("%s is expected to be in night", t1)
	}
}

func TestSetSeason(t *testing.T) {
	t.Log("should set the season of a time")
	s := newStatitistics()
	t1 := time.Date(2017, time.October, 1, 0, 0, 0, 0, time.UTC)
	s.setSeason(t1)
	if win := s.BySeason["AUT"]; win == 0 {
		t.Errorf("%s is expected to be an autumnal month", time.October)
	}

	t1 = time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	s.setSeason(t1)
	if win := s.BySeason["WIN"]; win == 0 {
		t.Errorf("%s is expected to be an wintry month", time.January)
	}

	t1 = time.Date(2017, time.April, 1, 0, 0, 0, 0, time.UTC)
	s.setSeason(t1)
	if spr := s.BySeason["SPR"]; spr == 0 {
		t.Errorf("%s is expected to be an springy month", time.April)
	}

	t1 = time.Date(2017, time.July, 1, 0, 0, 0, 0, time.UTC)
	s.setSeason(t1)
	if sum := s.BySeason["SUM"]; sum == 0 {
		t.Errorf("%s is expected to be a summer month", time.July)
	}
}
