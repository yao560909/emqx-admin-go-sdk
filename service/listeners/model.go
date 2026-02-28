package listeners

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Listener struct {
	Id         string        `json:"id"`
	Acceptors  int           `json:"acceptors"`
	Enable     bool          `json:"enable"`
	Name       string        `json:"name"`
	Type       string        `json:"type"`
	Status     *Status       `json:"status"`
	Bind       string        `json:"bind"`
	NodeStatus []*NodeStatus `json:"node_status"`
	Number     int           `json:"number"`
}

type ListenerStatus struct {
	Enable     bool          `json:"enable"`
	Ids        []string      `json:"ids"`
	NodeStatus []*NodeStatus `json:"node_status"`
	Status     *Status       `json:"status"`
	Type       string        `json:"type"`
}

type Status struct {
	CurrentConnections int     `json:"current_connections"`
	MaxConnections     MaxConn `json:"max_connections"`
	Running            bool    `json:"running"`
}

type NodeStatus struct {
	Node   string  `json:"node"`
	Status *Status `json:"status"`
}

type MaxConn struct {
	value int64
	inf   bool
}

func (m *MaxConn) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == `"infinity"` {
		m.inf = true
		m.value = 0
		return nil
	}
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("invalid max_connections: %s", string(data))
	}
	m.value = v
	m.inf = false
	return nil
}

func (m MaxConn) IsInfinity() bool {
	return m.inf
}

func (m MaxConn) Int64() int64 {
	if m.inf {
		return -1
	}
	return m.value
}

func (m MaxConn) String() string {
	if m.inf {
		return "infinity"
	}
	return strconv.FormatInt(m.value, 10)
}
