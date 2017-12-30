package pingdomclient

import (
	"net/url"
	"strconv"
)

type httpCheck struct {
	check
	URL              string            `url:"url,omitempty"`
	Encryption       bool              `url:"encryption"`
	Port             int               `url:"port,omitempty"`
	Auth             string            `url:"auth,omitempty"`
	ShouldContain    string            `url:"shouldcontain,omitempty"`
	ShouldNotContain string            `url:"shouldnotcontain,omitempty"`
	PostData         string            `url:"postdata,omitempty"`
	RequestHeader    map[string]string `url:"requestheader,omitempty"`
}

type pingdomCheck struct {
	Check struct {
		ID                       int
		Name                     string
		Resolution               int
		SendNotificationWhenDown int
		NotifyAgainEvery         int
		NotifyWhenBackup         bool
		Created                  int
		Type                     struct {
			HTTP struct {
				URL            string
				Encryption     bool
				Port           int
				RequestHeaders map[string]string
			}
		}
		HostName       string
		IPv6           bool
		IntegrationIDs []int
		Status         string
		Tags           []string
		ProbeFilters   []string
	}
}

// NewHTTPCheck creates a HTTP Check to send to the Pingdom API
func NewHTTPCheck(name, endpoint string) (Check, error) {
	ep, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	hc := &httpCheck{
		check: newCheck(name, ep.Hostname(), CheckTypeHTTP),
		URL:   ep.Path,
	}

	if ep.Scheme == "https" {
		hc.Encryption = true
	}

	if ep.User != nil {
		if ep.User.Username() != "" {
			hc.Auth = ep.User.Username()
		}

		if pass, ok := ep.User.Password(); ok {
			hc.Auth = hc.Auth + ":" + pass
		}
	}

	if ep.Port() != "" {
		hc.Port, err = strconv.Atoi(ep.Port())

		if err != nil {
			return nil, err
		}
	} else if ep.Scheme == "https" {
		hc.Port = 443
	}

	return hc, nil
}

// SetData uses the data in the map[string]interface{} to set the relevant fields
// for the httpCheck
func (hc *httpCheck) SetData(data map[string]interface{}) error {
	return nil
}

// Compare takes in a Check and compares the data with itself to see if they are
// the same or not. Returns true where they have the same values
func (hc *httpCheck) Compare(check Check) bool {
	if check == nil {
		return false
	}

	ohc := check.(*httpCheck)

	if hc.Name != ohc.Name {
		return false
	}

	if hc.Host != ohc.Host {
		return false
	}

	if hc.Encryption != ohc.Encryption {
		return false
	}

	if hc.Auth != ohc.Auth {
		return false
	}

	if hc.Port != ohc.Port {
		return false
	}

	if hc.URL != ohc.URL {
		return false
	}

	return true
}

func (hc *httpCheck) getAPI() string {
	return APIv21Checks
}
