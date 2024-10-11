package structs

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

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

func (c *commander) Ping(host string) (PingResult, error) {
	cmd := exec.Command("ping", "-c", "1", host)
	start := time.Now()
	err := cmd.Run()
	duration := time.Since(start)
	return PingResult{Successful: err == nil, Time: duration}, err
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return SystemInfo{}, err
	}

	ip, err := GetLocalIPAddress()
	if err != nil {
		return SystemInfo{}, err
	}

	return SystemInfo{Hostname: hostname, IPAddress: ip}, nil
}

func GetLocalIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("no IP address found")
}

func NewCommander() Commander {
	return &commander{}
}
