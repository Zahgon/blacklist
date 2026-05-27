package edgeos

import (
	"sort"
)

// Objects is a struct of []*source
type Objects struct {
	*Env
	iface IFace
	src   []*source
}

func (o *Objects) addObj(c *Config, node string) { _ = "STUB: not implemented"; return }

// Files returns a list of dnsmasq conf files from all srcs
func (o *Objects) Files() *CFile { _ = "STUB: not implemented"; return nil }

// Filter returns a subset of Objects filtered by ltype
func (o *Objects) Filter(ltype string) *Objects { _ = "STUB: not implemented"; return nil }

// Find returns the int position of an Objects' element
func (o *Objects) Find(elem string) int { _ = "STUB: not implemented"; return 0 }

func getLtypeDesc(s string) string { _ = "STUB: not implemented"; return "" }

func (o *Objects) procltypes(c *Config, node string, ltypes ...string) {
	_ = "STUB: not implemented"
	return
}

func (o *Objects) objects(c *Config, node string, ltypes ...string) {
	_ = "STUB: not implemented"
	return
}

// Names returns a sorted slice of Objects names
func (o *Objects) Names() (s sort.StringSlice) {
	_ = "STUB: not implemented"
	return *new(sort.StringSlice)
}

// Stringer for Objects
func (o *Objects) String() (s string) { _ = "STUB: not implemented"; return "" }

// Implement Sort Interface for Objects
func (o *Objects) Len() int           { _ = "STUB: not implemented"; return 0 }
func (o *Objects) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (o *Objects) Swap(i, j int)      { _ = "STUB: not implemented"; return }
