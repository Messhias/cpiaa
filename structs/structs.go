package structs

import "time"

type Commander interface {
	Ping(host string) (PingResult, error)
	GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
	Successful bool          `json:"successfull"`
	Time       time.Duration `json:"time"`
}

type SystemInfo struct {
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ip_address"`
}

type CommandRequest struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

type CommandResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type commander struct{}
