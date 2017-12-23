package pingdom

// Check allows us to work with multiple structs when working with the Pingdom API
type Check interface {
	GetType() string
	getAPI() string
	SetData(map[string]interface{}) error
}

type check struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Type string `json:"type"`
	// TODO: implement the below fields
	// Paused                   bool              `json:"paused,omitempty"`
	// Resolution               int               `json:"resolution,omitempty"`
	// UserIDs                  []int             `json:"userids,omitempty"`
	// SendNotificationWhenDown int               `json:"sendnotificationwhendown,omitempty"`
	// NotifyAgainEvery         int               `json:"notifyagainevery,omitempty"`
	// NotifyWhenBackUp         bool              `json:"notifywhenbackup,omitempty"`
	// Tags                     []string          `json:"tags,omitempty"`
	// ProbeFilters             map[string]string `json:"probe_filters,omitempty"`
	// IPv6                     bool              `json:"ipv6,omitempty"`
	// ResponseTimeThreshold    int               `json:"responsetime_threshold,omitempty"`
	// IntegrationIDs           []int             `json:"integrationids,omitempty"`
	// TeamIDs                  []int             `json:"teamids,omitempty"`
}

func newCheck(name, host, checkType string) check {
	return check{
		Name: name,
		Host: host,
		Type: checkType,
	}
}

// GetType returns the type for the Check that will be sent to the Pingdom API
func (c *check) GetType() string {
	return c.Type
}
