package main

import (
	"testing"
)

func TestParseUnix(t *testing.T) {
	got, err := parseUnix("1654041600")
	if err != nil {
		t.Errorf("parseUnix(\"1654041600\") = %d; want %d, error: %v", got, 1654041600, err)
	}
}

func TestParseUtc(t *testing.T) {
	got, err := parseUtc("2015-12-25")
	expect := "2015-12-12 00:00:00 +0000 UTC"
	if err != nil {
		t.Errorf("parseUtc(\"2015-12-25\") = %v; expected: %s, error: %v", got, expect, err)
	}
}

func TestParseDate(t *testing.T) {
	got, err := parseDate("2023-01-08")
	want := Timestamp{Unix: "1673136000000", Utc: "Sun, 08 Jan 2023 00:00:00 GMT"}
	if *got != want {
		t.Errorf("got %v, wanted: %v, error: %v", got, want, err)
	}

	got, err = parseDate("1673136000000")
	if *got != want {
		t.Errorf("got %v, wanted %v, error: %v", got, want, err)
	}

	got, err = parseDate("23-23-23")
	want = Timestamp{Unix: "null", Utc: "null"}
	if *got != want {
		t.Errorf("got %v, wanted %v, error: %v", got, nil, err)
	}
}
