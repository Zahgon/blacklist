package main

import (
	"runtime"

	logging "github.com/britannic/go-logging"
)

var (
	boldcolors = []string{
		logging.CRITICAL: logging.ColorSeqBold(logging.ColorMagenta),
		logging.ERROR:    logging.ColorSeqBold(logging.ColorRed),
		logging.INFO:     logging.ColorSeqBold(logging.ColorGreen),
		logging.WARNING:  logging.ColorSeqBold(logging.ColorYellow),
		logging.NOTICE:   logging.ColorSeqBold(logging.ColorCyan),
		logging.DEBUG:    logging.ColorSeqBold(logging.ColorBlue),
	}
	fdFmttr    logging.Backend
	haveTerm   = inTerminal
	log        = newLog(prefix)
	logCritf   = log.Criticalf
	logErrorf  = func(f string, args ...interface{}) { log.Errorf(f, args...) }
	logFatalf  = func(f string, args ...interface{}) { logCritf(f, args...); exitCmd(1) }
	logFile    = setLogFile(runtime.GOOS)
	logInfo    = log.Info
	logInfof   = log.Infof
	logNoticef = log.Noticef
	logPrintf  = logInfof
)

// inTerminal returns true if the current terminal is interactive
func inTerminal() bool { _ = "STUB: not implemented"; return false }

// setLogFile returns a log directory and file name dependent on the current OS
func setLogFile(os string) string { _ = "STUB: not implemented"; return "" }

// newLog returns a logging.Logger pointer
func newLog(prefix string) *logging.Logger { _ = "STUB: not implemented"; return nil }

// nolint

func newScreenLogBackend(colors []string, prefix string) *logging.LogBackend {
	_ = "STUB: not implemented"
	return nil
}

// screenLog adds stderr logging output to the screen
func screenLog(prefix string) logging.LeveledBackend {
	_ = "STUB: not implemented"
	return *new(logging.LeveledBackend)
}
