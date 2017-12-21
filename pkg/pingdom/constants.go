package pingdom

// This const block defines all the Check types that are supported by the Pingdom API
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
