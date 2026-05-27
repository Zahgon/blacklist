package edgeos

import (
	"io"
	"sort"
)

// ntype for labeling blacklist source types
type ntype int

// ntype label blacklist source types
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ntype
const (
	unknown ntype = iota // denotes a coding error
	domn                 // Format type e.g. address=/d.com/0.0.0.0
	excDomn              // Excluded from domains
	excHost              // Excluded from hosts
	excRoot              // Excluded globally
	host                 // Format type e.g. server=/www.d.com/0.0.0.0
	preDomn              // Pre-configured blacklisted domains
	preHost              // Pre-configured blacklisted hosts
	preRoot              // Pre-configured global blacklist domains
	root                 // Topmost root node
)

// booltoStr converts a boolean ("true" or "false") to a string equivalent
func booltoStr(b bool) string { _ = "STUB: not implemented"; return "" }

// diffArray returns the delta of two arrays
func diffArray(a, b []string) (diff sort.StringSlice) {
	_ = "STUB: not implemented"
	return *new(sort.StringSlice)
}

// formatData returns an io.Reader loaded with dnsmasq formatted data
func formatData(s string, l *list) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// getDnsmasqPrefix returns the dnsmasq conf file delimiter
func getDnsmasqPrefix(s *source) string { _ = "STUB: not implemented"; return "" }

// getType returns the converted "in" type
func getType(in interface{}) (out interface{}) { _ = "STUB: not implemented"; return nil }

// Iter iterates over ints - use it in for loops
func Iter(i int) []struct{} { _ = "STUB: not implemented"; return nil }

// NewWriter returns an io.Writer
func NewWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

// strToBool converts a string ("true" or "false") to boolean
func strToBool(s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func typeInt(n ntype) string { _ = "STUB: not implemented"; return "" }

func typeStr(s string) ntype { _ = "STUB: not implemented"; return *new(ntype) }
