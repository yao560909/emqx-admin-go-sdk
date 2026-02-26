package subscriptions

type Subscription struct {
	ClientID string `json:"clientid"`
	Nl       int    `json:"nl"`
	Node     string `json:"node"`
	Qos      int    `json:"qos"`
	Rap      int    `json:"rap"`
	rh       int    `json:"rh"`
	Topic    string `json:"topic"`
}

type Meta struct {
	Count   int  `json:"count"`
	Hasnext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}
