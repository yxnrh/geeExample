package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter,r *http.Request)*Context{
	return &Context{Writer: w,Req: r,Path: r.URL.Path,Method: r.Method}
}

//根据key获取表单里的value
func (c *Context)PostForm(key string)string{
	return c.Req.FormValue(key)
}

//根据key获取路由参数对应的值如果key不存在则返回""
func (c *Context)Query(key string)string{
	return c.Req.URL.Query().Get(key)
}

//把状态码写入返回的响应报文里
func (c *Context)Status(code int){
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//把数据写入响应报文头部
func (c *Context)SetHeader(key,value string){
	c.Writer.Header().Set(key,value)
}

//设置内容类型为"text/plain"
func (c *Context)String(code int,format string ,value ...interface{}){
	//把内容类型写入响应头里
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	//把内容写入响应头里
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}

func(c *Context)JSON(code int,object interface{}){
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	//对数据进行编码，如果编码失败则返回500及错误内容
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(object);err != nil{
		http.Error(c.Writer,err.Error(),500)
	}
}

//
func(c *Context)Data(code int ,data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

func(c *Context)HTML(code int ,html string){
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}