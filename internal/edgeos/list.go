package edgeos

import (
	"sync"
)

type entry map[string]struct{}

// list is a struct map of entry with a RW Mutex
type list struct {
	*sync.RWMutex
	entry
}

// set sets the int value of entry
func (l *list) keyExists(k []byte) bool { _ = "STUB: not implemented"; return false }

// merge returns a merge of two lists
func (l *list) merge(a *list) { _ = "STUB: not implemented"; return }

// set adds a list entry map member
func (l *list) set(k []byte) { _ = "STUB: not implemented"; return }

func (l *list) String() string { _ = "STUB: not implemented"; return "" }

// subKeyExists returns true if part or all of the key matches
func (l *list) subKeyExists(b []byte) bool { _ = "STUB: not implemented"; return false }
