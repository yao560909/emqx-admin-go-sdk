package core

type RequestOption struct {
	Header map[string]string
}

type RequestOptionFunc func(option *RequestOption)

func WithHeaders(header map[string]string) RequestOptionFunc {
	return func(option *RequestOption) {
		for k, v := range header {
			option.Header[k] = v
		}
	}
}
