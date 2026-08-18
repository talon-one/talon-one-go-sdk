package talon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewAPIClient_DefaultTimeout(t *testing.T) {
	cfg := NewConfiguration()
	_ = NewAPIClient(cfg)

	if cfg.HTTPClient == nil {
		t.Fatal("expected HTTPClient to be set, got nil")
	}
	if cfg.HTTPClient.Timeout != DefaultHTTPTimeout {
		t.Errorf("expected default timeout %v, got %v", DefaultHTTPTimeout, cfg.HTTPClient.Timeout)
	}
}

func TestNewAPIClient_CustomClientPreserved(t *testing.T) {
	custom := &http.Client{Timeout: 90 * time.Second}
	cfg := NewConfiguration()
	cfg.HTTPClient = custom
	_ = NewAPIClient(cfg)

	if cfg.HTTPClient != custom {
		t.Error("expected custom HTTPClient to be preserved")
	}
	if cfg.HTTPClient.Timeout != 90*time.Second {
		t.Errorf("expected 90s timeout, got %v", cfg.HTTPClient.Timeout)
	}
}

func TestNewAPIClient_DoesNotUseDefaultClient(t *testing.T) {
	cfg := NewConfiguration()
	_ = NewAPIClient(cfg)

	if cfg.HTTPClient == http.DefaultClient {
		t.Error("NewAPIClient must not fall back to http.DefaultClient (Timeout=0)")
	}
}

func TestPrepareRequest_UsesContext(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := client.prepareRequest(ctx, "https://example.com/test", http.MethodGet, nil,
		map[string]string{}, nil, nil, nil)
	if err != nil {
		t.Fatalf("prepareRequest returned error: %v", err)
	}
	if req.Context() != ctx {
		t.Error("expected request to carry the provided context")
	}
}

func TestPrepareRequest_NilContextFallsBackToBackground(t *testing.T) {
	cfg := NewConfiguration()
	client := NewAPIClient(cfg)

	req, err := client.prepareRequest(nil, "https://example.com/test", http.MethodGet, nil,
		map[string]string{}, nil, nil, nil)
	if err != nil {
		t.Fatalf("prepareRequest returned error: %v", err)
	}
	if req.Context() == nil {
		t.Error("expected request to have a non-nil context even when nil is passed")
	}
}

func TestDefaultHTTPTimeout_EnforcedOnSlowServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.HTTPClient = &http.Client{Timeout: 200 * time.Millisecond}
	client := NewAPIClient(cfg)

	ctx := context.Background()
	req, err := client.prepareRequest(ctx, srv.URL+"/slow", http.MethodGet, nil,
		map[string]string{}, nil, nil, nil)
	if err != nil {
		t.Fatalf("prepareRequest error: %v", err)
	}

	_, err = client.callAPI(req)
	if err == nil {
		t.Error("expected timeout error from slow server, got nil")
	}
}
