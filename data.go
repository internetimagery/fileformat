// Extract data from filename
package fileformat

import (
	"html"
	"strings"
)

// Where are we?
const (
	NONE   = iota // No data found yet
	FLAG          // -flag
	STRING        // -flag string
	ARRAY         // -flag string1 string2 etc...
)

// Pull out data from filename
// info before flags ignored
// --flag1 :: flag1 = true
// --flag2 data :: flag2 = data
// --flag3 data1 data2 data3 :: flag3 = []string{data1, data2, data3}
func pull_args(tokens []string, start int) (string, interface{}, int) {
	flag := ""
	str := ""
	array := []string{}

	last_state := NONE
	state := NONE

	i := start
	for i < len(tokens) {
		token := tokens[i]
		last_state = state
		if '-' == token[0] {
			if last_state == FLAG {
				break
			}
			state = FLAG
			flag = token
		} else if last_state == FLAG {
			state = STRING
			str = token
		} else if last_state == STRING {
			state = ARRAY
			array = []string{str, token}
		} else if last_state == ARRAY {
			array = append(array, token)
		}
		i++
	}
	var arg interface{}
	switch state {
	case FLAG:
		arg = true
	case STRING:
		arg = str
	case ARRAY:
		arg = array
	}
	return flag, arg, i
}

func Parse(name string) map[string]interface{} {
	clean := html.UnescapeString(name) // First clean up string
	tokens := strings.Split(clean, " ")
	res := make(map[string]interface{})

	i := 0
	for i < len(tokens) {
		flag, arg, i = pull_args(tokens, i)
		res[flag] = arg
	}
	return res
}
