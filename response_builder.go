package validate

type response struct {
	fields map[string]string
}

func newResponse() *response {
	return &response{fields: make(map[string]string)}
}

func (r *response) set(key string, value string) *response {
	r.fields[key] = value
	return r
}

func (r *response) build() map[string]string {
	return r.fields
}
