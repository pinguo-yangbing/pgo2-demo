package Controller

import "github.com/pinguo/pgo"

type TestController struct {

	pgo.Controller
}

func (t *TestController) ActionIndex() {
	t.OutputJson("hello world", 200, "success")
}