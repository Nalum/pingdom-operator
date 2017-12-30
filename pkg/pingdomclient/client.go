package pingdomclient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
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

// GetCheck takes an ID for a Check and retreives it from the Pingdom API
func (c *Client) GetCheck(ID int) (Check, error) {
	req, _ := http.NewRequest(
		http.MethodGet,
		pingdomBaseAPI+APIv21Checks+"/"+strconv.Itoa(ID),
		nil,
	)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.User, c.Pass)
	req.Header.Add("App-Key", c.appKey)
	resp, _ := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Unable to get HTTP Check with ID " + strconv.Itoa(ID) + " - " + resp.Status)
	}

	var checkResp pingdomCheck
	json.NewDecoder(resp.Body).Decode(&checkResp)
	var endpoint string

	if checkResp.Check.Type.HTTP.Encryption {
		endpoint = "https://"
	} else {
		endpoint = "http://"
	}

	endpoint = endpoint + checkResp.Check.HostName

	if checkResp.Check.Type.HTTP.Port > 0 {
		endpoint = endpoint + ":" + strconv.Itoa(checkResp.Check.Type.HTTP.Port)
	}

	endpoint = endpoint + checkResp.Check.Type.HTTP.URL

	check, err := NewHTTPCheck(checkResp.Check.Name, endpoint)

	if err != nil {
		return nil, err
	}

	return check, nil
}

// CreateCheck takes a Check struct and creates a new check against the
// Pingdom API
func (c *Client) CreateCheck(check Check) error {
	checkQuery, _ := query.Values(check)
	req, _ := http.NewRequest(
		http.MethodPost,
		c.apiBase+check.getAPI()+"?"+checkQuery.Encode(),
		nil,
	)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.User, c.Pass)
	req.Header.Add("App-Key", c.appKey)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("Unauthorized Access make sure your credentials are correct")
	} else if resp.StatusCode != http.StatusOK {
		rsp, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Status: %s", resp.Status)
		log.Printf("Response: %s", rsp)
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
	checkQuery, _ := query.Values(check)
	checkQuery.Del("type")
	req, _ := http.NewRequest(
		http.MethodPut,
		c.apiBase+check.getAPI()+"/"+strconv.Itoa(check.GetID())+"?"+checkQuery.Encode(),
		nil,
	)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.User, c.Pass)
	req.Header.Add("App-Key", c.appKey)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		rsp, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Status: %s", resp.Status)
		log.Printf("Response: %s", rsp)
		return errors.New("Unable to update the check against the Pingdom API")
	}

	return nil
}

// DeleteCheck takes a Check struct and deletes the matching check in the
// Pingdom API
func (c *Client) DeleteCheck(check Check) error {
	req, _ := http.NewRequest(
		http.MethodDelete,
		c.apiBase+check.getAPI()+"/"+strconv.Itoa(check.GetID()),
		nil,
	)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(c.User, c.Pass)
	req.Header.Add("App-Key", c.appKey)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		rsp, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Status: %s", resp.Status)
		log.Printf("Response: %s", rsp)
		return errors.New("Unable to delete the check against the Pingdom API")
	}

	return nil
}
