package banned

type Banned struct {
	As     string `json:"as"`
	At     string `json:"at"`
	By     string `json:"by"`
	Reason string `json:"reason"`
	Until  string `json:"until"`
	Who    string `json:"who"`
}

type Meta struct {
	Count int `json:"count"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
