package pingdomclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// NewClient creates a new Pingdom Client to make requests against the Pingdom API
func NewClient(user, pass string) *Client {
	return &Client{
		apiBase: pingdomBaseAPI,
		appKey:  pingdomAppKey,
		Pass:    pass,
		User:    user,
	}
}

// Client handles sending requests to the Pingdom API
type Client struct {
	apiBase string
	appKey  string
	Pass    string
	User    string
}

// CreateCheck takes a Check struct and creates a new check against the
// Pingdom API
func (c *Client) CreateCheck(check Check) error {
	checkBytes, err := json.Marshal(check)

	if err != nil {
		return err
	}

	reader := bytes.NewReader(checkBytes)
	resp, err := http.Post(c.apiBase+check.getAPI(), "application/json", reader)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unable to create the check against the Pingdom API")
	}

	return nil
}

// UpdateCheck takes a Check struct and updates the matching check in the
// Pingdom API
func (c *Client) UpdateCheck(check Check) error {
	log.Println(c.apiBase + check.getAPI())
	return nil
}

// DeleteCheck takes a Check struct and deletes the matching check in the
// Pingdom API
func (c *Client) DeleteCheck(check Check) error {
	log.Println(c.apiBase + check.getAPI())
	return nil
}

func (c *Client) setBaseURL(url string) {
	c.apiBase = url
}
