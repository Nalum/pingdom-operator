package pingdom

const (
	pingdomBaseAPI = "https://api.pingdom.com/"
)

// NewClient creates a new Pingdom Client to make requests against the Pingdom API
func NewClient() *Client {
	return &Client{
		apiBase: pingdomBaseAPI,
	}
}

// Client handles sending requests to the Pingdom API
type Client struct {
	apiBase string
}

// CreateCheck takes a Check struct and creates a new check against the
// Pingdom API
func (c *Client) CreateCheck(check Check) error {
	return nil
}

// UpdateCheck takes a Check struct and updates the matching check in the
// Pingdom API
func (c *Client) UpdateCheck(check Check) error {
	return nil
}

// DeleteCheck takes a Check struct and deletes the matching check in the
// Pingdom API
func (c *Client) DeleteCheck(check Check) error {
	return nil
}
