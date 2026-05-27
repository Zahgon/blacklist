package edgeos

import (
	"time"

	logging "github.com/britannic/go-logging"
)

// Env is struct of parameters
type Env struct {
	ctr
	// ioWriter io.Writer
	Log      *logging.Logger
	API      string        `json:"API,omitempty"`
	Arch     string        `json:"Arch,omitempty"`
	Bash     string        `json:"Bash,omitempty"`
	Cores    int           `json:"Cores,omitempty"`
	Disabled bool          `json:"Disabled"`
	Dbug     bool          `json:"Dbug,omitempty"`
	Dex      *list         `json:"Dex,omitempty"`
	Dir      string        `json:"Dir,omitempty"`
	DNSsvc   string        `json:"dnsmasq service,omitempty"`
	Exc      *list         `json:"Exc,omitempty"`
	Ext      string        `json:"dnsmasq fileExt.,omitempty"`
	File     string        `json:"File,omitempty"`
	FnFmt    string        `json:"File name fmt,omitempty"`
	InCLI    string        `json:"-"`
	Method   string        `json:"HTTP method,omitempty"`
	Pfx      dnsPfx        `json:"Prefix,omitempty"`
	Test     bool          `json:"Test,omitempty"`
	Timeout  time.Duration `json:"Timeout,omitempty"`
	Verb     bool          `json:"Verbosity,omitempty"`
	Wildcard/*..........*/ `json:"Wildcard,omitempty"`
}

// dnsPfx defines the prefix entries in the dnsmasq configuration file
type dnsPfx struct {
	domain string
	host   string
}

// Wildcard struct sets globbing wildcards for filename searches
type Wildcard struct {
	Node string `json:"Node,omitempty"`
	Name string `json:"Name,omitempty"`
}

// Debug logs debug messages when the Dbug flag is true
func (e *Env) Debug(s ...interface{}) { _ = "STUB: not implemented"; return }

// Option is a recursive function
type Option func(c *Config) Option

// SetOpt sets the specified options passed as Env and returns an option to restore the last set of arg's previous values
func (c *Config) SetOpt(opts ...Option) Option {
	_ = "STUB: not implemented"
	// apply all the options, and replace each with its inverse
	return *new(Option)
}

// Arch sets target CPU architecture
func Arch(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// API sets the EdgeOS CLI API command
func API(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Bash sets the shell processor
func Bash(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Cores sets max CPU cores
func Cores(i int) Option { _ = "STUB: not implemented"; return *new(Option) }

// Disabled toggles Disabled
func Disabled(b bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Dbug toggles Debug level on or off
func Dbug(b bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Dir sets directory location
func Dir(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// DNSsvc sets dnsmasq restart command
func DNSsvc(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Ext sets the blacklist file n extension
func Ext(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// File sets the EdgeOS configuration file
func File(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// FileNameFmt sets the EdgeOS configuration file name format
func FileNameFmt(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// InCLI sets the CLI inSession command
func InCLI(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Logger sets a pointer to the logger
func Logger(l *logging.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

// Method sets the HTTP method
func Method(s string) Option { _ = "STUB: not implemented"; return *new(Option) }

// NewConfig returns a new *Config initialized with the parameter options passed to it
func NewConfig(opts ...Option) *Config { _ = "STUB: not implemented"; return nil }

// Prefix sets the dnsmasq configuration address line prefix
func Prefix(d string, h string) Option { _ = "STUB: not implemented"; return *new(Option) }

// Env Stringer interface
func (e *Env) String() string { _ = "STUB: not implemented"; return "" }

// Test toggles testing mode on or off
func Test(b bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// Timeout sets how long before an unresponsive goroutine is aborted
func Timeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// Verb sets the verbosity level to v
func Verb(b bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WCard sets file globbing wildcard values
func WCard(w Wildcard) Option { _ = "STUB: not implemented"; return *new(Option) }
