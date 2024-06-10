package main

import (
	"hash/fnv"
	"log/slog"
	"os"
	"time"
)

var httpTimeFormats = []string{
	"Mon, _2 Jan 2006 15:04:05 GMT",
	time.RFC850,
	time.ANSIC,
}

func parseHTTPTime(text string) (int64, bool) {
	if text == "" {
		return 0, true
	}

	for _, layout := range httpTimeFormats {
		if t, err := time.Parse(layout, text); err == nil {
			return t.Unix(), true
		}
	}

	return 0, false
}

func hash(s string) uint {
	h := fnv.New32a()
	h.Write([]byte(s))
	return uint(h.Sum32())
}

func fatal(msg string, attr ...any) {
	slog.Error(msg, attr...)
	os.Exit(1)
}