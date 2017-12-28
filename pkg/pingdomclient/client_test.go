package pingdomclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCheck(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request received on: %s", r.RequestURI)
		w.WriteHeader(http.StatusOK)
	}))
	client := NewClient("testing", "tester")
	client.setBaseURL(httpServer.URL)
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.CreateCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestCreateCheckFail(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request received on: %s", r.RequestURI)
		w.WriteHeader(http.StatusAccepted)
	}))
	client := NewClient("testing", "tester")
	client.setBaseURL(httpServer.URL)
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.CreateCheck(check)

	if err.Error() != "Unable to create the check against the Pingdom API" {
		t.Fail()
		t.Error(err)
	}
}

func TestUpdateCheck(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request received on: %s", r.RequestURI)
		w.WriteHeader(http.StatusOK)
	}))
	client := NewClient("testing", "tester")
	client.setBaseURL(httpServer.URL)
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.UpdateCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestDeleteCheck(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Request received on: %s", r.RequestURI)
		w.WriteHeader(http.StatusOK)
	}))
	client := NewClient("testing", "tester")
	client.setBaseURL(httpServer.URL)
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.DeleteCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}

func TestSetBaseURL(t *testing.T) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	client := NewClient("testing", "tester")

	if client.apiBase != pingdomBaseAPI {
		t.Error()
	}

	client.setBaseURL(httpServer.URL)

	if client.apiBase != httpServer.URL {
		t.Logf("Unable to set the base URL for the request. Expected: %s Got: %s", httpServer.URL, client.apiBase)
		t.Fail()
	}
}
