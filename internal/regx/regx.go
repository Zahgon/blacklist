// Package regx provides regex objects for processing data in files and web content
package regx

import (
	rx "regexp"
)

// Leaf is a config label
type Leaf int

// Leaf label regx map keys
//
//ggo:generate go run golang.org/x/tools/cmd/stringer -type=Leaf
const (
	CMNT Leaf = iota + 1000
	DESC
	DSBL
	FLIP
	FQDN
	HOST
	HTTP
	IPBH
	LEAF
	LBRC
	MISC
	MLTI
	MPTY
	NAME
	NODE
	RBRC
	SUFX
)

// OBJ is a map of regex precompiled objects

type regexMap map[Leaf]*rx.Regexp

// OBJ is a struct of regex precompiled objects
type OBJ struct {
	RX regexMap
}

// NewRegex returns a map of OBJ populated with a map of precompiled regex objects
func NewRegex() *OBJ { _ = "STUB: not implemented"; return nil }

// SubMatch extracts the configuration value for a matched label
func (o *OBJ) SubMatch(t Leaf, b []byte) [][]byte { _ = "STUB: not implemented"; return nil }

func (o *OBJ) String() string { _ = "STUB: not implemented"; return "" }

// StripPrefixAndSuffix strips the prefix and suffix
func (o *OBJ) StripPrefixAndSuffix(l []byte, p string) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
