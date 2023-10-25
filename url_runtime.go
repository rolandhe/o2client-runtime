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
	w.query[name] = values
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
			name = field.Tag.Get("form")
			if name == "" {
				name = field.Name
			}
		}
		if strValue, ok := w.maybeSimple(fValue); ok {
			w.query.Add(name, strValue)
			continue
		}

		if fValue.Kind() == reflect.Slice {
			if fValue.Len() == 0 {
				continue
			}
			if err := w.acceptSlice(fValue, name); err != nil {
				return err
			}
			continue
		}

		if fValue.Kind() == reflect.Struct {
			if field.Type.Name() != name {
				return errors.New("get simpleParam just support int,uint and string, array of these simple type,but support inherit")
			}
			if err := w.AddSimpleParamObject(fValue.Interface()); err != nil {
				return err
			}
			continue
		}

		if fValue.Kind() == reflect.Pointer {
			if fValue.Elem().Kind() == reflect.Struct {
				if field.Type.Elem().Name() != name {
					return errors.New("get simpleParam just support int,uint and string, array of these simple type,but support inherit)")
				}
				if err := w.AddSimpleParamObject(fValue.Interface()); err != nil {
					return err
				}
				continue
			}
		}

		return errors.New("get simpleParam just support int,uint and string, array of these simple type,but support inherit")

	}
	return nil
}

func (w *wrapUrl) ToString() string {
	w.url.RawQuery = w.query.Encode()
	return w.url.String()
}

func (w *wrapUrl) maybeSimple(fValue reflect.Value) (string, bool) {
	if fValue.Kind() == reflect.String {
		return fValue.String(), true
	}
	if fValue.Kind() == reflect.Bool {
		return strconv.FormatBool(fValue.Bool()), true
	}
	if fValue.CanInt() {
		return strconv.FormatInt(fValue.Int(), 10), true
	}
	if fValue.CanUint() {
		return strconv.FormatUint(fValue.Uint(), 10), true
	}
	return "", false
}

func (w *wrapUrl) acceptSlice(fValue reflect.Value, name string) error {
	var values []string
	for i := 0; i < fValue.Len(); i++ {
		v := fValue.Index(i)
		if strValue, ok := w.maybeSimple(v); ok {
			values = append(values, strValue)
		} else {
			return errors.New("slice element just support int,uint and string")
		}
	}
	w.AddParamWithValues(name, values)
	return nil
}
