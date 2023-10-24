package o2client_runtime

import (
	"fmt"
	"testing"
)

type Param struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func TestAddObjectParam(t *testing.T) {
	w, _ := NewWrapUrl("http://xx-service:8080")
	param := &struct {
		Name string `json:"name"`
		Id   int64  `json:"id"`
	}{
		Name: "Joe",
		Id:   101,
	}

	err := w.AddSimpleParamObject(param)
	fmt.Println(err)
	fmt.Println(w.ToString())
}
