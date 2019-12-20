package Controller

import "github.com/pinguo/pgo"

func init(){
	pgo.App.GetContainer().Bind(&TestController{})
}
