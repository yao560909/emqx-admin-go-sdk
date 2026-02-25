package topics

type Topic struct {
	Topic string `json:"topic"`
	Node  string `json:"node"`
}

type Meta struct {
	Count   int  `json:"count"`
	HasNext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}
