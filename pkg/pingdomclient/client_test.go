package pingdomclient

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"testing"

	"log"
)

type RewriteTransport struct {
	Transport http.RoundTripper
	URL       *url.URL
}

func (t RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// note that url.URL.ResolveReference doesn't work here
	// since t.u is an absolute url
	req.URL.Scheme = t.URL.Scheme
	req.URL.Host = t.URL.Host
	req.URL.Path = path.Join(t.URL.Path, req.URL.Path)
	rt := t.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(req)
}

func TestMain(m *testing.M) {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s Request received on: %s%s", r.Method, r.Host, r.RequestURI)

		if r.Method != http.MethodDelete {
			jsonBody := map[string]interface{}{}
			json.NewDecoder(r.Body).Decode(&jsonBody)

			if jsonBody["name"] == "accepted" {
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusOK)
			}

			w.Write([]byte(`{"check": {"id": 123, "name": "` + jsonBody["name"].(string) + `"}}`))
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer httpServer.Close()
	serverURL, err := url.Parse(httpServer.URL)

	if err != nil {
		log.Fatalf("failed to parse httptest.Server URL: %s", err)
	}

	defaultTransport := http.DefaultClient.Transport
	http.DefaultClient.Transport = RewriteTransport{URL: serverURL}
	retCode := m.Run()
	http.DefaultClient.Transport = defaultTransport
	os.Exit(retCode)
}

func TestCreateCheck(t *testing.T) {
	client := NewClient("testing", "tester")
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.CreateCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}

	if check.GetID() != 123 {
		t.Fail()
	}
}

func TestCreateCheckFail(t *testing.T) {
	client := NewClient("testing", "tester")
	check, err := NewHTTPCheck("accepted", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.CreateCheck(check)

	if err != nil {
		if err.Error() != "Unable to create the check against the Pingdom API" {
			t.Fail()
			t.Error(err)
		}
	} else {
		t.Fail()
		t.Error("Expected an error from the CreateCheck call but got nil")
	}
}

func TestUpdateCheck(t *testing.T) {
	client := NewClient("testing", "tester")
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")
	check.SetID(123)

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
	client := NewClient("testing", "tester")
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")
	check.SetID(123)

	if err != nil {
		t.Error(err)
	}

	err = client.DeleteCheck(check)

	if err != nil {
		t.Fail()
		t.Error(err)
	}
}
