package o2client_runtime

import "net/url"

func NewWrapUrl(urlPath string) (WrapUrl, error) {
	url, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	return &wrapUrl{
		url:   url,
		query: url.Query(),
	}, nil
}

type WrapUrl interface {
	AddParam(name string, value string)
	AddParamWithValues(name string, values []string)
	ToString() string
}

type wrapUrl struct {
	url   *url.URL
	query url.Values
}

func (w *wrapUrl) AddParam(name string, value string) {
	w.query.Set(name, value)
}

func (w *wrapUrl) AddParamWithValues(name string, values []string) {
	for _, v := range values {
		w.query.Add(name, v)
	}
}

func (w *wrapUrl) ToString() string {
	w.url.RawQuery = w.query.Encode()
	return w.url.String()
}
