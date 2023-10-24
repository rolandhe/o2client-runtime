package o2client_runtime

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
)

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
	AddSimpleParamObject(param any) error
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

func (w *wrapUrl) AddSimpleParamObject(simpleParam any) error {
	value := reflect.ValueOf(simpleParam)
	tp := reflect.TypeOf(simpleParam)
	if tp.Kind() == reflect.Pointer {
		tp = tp.Elem()
		value = value.Elem()
	}

	for i := 0; i < tp.NumField(); i++ {
		fValue := value.Field(i)
		field := tp.Field(i)
		name := field.Tag.Get("json")
		if name == "" {
			name = field.Name
		}
		if fValue.Kind() == reflect.String {
			w.query.Add(name, fValue.String())
			continue
		}
		if fValue.Kind() == reflect.Bool {
			w.query.Add(name, strconv.FormatBool(fValue.Bool()))
			continue
		}
		if fValue.CanInt() {
			w.query.Add(name, strconv.FormatInt(fValue.Int(), 10))
			continue
		}
		if fValue.CanUint() {
			w.query.Add(name, strconv.FormatUint(fValue.Uint(), 10))
			continue
		}

		return errors.New("get simpleParam just support int,uint and string")

	}
	return nil
}

func (w *wrapUrl) ToString() string {
	w.url.RawQuery = w.query.Encode()
	return w.url.String()
}
