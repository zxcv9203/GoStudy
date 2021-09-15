package comma

import (
	"strings"
)

func JoinWithCommas(parses []string) string {
	result := strings.Join(parses[:len(parses)-1], ", ")
	result += " and "
	result += parses[len(parses)-1]
	return result
}
