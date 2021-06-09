package chat_subsystem

import (
	"testing"
)

func TestUnderscore(t *testing.T) {
	var pairs = []struct {
		k string
		v string
	}{
		{"ILoveGoAndJSONSoMuch", "i_love_go_and_json_so_much"},
		{"CamelCase", "camel_case"},
		{"Camel", "camel"},
		{"CAMEL", "camel"},
		{"camel", "camel"},
		{"BIGCase", "big_case"},
	}
	
	for _, tt := range pairs {
        actual := Underscore(tt.k)
        if actual != tt.v {
            t.Errorf("Underscore(%s) = %s; expected %s", tt.k, actual, tt.v)
        }
    }
}