package monitor

type PrometheusConfig struct {
	Collectors      *Collectors  `json:"collectors"`
	EnableBasicAuth bool         `json:"enable_basic_auth"`
	PushGateway     *PushGateway `json:"push_gateway"`
}

type Collectors struct {
	Mnesia       string `json:"mnesia"`
	VmDist       string `json:"vm_dist"`
	VmNemory     string `json:"vm_nemory"`
	VmMsacc      string `json:"vm_msacc"`
	VmStatistics string `json:"vm_statistics"`
	VmSystemInfo string `json:"vm_system_info"`
}

type PushGateway struct {
	Enable   bool              `json:"enable"`
	Headers  map[string]string `json:"headers"`
	Interval string            `json:"interval"`
	JobName  string            `json:"job_name"`
	Url      string            `json:"url"`
}

type OpenTelemetryConfig struct {
	Exporter *OpenTelemetryExporter `json:"exporter"`
	Logs     *OpenTelemetryLogs     `json:"logs"`
	Metrics  *OpenTelemetryMetrics  `json:"metrics"`
	Traces   *OpenTelemetryTraces   `json:"traces"`
}

type OpenTelemetryExporter struct {
	Endpoint   string     `json:"endpoint"`
	SslOptions SslOptions `json:"ssl_options"`
}

type SslOptions struct {
	Ciphers              []*string `json:"ciphers"`
	Depth                int       `json:"depth"`
	Enable               bool      `json:"enable"`
	HibernateAfter       string    `json:"hibernate_after"`
	LogLevel             string    `json:"log_level"`
	Password             string    `json:"password"`
	ReuseSessions        bool      `json:"reuse_sessions"`
	SecureRenegotiate    bool      `json:"secure_renegotiate"`
	Verify               string    `json:"verify"`
	Versions             []*string `json:"versions"`
	Cacertfile           string    `json:"cacertfile"`
	Certfile             string    `json:"certfile"`
	Keyfile              string    `json:"keyfile"`
	ServerNameIndication string    `json:"server_name_indication"`
}

type OpenTelemetryLogs struct {
	Enable         bool   `json:"enable"`
	Level          string `json:"level"`
	ScheduledDelay string `json:"scheduled_delay"`
}

type OpenTelemetryMetrics struct {
	Enable   bool   `json:"enable"`
	Interval string `json:"interval"`
}

type OpenTelemetryTraces struct {
	Enable         bool   `json:"enable"`
	Filter         Filter `json:"filter"`
	ScheduledDelay string `json:"scheduled_delay"`
}

type Filter struct {
	TraceAll bool `json:"trace_all"`
}
