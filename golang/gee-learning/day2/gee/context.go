package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// for json-format-data
type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	// Request info
	Path   string
	Method string
	// Response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

/* Param Methods */
func (c *Context) PostForm(key string) string {
	// Get the form-data
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	// Get the query-paramter
	return c.Req.URL.Query().Get(key)
}

/* Response Format Methods */
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// Can also use: json.Encode(c.Writer)
	content, err := json.Marshal(obj)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
	c.Writer.Write(content)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
