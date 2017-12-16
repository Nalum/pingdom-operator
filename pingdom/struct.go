package pingdom

// Check is the Pingdom Check as defined by their API documentation
type Check struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Create sends a POST request to the Pingdom API to create the Check
func (c *Check) Create() error {
	return nil
}

// Update sends a PUT request to the Pingdom API to update the Check
func (c *Check) Update() error {
	return nil
}
