package pingdomclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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
	checkBytes, _ := json.Marshal(check)
	reader := bytes.NewReader(checkBytes)
	resp, err := http.Post(c.apiBase+check.getAPI(), "application/json", reader)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unable to create the check against the Pingdom API")
	}

	jsonBody := map[string]interface{}{}
	json.NewDecoder(resp.Body).Decode(&jsonBody)

	if checkObj, ok := jsonBody["check"].(map[string]interface{}); ok {
		if id, ok := checkObj["id"].(float64); ok {
			check.SetID(int(id))
		}
	}

	if check.GetID() == 0 {
		return errors.New("Something went wrong with the request to create the Check")
	}

	return nil
}

// UpdateCheck takes a Check struct and updates the matching check in the
// Pingdom API
func (c *Client) UpdateCheck(check Check) error {
	checkBytes, _ := json.Marshal(check)
	reader := bytes.NewReader(checkBytes)
	req, err := http.NewRequest(http.MethodPut, c.apiBase+check.getAPI()+"/"+strconv.Itoa(check.GetID()), reader)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unable to update the check against the Pingdom API")
	}

	return nil
}

// DeleteCheck takes a Check struct and deletes the matching check in the
// Pingdom API
func (c *Client) DeleteCheck(check Check) error {
	reader := bytes.NewReader([]byte{})
	req, err := http.NewRequest(http.MethodDelete, c.apiBase+check.getAPI()+"/"+strconv.Itoa(check.GetID()), reader)

	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unable to create the check against the Pingdom API")
	}

	return nil
}
