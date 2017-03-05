// Code generated by hero.
// source: /home/ro/go/src/github.com/rohanthewiz/echo_basic_exercise/template/userlist.html
// DO NOT EDIT!
package template

import "github.com/shiyanhui/hero"
import "bytes"

func UserList(userList []string, buffer *bytes.Buffer) {
	buffer.WriteString(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
  </head>

  <body>
  	<nav><li>One</li><li>Two</li><li>Three</li></nav>
  	`)
	for _, user := range userList {
		buffer.WriteString(`
		<ul>
			`)
		buffer.WriteString(`<li>
	`)
		hero.EscapeHTML(user, buffer)
		buffer.WriteString(`
</li>
`)

		buffer.WriteString(`
		</ul>
		
	`)
	}

	buffer.WriteString(`
  </body>
</html>
	`)

}
