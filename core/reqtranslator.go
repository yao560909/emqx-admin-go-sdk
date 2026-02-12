package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"strings"
)

const (
	contentTypeHeader  = "Content-Type"
	contentTypeJson    = "application/json"
	defaultContentType = contentTypeJson + "; charset=utf-8"
)

type ReqTranslator struct {
}

type FormData struct {
	mimeType string
	fields   map[string]interface{}
	data     *struct {
		content     []byte
		contentType string
	}
}

func (fd *FormData) content() (string, []byte, error) {
	if fd.data != nil {
		return fd.data.contentType, fd.data.content, nil
	}
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, val := range fd.fields {
		if r, ok := val.(io.Reader); ok {
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition",
				fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
					escapeQuotes("file"), escapeQuotes(key)))
			h.Set("Content-Type", fd.mimeType)
			part, err := writer.CreatePart(h)
			if err != nil {
				return "", nil, err
			}
			_, err = io.Copy(part, r)
			if err != nil {
				return "", nil, err
			}
			continue
		}
		err := writer.WriteField(key, fmt.Sprint(val))
		if err != nil {
			return "", nil, err
		}
	}
	contentType := writer.FormDataContentType()
	err := writer.Close()
	if err != nil {
		return "", nil, err
	}
	fd.data = &struct {
		content     []byte
		contentType string
	}{content: buf.Bytes(), contentType: contentType}
	return fd.data.contentType, fd.data.content, nil
}

func (translator *ReqTranslator) Payload(body interface{}) (string, []byte, error) {
	if fd, ok := body.(*FormData); ok {
		return fd.content()
	}
	contentType := defaultContentType
	if body == nil {
		return contentType, nil, nil
	}
	bs, err := json.Marshal(body)
	return contentType, bs, err
}

func (translator *ReqTranslator) Path(path string, pathParams PathParams) (string, error) {
	for {
		start := strings.Index(path, "{")
		if start == -1 {
			return path, nil
		}
		end := strings.Index(path[start:], "}")
		if end == -1 {
			return "", fmt.Errorf("invalid template: %s", path)
		}
		end += start
		key := path[start+1 : end]
		v := pathParams.Get(key)
		if v == "" {
			return "", fmt.Errorf("missing path param: %s", key)
		}
		path = path[:start] + url.PathEscape(v) + path[end+1:]
	}
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
