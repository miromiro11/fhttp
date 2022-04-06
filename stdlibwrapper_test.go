package http

import (
	"context"
	"errors"
	"net/http"
	"testing"
)

func TestStdlibTransport(t *testing.T) {
	var txp http.RoundTripper = &StdlibTransport{
		Transport: &Transport{},
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
