package authentication

// Keep only common fields
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
