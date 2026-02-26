package apikeys

type APIKey struct {
	APIKey  string `json:"api_key"`
	Secret  string `json:"secret"`
	Expired string `json:"expired"`
	Note    string `json:"note"`
	Tag     string `json:"tag"`
}

type Meta struct {
	Count int `json:"count"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
