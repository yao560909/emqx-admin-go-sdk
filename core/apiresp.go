package core

type APIResp struct {
	StatusCode int               `json:"-"`
	Header     map[string]string `json:"-"`
	RawBody    []byte            `json:"-"`
}

type CodeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
