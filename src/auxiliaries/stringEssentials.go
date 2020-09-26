package auxiliaries

import "strings"

// StringEssentials works string
type StringEssentials struct {
}

// IsNilEmptyOrWhiteSpace returns true if the providered string in Nil or Empty or Only Whitespace
func (se *StringEssentials) IsNilEmptyOrWhiteSpace(source string) bool {
	return len(strings.TrimSpace(source)) == 0
}
