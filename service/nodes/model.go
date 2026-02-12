package nodes

type Node struct {
	Node            string `json:"node"`
	Connections     int64  `json:"connections"`
	LiveConnections int64  `json:"live_connections"`

	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`

	MaxFDs      int64  `json:"max_fds"`
	MemoryTotal string `json:"memory_total"`
	MemoryUsed  string `json:"memory_used"`

	NodeStatus       string `json:"node_status"`
	OTPRelease       string `json:"otp_release"`
	ProcessAvailable int64  `json:"process_available"`
	ProcessUsed      int64  `json:"process_used"`

	Uptime  int64  `json:"uptime"`
	Version string `json:"version"`
	Edition string `json:"edition"`

	SysPath string `json:"sys_path"`
	LogPath string `json:"log_path"`

	Role string `json:"role"`
}
