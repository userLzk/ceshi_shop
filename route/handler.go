package route

import (
	"fmt"

	"ceshi_shop/context"
)
type Handlers struct {
	Handler map[string]HandlerFunc
}

func newHandler()*Handlers  {
	return &Handlers{Handler:make(map[string]HandlerFunc)}
}

//服务注册路由
func (H *Handlers) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	H.Handler[key] = handler
}


func (H *Handlers) setHandler(Ctx *context.Context) {
	//获取路由
	key := Ctx.Method + "-"+  Ctx.Path
	fmt.Println("key:",key)
	if handler, ok := H.Handler[key]; ok {
		handler(Ctx)
	} else {
		fmt.Sprintf("在这我有问题了！！")
	}
}