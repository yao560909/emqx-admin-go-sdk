package listeners

type Listener struct {
	Id         string        `json:"id"`
	Acceptors  int           `json:"acceptors"`
	Enable     bool          `json:"enable"`
	Name       string        `json:"name"`
	Type       string        `json:"type"`
	Status     *Status       `json:"status"`
	Bind       string        `json:"bind"`
	NodeStatus []*NodeStatus `json:"node_status"`
	Number     int   `json:"number"`
}

type ListenerStatus struct {
	Enable     bool          `json:"enable"`
	Ids        []string      `json:"ids"`
	NodeStatus []*NodeStatus `json:"node_status"`
	Status     *Status       `json:"status"`
	Type       string        `json:"type"`
}

type Status struct {
	CurrentConnections int    `json:"current_connections"`
	MaxConnections     string `json:"max_connections"`
	Running            bool   `json:"running"`
}

type NodeStatus struct {
	Node   string `json:"node"`
	Status *Status `json:"status"`
}
