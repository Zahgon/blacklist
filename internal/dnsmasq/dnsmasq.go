// Package dnsmasq parses dnsmasq.conf address and server name IP mapping files
package dnsmasq

import (
	"io"
)

const (
	address = "address="
	server  = "server="
)

// Host is a container for IP addresses
type Host struct {
	IP     string `json:"IP,omitempty"`
	Server bool   `json:"Server,omitempty"`
}

// Conf is map of Hosts
type Conf map[string]Host

type confLoader interface {
	read() io.Reader
}

// Mapping holds a dnsmasq configuration file contents
type Mapping struct {
	Contents []byte
}

// ConfigFile reads a file and returns an io.Reader
func ConfigFile(f string) (io.Reader, error) {
	_ = "STUB: not implemented"
	// nolint
	return *new(io.Reader), nil
}

func fetchHost(dns, ip string) bool { _ = "STUB: not implemented"; return false }

func ipOK(i, x string) bool { _ = "STUB: not implemented"; return false }

func matchIP(ip string, ips []string) bool { _ = "STUB: not implemented"; return false }

// Parse extracts host to IP mappings from a dnsmasq configuration file
func (c Conf) Parse(r confLoader) error { _ = "STUB: not implemented"; return nil }

func (m *Mapping) read() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

// Redirect returns true if the resolved IP address matches the correct IP (redirected or normal)
func (c Conf) Redirect(k, ip string) bool { _ = "STUB: not implemented"; return false }

func (c Conf) String() string { _ = "STUB: not implemented"; return "" }
