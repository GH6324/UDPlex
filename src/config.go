package main

// Config represents the top-level configuration structure
type Config struct {
	BufferSize        int                           `json:"buffer_size"`
	BufferOffset      int                           `json:"buffer_offset"`
	QueueSize         int                           `json:"queue_size"`
	WorkerCount       int                           `json:"worker_count"`
	Services          []map[string]any              `json:"services"`
	ProtocolDetectors map[string]ProtocolDefinition `json:"protocol_detectors"`
	Logging           LoggingConfig                 `json:"logging"`
	API               APIConfig                     `json:"api"`
}

// ComponentConfig represents the common configuration for all components
type ComponentConfig struct {
	Type                string      `json:"type"`
	Tag                 string      `json:"tag"`
	ListenAddr          string      `json:"listen_addr"`
	Timeout             int         `json:"timeout"`
	ReplaceOldMapping   bool        `json:"replace_old_mapping"`
	Forwarders          []string    `json:"forwarders"`
	ReconnectInterval   int         `json:"reconnect_interval"`
	ConnectionCheckTime int         `json:"connection_check_time"`
	Detour              []string    `json:"detour"`
	SendKeepalive       *bool       `json:"send_keepalive"`
	Auth                *AuthConfig `json:"auth,omitempty"`
	BroadcastMode       *bool       `json:"broadcast_mode"`       // When false, only send to the specific connection ID
	ConnectionPoolSize  int         `json:"connection_pool_size"` // Number of connections in the pool
	NoDelay             *bool       `json:"no_delay"`
	SendTimeout         int         `json:"send_timeout"`     // ms
	RecvBufferSize      int         `json:"recv_buffer_size"` // UDP socket receive buffer size in bytes
	SendBufferSize      int         `json:"send_buffer_size"` // UDP socket send buffer size in bytes
}

// AuthConfig represents authentication and encryption settings
type AuthConfig struct {
	Enabled           bool   `json:"enabled"`
	Secret            string `json:"secret"`
	EnableEncryption  bool   `json:"enable_encryption"`
	HeartbeatInterval int    `json:"heartbeat_interval"` // seconds
	AuthTimeout       int    `json:"auth_timeout"`       // seconds
}

// FilterComponentConfig represents the configuration for a filter component
type FilterComponentConfig struct {
	Type              string              `json:"type"`
	Tag               string              `json:"tag"`
	Detour            map[string][]string `json:"detour"`
	DetourMiss        []string            `json:"detour_miss"`
	UseProtoDetectors []string            `json:"use_proto_detectors"`
}

// LoggingConfig holds all logging-related configuration
type LoggingConfig struct {
	Level      string `json:"level"`       // debug, info, warn, error, dpanic, panic, fatal
	Format     string `json:"format"`      // json or console
	OutputPath string `json:"output_path"` // file path or "stdout"
	Caller     bool   `json:"caller"`      // include caller information
}

// LoadBalancerDetourRule represents a single detour rule for load balancer
type LoadBalancerDetourRule struct {
	Rule    string   `json:"rule"`    // Expression rule for matching
	Targets []string `json:"targets"` // Target component tags (array)
}

// LoadBalancerComponentConfig represents the configuration for a load balancer component
type LoadBalancerComponentConfig struct {
	Type        string                   `json:"type"`
	Tag         string                   `json:"tag"`
	Detour      []LoadBalancerDetourRule `json:"detour"`
	WindowSize  uint32                   `json:"window_size"`
	EnableCache bool                     `json:"enable_cache"`
}
