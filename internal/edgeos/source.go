package edgeos

import (
	"io"
)

// source struct for normalizing EdgeOS data.
type source struct {
	*Env
	Objects
	desc     string
	disabled bool
	err      error
	exc      []string
	file     string
	inc      []string
	ip       string
	iface    IFace
	ltype    string
	nType    ntype
	name     string
	prefix   string
	r        io.Reader
	url      string
}

func (s *source) addSource(srcName [][]byte, n string) { _ = "STUB: not implemented"; return }

func (s *source) area() string { _ = "STUB: not implemented"; return "" }

// excludes returns an io.Reader of blacklist includes
func (s *source) excludes() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (s *source) filename(area string) string { _ = "STUB: not implemented"; return "" }

// includes returns an io.Reader of blacklist includes
func (s *source) includes() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func isntSource(nx []string) bool { _ = "STUB: not implemented"; return false }

func newSource() *source { _ = "STUB: not implemented"; return nil }

func (s *source) setFilePrefix(format string) string { _ = "STUB: not implemented"; return "" }

func printArray(a []string) (s string) { _ = "STUB: not implemented"; return "" }

func pad(s string) string { _ = "STUB: not implemented"; return "" }

// Process extracts hosts/domains from downloaded raw content
func (s *source) process() *bList { _ = "STUB: not implemented"; return nil }

// if b != nil {

// }

// Stringer for *source
func (s *source) String() string { _ = "STUB: not implemented"; return "" }

func (s *source) sum(area string, dropped, extracted, kept int) {
	_ = "STUB: not implemented"
	// Let's do some accounting
	return
}
