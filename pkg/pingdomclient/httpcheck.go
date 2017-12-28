package pingdomclient

import (
	"net/url"
	"strconv"
)

type httpCheck struct {
	check
	URL              string            `json:"url,omitempty"`
	Encryption       bool              `json:"encryption"`
	Port             int               `json:"port,omitempty"`
	Auth             string            `json:"auth,omitempty"`
	ShouldContain    string            `json:"shouldcontain,omitempty"`
	ShouldNotContain string            `json:"shouldnotcontain,omitempty"`
	PostData         string            `json:"postdata,omitempty"`
	RequestHeader    map[string]string `json:"requestheader,omitempty"`
}

// NewHTTPCheck creates a HTTP Check to send to the Pingdom API
func NewHTTPCheck(name, endpoint string) (Check, error) {
	ep, err := url.Parse(endpoint)

	if err != nil {
		return nil, err
	}

	hc := &httpCheck{
		check: newCheck(name, ep.Host, CheckTypeHTTP),
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

func (hc *httpCheck) getAPI() string {
	return APIv21Checks
}