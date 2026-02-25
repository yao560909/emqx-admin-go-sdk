package authorization

type Rule struct {
	Action     string `json:"action"`
	Permission string `json:"permission"`
	Topic      string `json:"topic"`
	Qos        []int  `json:"qos"`
	Retain     string `json:"retain"`
}

type UserRules struct {
	Username string  `json:"username"`
	Rules    []*Rule `json:"rules"`
}

type ClientRules struct {
	ClientId string  `json:"clientid"`
	Rules    []*Rule `json:"rules"`
}

type Meta struct {
	Count   int  `json:"count"`
	HasNext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}

type AuthorizationSourceStatusMetrics struct {
	Allow      int `json:"allow"`
	Deny       int `json:"deny"`
	Nomatch    int `json:"nomatch"`
	Rate       int `json:"rate"`
	RateLast5m int `json:"rate_last5m"`
	RateMax    int `json:"rate_max"`
	Total      int `json:"total"`
}

type NodeError struct {
	Node  string `json:"node"`
	Error string `json:"error"`
}

type NodeMetrics struct {
	Node    string                            `json:"node"`
	Metrics *AuthorizationSourceStatusMetrics `json:"metrics"`
}

type NodeResourceMetrics struct {
	Node    string           `json:"node"`
	Metrics *ResourceMetrics `json:"metrics"`
}

type ResourceMetrics struct {
	Success    int `json:"success"`
	Matched    int `json:"matched"`
	Failed     int `json:"failed"`
	Rate       int `json:"rate"`
	RateLast5m int `json:"rate_last5m"`
	RateMax    int `json:"rate_max"`
}

type NodeStatus struct {
	Node   string `json:"node"`
	Status string `json:"status"`
}

type AuthorizationSourceStatus struct {
	Status              string                            `json:"status"`
	Metrics             *AuthorizationSourceStatusMetrics `json:"metrics"`
	NodeError           []*NodeError                      `json:"node_error"`
	NodeMetrics         []*NodeMetrics                    `json:"node_metrics"`
	NodeResourceMetrics []*NodeResourceMetrics            `json:"node_resource_metrics"`
	NodeStatus          []*NodeStatus                     `json:"node_status"`
	ResourceMetrics     *ResourceMetrics                  `json:"resource_metrics"`
}

type AuthorizationSourceListItem struct {
	Type   string `json:"type"`
	Enable bool   `json:"enable"`
}
