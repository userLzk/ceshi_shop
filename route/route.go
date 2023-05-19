package route

import (
	"net/http"

	"ceshi_shop/context"
)

//定义结构体 结构体包含 writer request
type HandlerFunc func(h *context.Context)

//路由存储
type Engine struct {
	router *Handlers
}

func New() *Engine {
	return &Engine{router: newHandler()}
}

//服务注册路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlers := context.NewContext(w, r)
	engine.router.setHandler(handlers)
}

//启动服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
