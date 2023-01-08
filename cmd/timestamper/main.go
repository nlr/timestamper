package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Timestamp struct {
	Unix string `json:"unix"`
	Utc  string `json:"utc"`
}

func main() {
	router := httprouter.New()
	router.GET("/api/:date", handleDate)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func handleDate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	dateStr := ps.ByName("date")

	w.Header().Set("Content-Type", "application/json")

	t, err := parseDate(dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := make(map[string]string)
		resp["message"] = "Bad Request"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)
}

func parseDate(date string) (*Timestamp, error) {
	var err error
	const layoutISO = "Mon, 02 Jan 2006 00:00:00 GMT"

	if unix, err := parseUnix(date); err == nil {
		return &Timestamp{Unix: strconv.FormatInt(unix, 10), Utc: time.UnixMilli(unix).Format(layoutISO)}, nil
	}

	if utc, err := parseUtc(date); err == nil {
		return &Timestamp{Unix: strconv.FormatInt(utc.UnixMilli(), 10), Utc: utc.Format(layoutISO)}, nil
	}

	return &Timestamp{Unix: "null", Utc: "null"}, err
}

func parseUnix(date string) (int64, error) {
	unix, err := strconv.ParseInt(date, 10, 64)
	if err != nil {
		return 0, err
	}
	return unix, nil
}

func parseUtc(date string) (time.Time, error) {
	const layoutISO = "2006-01-02"
	utc, err := time.Parse(layoutISO, date)
	if err != nil {
		return time.Time{}, err
	}
	return utc, nil
}
