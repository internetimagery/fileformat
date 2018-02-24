// Escape bad characters for file pathnames
package main

import (
	"fmt"
)

// https://superuser.com/questions/358855/what-characters-are-safe-in-cross-platform-file-names-for-linux-windows-and-os#358861
const BLACKLIST = "/\\<|>?*!:\"" // Illegal filename characters

// https://www.ascii.cl/htmlcodes.htm
// Escape filename using HTML compliant characters
func Escape(text string) string {
	res := ""
	for i, char := range text {
		esc := string(char)
		for _, bl := range BLACKLIST {
			if char == bl {
				esc = fmt.Sprintf("&#%d;", char)
				break
			}
		}
		// If & is followed by #, encode the &.
		// This is to encode existing codes.
		i++
		if char == '&' && i < len(text) && text[i] == '#' {
			esc = fmt.Sprintf("&#%d;", char)
		}
		res += esc
	}
	return res
}
