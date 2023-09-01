package db

type DomainScheme struct {
	Id           int64       `json:"id,omitempty"`
	Domain       string      `json:"domain"`
	Port         int         `json:"port"`
	StatusCode   int         `json:"status_code"`
	WebServer    string      `json:"webserver"`
	Host         string      `json:"host"`
	ResponseTime string      `json:"response_time"`
	Raw          interface{} `json:"raw"`
	LastUpdated  string      `json:"last_updated"`
}

type LibraryScheme struct {
	Id          int64  `json:"id,omitempty"`
	Domain      string `json:"domain"`
	LastScanned string `json:"last_scanned"`
}
