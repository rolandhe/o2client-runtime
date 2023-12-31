package o2client_runtime

import (
	"fmt"
	"github.com/rolandhe/o2client-runtime/test"
	"testing"
)

type Param struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
}

func TestAddObjectParam(t *testing.T) {
	w, _ := NewWrapUrl("http://xx-service:8080")
	param := &struct {
		//Name  string  `json:"name"`
		Id int64 `json:"id"`
		*test.PageParam
		//Kinds []int64 `json:"kinds"`
	}{
		//Name:  "Joe",
		Id: 101,
		//Kinds: []int64{1, 2, 3},
		PageParam: &test.PageParam{
			PageNO: 101,
		},
	}

	err := w.AddSimpleParamObject(param)
	fmt.Println(err)
	fmt.Println(w.ToString())
}
