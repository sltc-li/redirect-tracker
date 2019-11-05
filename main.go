package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

type RedirectTracker interface {
	Track(ctx context.Context, url string) ([]TrackResult, error)
}

var (
	tracker RedirectTracker
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:\n\n\t%s https://google.com\n\n", os.Args[0])
		os.Exit(2)
	}

	tracker = NewRedirectTrackerByTransport()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	results, err := tracker.Track(ctx, os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}

	for _, r := range results {
		fmt.Printf("requesting %s (%s)\n", r.URL, r.Status)
	}
}
