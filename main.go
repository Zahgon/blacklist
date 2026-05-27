package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"time"

	e "github.com/britannic/blacklist/internal/edgeos"
)

var (
	// updated by go build -ldflags
	architecture = "UNKNOWN"
	build        = "UNKNOWN"
	githash      = "UNKNOWN"
	hostOS       = "UNKNOWN"
	version      = "UNKNOWN"
	// ----------------------------

	exitCmd      = os.Exit
	initEnvirons = initEnv
	prog         = basename(os.Args[0])
	prefix       = fmt.Sprintf("%s: ", prog)
	bkpCfgFile   = "/config/user-data/blacklist.failover.cfg"
	stdCfgFile   = "/config/config.boot"
)

// Hack to reduce memory usage in Go 1.17
func init() {
	go func() {
		t := time.Tick(time.Second)
		for {
			<-t
			debug.FreeOSMemory()
		}
	}()
}

func main() {
	// Memory profiling
	// defer profile.Start(profile.MemProfile).Stop()

	objex := []e.IFace{
		e.PreRObj,
		e.PreDObj,
		e.PreHObj,
		e.ExRtObj,
		e.ExDmObj,
		e.ExHtObj,
		e.FileObj,
		e.URLdObj,
		e.URLhObj,
	}

	if os.Geteuid() != 0 {
		fmt.Printf("%s must be run as sudo\n", prog)
		logErrorf("%s must be run as sudo", prog)
		exitCmd(0)
	}
	c, err := initEnvirons()
	if err != nil {
		logErrorf("Cannot continue due to error: %s", err.Error())
		exitCmd(0)
	}

	c.Debug(fmt.Sprintf("Dumping commandline args: %v", os.Args[1:]))
	c.Debug(fmt.Sprintf("Dumping env variables: %v", c))
	logNoticef("%v", "Starting blacklist update...")

	if !e.ChkWeb("www.google.com", 443) {
		logFatalf("%s", "No internet access, aborting blacklist update!")
	}

	logInfo("Checking for stale blacklists...")
	if err = removeStaleFiles(c); err != nil {
		logFatalf("%v", err.Error())
	}

	// _, _ = context.WithTimeout(context.Background(), c.Timeout)

	if !c.Disabled {
		if err := processObjects(c, objex); err != nil {
			logErrorf("%v", err.Error())
		}
	}

	dropped, extracted, kept := c.GetTotalStats()
	if kept+dropped != 0 {
		c.Log.Noticef("Total entries found: %d", extracted)
		c.Log.Noticef("Total entries extracted %d", kept)
		c.Log.Noticef("Total entries dropped %d", dropped)
	}

	reloadDNS(c)

	logNoticef("%v", "Blacklist update completed......")
}

// basename removes directory components and file extensions.
func basename(s string) string {
	_ = "STUB: not implemented"
	// Discard last '/' and everything before.
	return ""
}

// Preserve everything before last '.'

// files returns an empty *e.CFile string array
func files(c *e.Config) *e.CFile { _ = "STUB: not implemented"; return nil }

func initEnv() (c *e.Config, err error) { _ = "STUB: not implemented"; return nil, nil }

func loadConfig(c *e.Config, o *opts) (*e.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processObjects processes local sources, downloads Internet sources and creates
// dnsmasq configuration files
func processObjects(c *e.Config, objects []e.IFace) error { _ = "STUB: not implemented"; return nil }

// reloadDNS reloads the latest processed dnsmasq configuration files
func reloadDNS(c *e.Config) { _ = "STUB: not implemented"; return }

// removeStaleFiles deletes redundant files
func removeStaleFiles(c *e.Config) error { _ = "STUB: not implemented"; return nil }
