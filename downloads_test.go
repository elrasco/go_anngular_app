package main

import (
	"testing"
	"time"
)

func TestInitOK(t *testing.T) {
	t.Log("should create the database")
	var d = Downloads{}
	if l := len(d.Init().downloads); l < 1000 {
		t.Errorf("Expected downloads big enough")
	}
}

func TestFilterOK(t *testing.T) {
	t.Log("should choose the ones that satisfy the filter function")
	var d = Downloads{}
	d.downloads = []Download{Download{"1", 90.124, 290.00, time.Now()}}

	var filtered = d.Filter(func(item Download) bool {
		return true
	})

	if l := len(filtered); l == 0 {
		t.Errorf("Expected length greater than zero")
	}
}

func TestFilterKO(t *testing.T) {
	t.Log("should return empty array if function condition fails")
	var d = Downloads{}
	d.downloads = []Download{Download{"1", 90.124, 290.00, time.Now()}}

	var filtered = d.Filter(func(item Download) bool {
		return false
	})

	if l := len(filtered); l > 0 {
		t.Errorf("Expected length be zero")
	}
}

func TestHasBeenDownloadeAfterOk(t *testing.T) {
	t.Log("should return true if downloaded_at is after or equal the given date")

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, "2017-05-10")

	var downloaded_at, _ = time.Parse(layout, "2017-05-10")
	var d = Download{"1", 90.124, 290.00, downloaded_at}

	if hasBeenDowloadAfter := d.HasBeenDownloadeAfter(start); !hasBeenDowloadAfter {
		t.Errorf("Expected %s after of %s", d.Downloaded_at, start)
	}

}

func TestHasBeenDownloadeAfterKO(t *testing.T) {
	t.Log("should return false if downloaded_at is before the given date")

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, "2017-05-11")

	var downloaded_at, _ = time.Parse(layout, "2017-05-10")
	var d = Download{"1", 90.124, 290.00, downloaded_at}

	if hasBeenDowloadAfter := d.HasBeenDownloadeAfter(start); hasBeenDowloadAfter {
		t.Errorf("Expected %s after of %s", d.Downloaded_at, start)
	}

}

func TestHasBeenDownloadeBeforeOK(t *testing.T) {
	t.Log("should return true if downloaded_at is before or equal the given date")

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, "2017-05-10")

	var downloaded_at, _ = time.Parse(layout, "2017-05-10")
	var d = Download{"1", 90.124, 290.00, downloaded_at}

	if hasBeenDowloadBefore := d.HasBeenDownloadeBefore(start); !hasBeenDowloadBefore {
		t.Errorf("Expected %s before of %s", d.Downloaded_at, start)
	}
}

func TestHasBeenDownloadeBeforeKO(t *testing.T) {
	t.Log("should return false if downloaded_at is after the given date")

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, "2017-05-10")

	var downloaded_at, _ = time.Parse(layout, "2017-05-11")
	var d = Download{"1", 90.124, 290.00, downloaded_at}

	if hasBeenDowloadBefore := d.HasBeenDownloadeBefore(start); hasBeenDowloadBefore {
		t.Errorf("Expected %s before of %s", d.Downloaded_at, start)
	}
}

func TestHasBeenDownloadeWithinOK(t *testing.T) {
	t.Log("should return true if downloaded_at is between a given range")

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, "2017-05-10")
	end, _ := time.Parse(layout, "2017-05-12")

	var downloaded_at, _ = time.Parse(layout, "2017-05-12")
	var d = Download{"1", 90.124, 290.00, downloaded_at}

	if hasBeenDownloadeWithin := d.HasBeenDownloadeWithin(start, end); !hasBeenDownloadeWithin {
		t.Errorf("Expected %s between %s and ", d.Downloaded_at, start, end)
	}
}

func TestHasBeenDownloadeWithinKO(t *testing.T) {
	t.Log("should return false if downloaded_at is out a given range")

	const layout = "2006-01-02"
	start, _ := time.Parse(layout, "2017-05-10")
	end, _ := time.Parse(layout, "2017-05-12")

	var downloaded_at, _ = time.Parse(layout, "2017-05-13")
	var d = Download{"1", 90.124, 290.00, downloaded_at}

	if hasBeenDownloadeWithin := d.HasBeenDownloadeWithin(start, end); hasBeenDownloadeWithin {
		t.Errorf("Expected %s out %s and %s", d.Downloaded_at, start, end)
	}
}

func TestFindOK(t *testing.T) {
	t.Log("should return an array of downloads within the range")

	const layout = "2006-01-02"
	start := "2017-05-10"
	end := "2017-05-12"

	var downloaded_at, _ = time.Parse(layout, "2017-05-12")
	var d = Downloads{[]Download{Download{"1", 90.124, 290.00, downloaded_at}}}

	if l := len(d.Find(start, end)); l != 1 {
		t.Errorf("Expected length eq 1")
	}
}

func TestFindDefaultValues(t *testing.T) {
	t.Log("should return an array of downloads within the range")

	const layout = "2006-01-02"
	start := ""
	end := ""

	var downloaded_at = time.Now()
	var d = Downloads{[]Download{Download{"1", 90.124, 290.00, downloaded_at}}}

	if l := len(d.Find(start, end)); l != 1 {
		t.Errorf("Expected length eq 1")
	}
}

func TestFindKO(t *testing.T) {
	t.Log("should return an empty array if out of range")

	const layout = "2006-01-02"
	start := "2017-05-10"
	end := "2017-05-12"

	var downloaded_at, _ = time.Parse(layout, "2017-05-13")
	var d = Downloads{[]Download{Download{"1", 90.124, 290.00, downloaded_at}}}

	if l := len(d.Find(start, end)); l == 1 {
		t.Errorf("Expected length eq 0")
	}
}
