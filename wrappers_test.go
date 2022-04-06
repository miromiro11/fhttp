package http_test

import (
	"context"
	"errors"
	fhttp "github.com/saucesteals/fhttp"
	"net/http"
	"testing"
)

func TestStdlibTransport(t *testing.T) {
	var txp http.RoundTripper = &fhttp.StdlibTransport{
		Transport: &fhttp.Transport{},
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // we expect the request with context to fail immediately
	req, err := http.NewRequestWithContext(
		ctx, "GET", "https://google.com", nil)
	if err != nil {
		t.Fatal(err) // should not fail
	}
	resp, err := txp.RoundTrip(req)
	if !errors.Is(err, context.Canceled) {
		t.Fatal("unexpected err", err)
	}
	if resp != nil {
		t.Fatal("unexpected resp")
	}

}

func TestFHttpTransport(t *testing.T) {
	var txp fhttp.RoundTripper = &fhttp.FHttpTransport{
		Transport: &http.Transport{},
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // we expect the request with context to fail immediately
	req, err := fhttp.NewRequestWithContext(
		ctx, "GET", "https://google.com", nil)
	if err != nil {
		t.Fatal(err) // should not fail
	}
	resp, err := txp.RoundTrip(req)
	if !errors.Is(err, context.Canceled) {
		t.Fatal("unexpected err", err)
	}
	if resp != nil {
		t.Fatal("unexpected resp")
	}

}
