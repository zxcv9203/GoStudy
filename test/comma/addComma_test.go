package comma

import (
	"testing"
)

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		testData{list: []string{"hello"}, want: "hello"},
		testData{list: []string{"hello", "world"}, want: "hello and world"},
		testData{list: []string{"hello", "world", "go"}, want: "hello, world and go"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}
