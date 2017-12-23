package pingdom

const (
	pingdomBaseAPI = "https://api.pingdom.com"
)

// This const block defines all the Check types that are supported by the Pingdom API
// TODO: currently only HTTP is supported, need to add support for all the others
const (
	CheckTypeHTTP       = "http"
	CheckTypeHTTPCustom = "httpcustom"
	CheckTypeTCP        = "tcp"
	CheckTypePing       = "ping"
	CheckTypeDNS        = "dns"
	CheckTypeUDP        = "udp"
	CheckTypeSMTP       = "smtp"
	CheckTypePOP3       = "pop3"
	CheckTypeIMAP       = "imap"
)

// This const block defines all the Pingdom API Endpoints supported by this package
const (
	APIv21Checks = "/api/2.1/checks"
)
