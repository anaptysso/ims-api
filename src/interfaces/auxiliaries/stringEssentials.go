package auxiliaries

// StringEssentials interface is a interface for string essentials in account logic
type StringEssentials interface {
	IsNilEmptyOrWhiteSpace(source string) bool
}
