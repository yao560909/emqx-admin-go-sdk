package authorization

type Rule struct {
	Action    string `json:"action"`
	Permission string `json:"permission"`
	Topic     string `json:"topic"`
	Qos       []int  `json:"qos"`
	Retain    string `json:"retain"`
}

type UserRules struct {
	Username string `json:"username"`
	Rules    []*Rule `json:"rules"`
}

type Meta struct {
	Count   int  `json:"count"`
	HasNext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}
