package models

import (
	"fmt"
	"time"
)

type Result struct {
	ID                int        `json:"id" gorm:"type:integer;primary_key;AUTO_INCREMENT;unique_index"`
	User              *User      `json:"user,omitempty" gorm:"foreignkey:UserID"`
	UserID            int        `json:"-" gorm:"type:int;not null"`
	CreatedAt         *time.Time `json:"created_at"`
	Jitter            float64    `json:"jitter"`
	Latency           float64    `json:"latency"`
	DownloadBandwidth int        `json:"download_bandwidth"`
	DownloadBytes     int        `json:"download_bytes"`
	DownloadElapsed   int        `json:"download_elapsed"`
	UploadBandwidth   int        `json:"upload_bandwidth"`
	UploadBytes       int        `json:"upload_bytes"`
	UploadElapsed     int        `json:"upload_elapsed"`
	Isp               string     `json:"isp"`
	UserInternalIP    string     `json:"user_internalIp"`
	UserName          string     `json:"user_name"`
	UserMacAddr       string     `json:"user_macAddr"`
	UserIsVpn         bool       `json:"user_isVpn"`
	UserExternalIP    string     `json:"user_externalIp"`
	ServerID          int        `json:"server_id"`
	ServerName        string     `json:"server_name"`
	ServerLocation    string     `json:"server_location"`
	ServerCountry     string     `json:"server_country"`
	ServerHost        string     `json:"server_host"`
	ServerPort        int        `json:"server_port"`
	ServerIP          string     `json:"server_ip"`
	ResultID          string     `json:"result_id"`
	ResultURL         string     `json:"result_url"`
}

func (r *Result) TableName() string {
	return "results"
}

func (r *Result) ToMsg() string {
	if r == nil {
		return "Speed Test Failed. 😇"
	}
	var isVpn string
	downloadSpeed := float64(r.DownloadBandwidth) * 8 / 1000000
	uploadSpeed := float64(r.UploadBandwidth) * 8 / 1000000

	switch r.UserIsVpn {
	case true:
		isVpn = "Enable"
	case false:
		isVpn = "Disable"
	}

	return fmt.Sprintf(
		"Download/Upload: %.2f/%.2f Mbps\n"+
			"Jitter/Latency: %.2f/%.2f ms\n"+
			"ISP: %s\n"+
			"Server: %s - Location: %s/%s\n"+
			"External/Internal IP: %s/%s - VPN: %s\n"+
			"Result: %s",
		downloadSpeed, uploadSpeed,
		r.Jitter, r.Latency,
		r.Isp,
		r.ServerName, r.ServerLocation, r.ServerCountry,
		r.UserExternalIP, r.UserInternalIP, isVpn,
		r.ResultURL)
}
