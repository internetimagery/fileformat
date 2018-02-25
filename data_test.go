// Testing data parsing, forming
package fileformat

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	parser := NewParser("hello darkness -my old friend...")
	fmt.Println(parser.Arg())
}
