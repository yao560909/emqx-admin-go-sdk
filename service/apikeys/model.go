package apikeys

type APIKey struct {
	Name      string `json:"name"`
	ApiKey    string `json:"api_key"`
	ExpiredAt string `json:"expired_at"`
	CreatedAt string `json:"created_at"`
	Desc      string `json:"desc"`
	Enable    bool   `json:"enable"`
	Expired   bool   `json:"expired"`
	ApiSecret string `json:"api_secret"`
}
