package pingdomclient

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"testing"

	"log"
)

const (
	username = "testing"
	password = "testing"
)

var exampleResponse = `{
  "check": {
    "id": 123,
    "name": "testing",
    "resolution": 5,
    "sendnotificationwhendown": 2,
    "notifyagainevery": 0,
    "notifywhenbackup": true,
    "created": 1514642336,
    "type": {
      "http": {
        "url": "/a/test",
        "encryption": true,
        "port": 443,
        "requestheaders": {
          "User-Agent": "Pingdom.com_bot_version_1.4_(http://www.pingdom.com/)"
        }
      }
    },
    "hostname": "this.is",
    "ipv6": false,
    "integrationids": [],
    "lasttesttime": 1514646668,
    "lastresponsetime": 712,
    "status": "up",
    "tags": [],
    "probe_filters": []
  }
}`

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

		if u, p, ok := r.BasicAuth(); ok {
			if u != username || p != password {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodDelete {
			checkName := r.URL.Query().Get("name")

			if checkName == "accepted" {
				w.WriteHeader(http.StatusAccepted)
			} else {
				w.WriteHeader(http.StatusOK)
			}

			w.Write([]byte(exampleResponse))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(exampleResponse))
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

func TestBadCredentials(t *testing.T) {
	client := NewClient("user", "pass")
	check, err := NewHTTPCheck("testing", "https://this.is/a/test")

	if err != nil {
		t.Error(err)
	}

	err = client.CreateCheck(check)

	if err != nil {
		if err.Error() != "Unauthorized Access make sure your credentials are correct" {
			t.Fail()
			t.Error(err)
		}
	} else {
		t.Fail()
		t.Error("Expected an error from the CreateCheck call but got nil")
	}
}

func TestCreateCheck(t *testing.T) {
	client := NewClient(username, password)
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
	client := NewClient(username, password)
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
	client := NewClient(username, password)
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
	client := NewClient(username, password)
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

func TestGetCheck(t *testing.T) {
	client := NewClient(username, password)
	check1, err := NewHTTPCheck("testing", "https://this.is/a/test")
	check1.SetID(123)

	if err != nil {
		t.Error(err)
	}

	check2, err := client.GetCheck(123)

	if err != nil {
		t.Fail()
		t.Error(err)
	}

	if !check1.Compare(check2) {
		t.Fail()
		t.Errorf("Got: %+v\n\nExpected: %+v", check2, check1)
	}
}
