package dashboard

type User struct {
	UserName    string `json:"username"`
	Description string `json:"description"`
	Backend     string `json:"backend"`
}

type LoginResponse struct {
	License License `json:"license"`
	Token   string  `json:"token"`
	Version string  `json:"version"`
}

type License struct {
	Edition string `json:"edition"`
}
