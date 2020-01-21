package speedtest

import "time"

type Result struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Ping      struct {
		Jitter  float64 `json:"jitter"`
		Latency float64 `json:"latency"`
	} `json:"ping"`
	Download struct {
		Bandwidth int `json:"bandwidth"`
		Bytes     int `json:"bytes"`
		Elapsed   int `json:"elapsed"`
	} `json:"download"`
	Upload struct {
		Bandwidth int `json:"bandwidth"`
		Bytes     int `json:"bytes"`
		Elapsed   int `json:"elapsed"`
	} `json:"upload"`
	Isp       string `json:"isp"`
	Interface struct {
		InternalIP string `json:"internalIp"`
		Name       string `json:"name"`
		MacAddr    string `json:"macAddr"`
		IsVpn      bool   `json:"isVpn"`
		ExternalIP string `json:"externalIp"`
	} `json:"interface"`
	Server struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Location string `json:"location"`
		Country  string `json:"country"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		IP       string `json:"ip"`
	} `json:"server"`
	Result struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"result"`
}

type ServerList struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	Servers   []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Location string `json:"location"`
		Country  string `json:"country"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
	} `json:"servers"`
}