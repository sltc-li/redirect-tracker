package main

import (
	"context"
	"log"
	"net/http"
)

type RedirectTrackerByCheckRedirect struct {
}

func NewRedirectTrackerByCheckRedirect() *RedirectTrackerByCheckRedirect {
	return &RedirectTrackerByCheckRedirect{}
}

func (rt *RedirectTrackerByCheckRedirect) Track(ctx context.Context, url string) ([]TrackResult, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	results := []TrackResult{{URL: url}}
	client := http.Client{
		Transport: nil,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// panic: `via[len(via)-1].Response` is nil
			log.Printf("%#v", via[len(via)-1].Response)
			results[len(results)-1].Status = via[len(via)-1].Response.Status
			results = append(results, TrackResult{URL: req.URL.String()})
			return nil
		},
		Jar:     nil,
		Timeout: 0,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	results[len(results)-1].Status = resp.Status
	return results, nil
}
