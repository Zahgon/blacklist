package main

import (
	e "github.com/britannic/blacklist/internal/edgeos"

	"github.com/britannic/mflag"
)

// opts struct for command line options and setting initial variables
type opts struct {
	*mflag.FlagSet
	ARCH    *string
	Dbug    *bool
	DNSdir  *string
	DNStmp  *string
	File    *string
	Help    *bool
	MIPSLE  *string
	MIPS64  *string
	OS      *string
	Safe    *bool
	Test    *bool
	Verb    *bool
	Version *bool
}

// cleanArgs removes flags when code is being tested
func cleanArgs(args []string) (r []string) { _ = "STUB: not implemented"; return nil }

// getCFG returns a e.ConfLoader
func (o *opts) getCFG(c *e.Config) e.ConfLoader {
	_ = "STUB: not implemented"
	return *new(e.ConfLoader)
}

// getOpts returns command line flags and values or displays help
func getOpts() *opts { _ = "STUB: not implemented"; return nil }

func (o *opts) initEdgeOS() *e.Config { _ = "STUB: not implemented"; return nil }

// setArgs retrieves arguments entered on the command line
func (o *opts) setArgs() { _ = "STUB: not implemented"; return }

// setDir sets the directory according to the host CPU arch
func (o *opts) setDir(arch string) string { _ = "STUB: not implemented"; return "" }
