// Extract data from filename
package fileformat

import (
	"html"
	"strings"

	"github.com/flynn-archive/go-shlex"
)

// Where are we?
const (
	NONE   = iota // No data found yet
	FLAG          // -flag
	STRING        // -flag string
	ARRAY         // -flag string1 string2 etc...
)

type Parser struct {
	pos    int
	tokens []string
	seen   map[string]bool
}

func NewParser(name string) *Parser {
	clean := html.UnescapeString(name) // First clean up string
	tokens, err := shlex.Split(clean)
	if err != nil {
		panic(err)
	}
	parser := new(Parser)
	parser.tokens = tokens
	parser.seen = make(map[string]bool)
	return parser
}

// Pull out data from filename
// info before flags ignored
// --flag1 :: flag1 = true
// --flag2 data :: flag2 = data
// --flag3 data1 data2 data3 :: flag3 = []string{data1, data2, data3}
func (self *Parser) Next() (string, interface{}) {
	flag := ""
	str := ""
	array := []string{}

	state := NONE

	for self.pos < len(self.tokens) {
		token := self.tokens[self.pos]
		if '-' == token[0] {
			if state != NONE {
				break
			}
			flag = strings.TrimLeft(token, "-") // Strip leading dashes
			state = FLAG
		} else if state == FLAG {
			state = STRING
			str = token
		} else if state == STRING {
			state = ARRAY
			array = []string{str, token}
		} else if state == ARRAY {
			array = append(array, token)
		}
		self.pos++
	}
	var arg interface{}
	switch state {
	case FLAG:
		arg = true
	case STRING:
		arg = str
	case ARRAY:
		arg = array
	default:
		arg = nil
	}
	if self.seen[flag] {
		return "", nil
	}
	self.seen[flag] = true
	return flag, arg
}
