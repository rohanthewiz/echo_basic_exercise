package main

import (
	"github.com/labstack/echo"
	"bytes"
	"github.com/rohanthewiz/echo_basic_exercise/template"
	//"net/http"
)	

func main() {
	var userList = []string {
		"Alice",
		"Tiger",
		"Tom",
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) (err error) {
		buf := new(bytes.Buffer)
		template.UserList(userList, buf)
		c.HTMLBlob(200, buf.Bytes())
		
		//resp := c.Response()
		//resp.Write([]byte("Hello World"))
		//fmt.Fprint(c.Response(), "Hello World!")
		return
	})

	e.Start(":3000")
}
