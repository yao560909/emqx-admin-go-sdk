package alarms

type Alarm struct {
	ActivateAt string             `json:"activate_at"`
	Details    map[string]float64 `json:"details"`
	Duration   int64              `json:"duration"`
	Message    string             `json:"message"`
	Name       string             `json:"name"`
	Node       string             `json:"node"`
}

type Meta struct {
	Count   int  `json:"count"`
	Hasnext bool `json:"hasnext"`
	Limit   int  `json:"limit"`
	Page    int  `json:"page"`
}
