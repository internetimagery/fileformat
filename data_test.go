// Testing data parsing, forming
package fileformat

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tmp := Parse("hello darkness -my old friend...")
	fmt.Println(tmp)
}
