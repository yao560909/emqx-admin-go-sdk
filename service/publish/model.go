package publish

type PublishResponse struct {
	Id         string `json:"id"`
	ReasonCode int    `json:"reason_code"`
	Message    string `json:"message"`
}
