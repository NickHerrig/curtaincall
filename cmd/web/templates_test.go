package main

import (
    "testing"
    "time"
)

func TestHumanDate(t *testing.T) {
    tm := time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC)
    hd := humanDate(tm)
    if hd != "17 Dec 2020 at 10:00" {
        t.Errorf("want %q: got %q", "17 Dec 2020 at 10:00", hd)
    }
}
