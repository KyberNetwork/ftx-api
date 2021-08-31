package ftxapi

import "strings"

type request struct {
	httpMethod string
	endpoint   string
	needSigned bool
	params     map[string]string
	body       []byte
}

func newRequest(httpMethod, endpoint string, needSigned bool) *request {
	return &request{
		httpMethod: httpMethod,
		endpoint:   strings.TrimPrefix(endpoint, "/"),
		needSigned: needSigned,
		params:     make(map[string]string),
	}
}

func (r *request) setParam(key, value string) {
	r.params[key] = value
}

func (r *request) setBody(body []byte) {
	r.body = body
}
