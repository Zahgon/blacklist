package edgeos

import (
	"io"
)

// CFGcli loads configurations using the EdgeOS CFGcli
type CFGcli struct {
	*Config
	Cfg string
}

// CFGstatic loads static configurations for testing
type CFGstatic struct {
	*Config
	Cfg string
}

func active(a string, inCLI bool) string { _ = "STUB: not implemented"; return "" }

// apiCMD returns a map of CLI commands
func apiCMD(a string, inCLI bool) string { _ = "STUB: not implemented"; return "" }

// deleteFile removes a file if it exists
func deleteFile(f string) bool { _ = "STUB: not implemented"; return false }

// GetFile reads a file and returns an io.Reader
func GetFile(f string) (io.Reader, error) {
	_ = "STUB: not implemented"
	// nolint
	return *new(io.Reader), nil
}

// purgeFiles removes any orphaned blacklist files that don't have sources
func purgeFiles(files []string) error { _ = "STUB: not implemented"; return nil }

// read returns an EdgeOS API configuration io.Reader
func (c *CFGcli) read() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// read returns an EdgeOS config file io.Reader
func (c *CFGstatic) read() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// writeFile saves domains/hosts/roots data to disk
func (b *bList) writeFile() error { _ = "STUB: not implemented"; return nil }
