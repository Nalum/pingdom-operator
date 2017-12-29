package pingdomclient

// Check allows us to work with multiple structs when working with the Pingdom API
type Check interface {
	GetType() string
	getAPI() string
	SetData(map[string]interface{}) error
	SetID(int)
	GetID() int
}

type check struct {
	Name string `url:"name"`
	Host string `url:"host"`
	Type string `url:"type"`
	ID   int    `url:"-"`
	// TODO: implement the below fields
	// Paused                   bool              `url:"paused,omitempty"`
	// Resolution               int               `url:"resolution,omitempty"`
	// UserIDs                  []int             `url:"userids,omitempty"`
	// SendNotificationWhenDown int               `url:"sendnotificationwhendown,omitempty"`
	// NotifyAgainEvery         int               `url:"notifyagainevery,omitempty"`
	// NotifyWhenBackUp         bool              `url:"notifywhenbackup,omitempty"`
	// Tags                     []string          `url:"tags,omitempty"`
	// ProbeFilters             map[string]string `url:"probe_filters,omitempty"`
	// IPv6                     bool              `url:"ipv6,omitempty"`
	// ResponseTimeThreshold    int               `url:"responsetime_threshold,omitempty"`
	// IntegrationIDs           []int             `url:"integrationids,omitempty"`
	// TeamIDs                  []int             `url:"teamids,omitempty"`
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

// GetID returns the ID of the Check in the Pingdom API
func (c *check) GetID() int {
	return c.ID
}

// GetID returns the ID of the Check in the Pingdom API
func (c *check) SetID(id int) {
	c.ID = id
}
