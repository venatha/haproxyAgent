package main

// AppConfig is basic application context requirements
type AppConfig struct {
	ThisHostname string   `json:"this_hostname"`
	AllowedHosts []string `json:"allowed_hosts"`
	LogFile      string   `json:"log_file"`
	Debug        bool     `json:"debug"`
}

// HTTPResponse struct will be used to return JSON errors
type HTTPResponse struct {
	ResponseCode int    `json:"response_code"`
	Response     string `json:"response"`
}
