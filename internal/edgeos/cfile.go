package edgeos

// CFile holds an array of file names
type CFile struct {
	*Env
	Names []string
}

// readDir returns a listing of dnsmasq blacklist configuration files
func (c *CFile) readDir(pattern string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Remove deletes a CFile array of file names
func (c *CFile) Remove() error { _ = "STUB: not implemented"; return nil }

// String implements string method
func (c *CFile) String() string { _ = "STUB: not implemented"; return "" }

// Strings returns a sorted array of strings.
func (c *CFile) Strings() []string { _ = "STUB: not implemented"; return nil }
