package client

type Service interface {
	CURL(method string, url string, opts ...CURLOption) (interface{}, error)
	WithBody(body interface{}) CURLOption
	WithHeader(header map[string]string) CURLOption
}

type S struct{}

func New() *S {
	return &S{}
}
