// Testing escaping

package main

import (
	"fmt"
	"html"
	"strings"
	"testing"
)

func TestEscape(t *testing.T) {
	names := []string{
		"--note this is a fine name. .jpg",
		"--fail this! is<454> fail.jpg&",
		"&#45; | or that.bmp",
		"/this/is/not/a/path.jpg",
		"--date 20180203 --event coding --notes none"}
	for _, name := range names {
		esc := Escape(name)
		if strings.ContainsAny(esc, BLACKLIST) || name != html.UnescapeString(esc) {
			fmt.Println("Old:", name)
			fmt.Println("Esc:", esc)
			fmt.Println("New:", html.UnescapeString(esc))
			t.Fail()
		}
	}
}
