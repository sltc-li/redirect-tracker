package main

import (
	"context"
	"net/http"
)

type RedirectTrackerByTransport struct {
}

func NewRedirectTrackerByTransport() *RedirectTrackerByTransport {
	return &RedirectTrackerByTransport{}
}

func (rt *RedirectTrackerByTransport) Track(ctx context.Context, url string) ([]TrackResult, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	transport := &RedirectTrackTransport{Base: http.DefaultTransport}
	client := http.Client{
		Transport:     transport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	_, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	return transport.Redirects, nil
}

type RedirectTrackTransport struct {
	Base      http.RoundTripper
	Redirects []TrackResult
}

func (rtt *RedirectTrackTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := rtt.Base.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	rtt.Redirects = append(rtt.Redirects, TrackResult{
		Status: resp.Status,
		URL:    req.URL.String(),
	})
	return resp, err
}
