package authentication

type Authentication struct {
	Id        string `json:"id"`
	Backend   string `json:"backend"`
	Enable    bool   `json:"enable"`
	Mechanism string `json:"mechanism"`
}

type User struct {
	UserId      string `json:"user_id"`
	IsSuperuser bool   `json:"is_superuser"`
}

type Meta struct {
	Count   int  `json:"count"`
	HasNext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}

type AuthenticatorStatus struct {
	Status              string                 `json:"status"`
	NodeStatus          []*NodeStatus          `json:"node_status"`
	Metrics             *AuthenticatorMetrics  `json:"metrics"`
	ResourceMetrics     *ResourceMetrics       `json:"resource_metrics"`
	NodeError           []*NodeError           `json:"node_error"`
	NodeMetrics         []*NodeMetrics         `json:"node_metrics"`
	NodeResourceMetrics []*NodeResourceMetrics `json:"node_resource_metrics"`
}

type NodeStatus struct {
	Node   string `json:"node"`
	Status string `json:"status"`
}

type NodeError struct {
	Node  string `json:"node"`
	Error string `json:"error"`
}

type AuthenticatorMetrics struct {
	Nomatch    int     `json:"nomatch"`
	Total      int     `json:"total"`
	Success    int     `json:"success"`
	Failed     int     `json:"failed"`
	Rate       float64 `json:"rate"`
	RateLast5m float64 `json:"rate_last5m"`
	RateMax    float64 `json:"rate_max"`
}

type ResourceMetrics struct {
	Success    int     `json:"success"`
	Matched    int     `json:"matched"`
	Failed     int     `json:"failed"`
	Rate       float64 `json:"rate"`
	RateLast5m float64 `json:"rate_last5m"`
	RateMax    float64 `json:"rate_max"`
}

type NodeMetrics struct {
	Node    string                `json:"node"`
	Metrics *AuthenticatorMetrics `json:"metrics"`
}

type NodeResourceMetrics struct {
	Node    string           `json:"node"`
	Metrics *ResourceMetrics `json:"metrics"`
}
