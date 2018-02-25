// Testing data parsing, forming
package fileformat

import (
	"testing"
)

// NOTE: extensions (.jpg .tct etc) should be stripped before parsing.
func TestParseSimple(t *testing.T) {
	flag, data := NewParser("--taken 20170213").Next()
	if flag != "taken" || data != "20170213" {
		t.Fail()
	}
}

func TestParsePrefixed(t *testing.T) {
	flag, data := NewParser("IMG --date-captured 20170213").Next()
	if flag != "date-captured" || data != "20170213" {
		t.Fail()
	}
}
func TestParseArray(t *testing.T) {
	flag, Idata := NewParser(" --date 2017 02 13").Next()
	data := Idata.([]string)
	if flag != "date" {
		t.Fail()
	} else if data[0] != "2017" || data[1] != "02" || data[2] != "13" {
		t.Fail()
	}
}
func TestParseBool(t *testing.T) {
	flag, Idata := NewParser("--flag_1 --flag_2 data").Next()
	data := Idata.(bool)
	if flag != "flag_1" || data != true {
		t.Fail()
	}
}
func TestParseNothing(t *testing.T) {
	flag, data := NewParser(" ").Next()
	if flag != "" || data != nil {
		t.Fail()
	}
}
func TestParseRepeated(t *testing.T) {
	// Only keep first flag found
	parser := NewParser("--flag1 data --flag1 more data")
	flag, data := parser.Next()
	if flag != "flag1" || data != "data" {
		t.Fail()
	}
	flag, data = parser.Next()
	if flag != "" || data != nil {
		t.Fail()
	}
}
