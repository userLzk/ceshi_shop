package context

import (
	"encoding/json"
	"net/http"
)

type ReturnStruct  map[string]interface{}
//定义结构体
type Context struct {
	//writer
	Writer http.ResponseWriter
	//request 请求信息
	Request *http.Request
	//path 请求协议地址
	Path string
	//请求方式
	Method string
	//状态码
	StatusCode int

}

//初始化结构体函数
func NewContext(w http.ResponseWriter,r *http.Request)*Context  {
	//设置结构体参数值
	return &Context{
		Writer: w,
		Request: r,
		Path: r.URL.Path,
		Method: r.Method,
	}
}
func (Ctx *Context) PostForm(key string) string {
	return Ctx.Request.FormValue(key)
}


func (Ctx *Context) Query(key string) string {
	return Ctx.Request.URL.Query().Get(key)
}
//设置status 状态码
func (Ctx *Context) SetResponseStatus(code int)  {
	Ctx.StatusCode = code
	Ctx.Writer.WriteHeader(code)
}

//设置 header头
func (Ctx *Context) SetHeader(key string, value string) {
	Ctx.Writer.Header().Set(key, value)
}

//设置json返回值
func (Ctx *Context)ReturnJson(code int,obj interface{})  {
	Ctx.SetHeader("content-Type","application/json")
	Ctx.SetResponseStatus(code)
	encoder := json.NewEncoder(Ctx.Writer)
	if err:= encoder.Encode(obj);err !=nil{
		http.Error(Ctx.Writer,err.Error(),http.StatusInternalServerError)
	}
}
