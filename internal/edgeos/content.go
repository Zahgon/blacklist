package edgeos

import (
	"io"
)

// IFace type for labeling interface types
type IFace int

// IFace types for labeling Content interfaces
const (
	notfound       = -1
	Invalid  IFace = iota + 100
	ExDmObj
	ExHtObj
	ExRtObj
	FileObj
	FylDObj
	FylHObj
	PreDObj
	PreHObj
	PreRObj
	URLdObj
	URLhObj
)

type bList struct {
	file string
	r    io.Reader
	size int
}

// Contenter is an interface for handling the different file/http data sources
type Contenter interface {
	Find(string) int
	GetList() *Objects
	Len() int
	SetURL(string, string)
	String() string
}

// ExcDomnObjects struct of *Objects for domain exclusions
type ExcDomnObjects struct {
	*Objects
}

// ExcHostObjects struct of *Objects for host exclusions
type ExcHostObjects struct {
	*Objects
}

// ExcRootObjects struct of *Objects for global domain exclusions
type ExcRootObjects struct {
	*Objects
}

// FIODataObjects struct of *Objects for files
type FIODataObjects struct {
	*Objects
}

// FIODomnObjects struct of *Objects for files
type FIODomnObjects struct {
	*Objects
}

// FIOHostObjects struct of *Objects for files
type FIOHostObjects struct {
	*Objects
}

// PreDomnObjects struct of *Objects for pre-configured domains content
type PreDomnObjects struct {
	*Objects
}

// PreHostObjects struct of *Objects for pre-configured hosts content
type PreHostObjects struct {
	*Objects
}

// PreRootObjects struct of *Objects for pre-configured hosts content
type PreRootObjects struct {
	*Objects
}

// URLDomnObjects struct of *Objects for domain URLs
type URLDomnObjects struct {
	*Objects
}

// URLHostObjects struct of *Objects for host URLs
type URLHostObjects struct {
	*Objects
}

// Find returns the int position of an Objects' element
func (e *ExcDomnObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (e *ExcHostObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (e *ExcRootObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (f *FIODataObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (p *PreDomnObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (p *PreHostObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (p *PreRootObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (u *URLHostObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// Find returns the int position of an Objects' element
func (u *URLDomnObjects) Find(s string) int { _ = "STUB: not implemented"; return 0 }

// GetList implements the Contenter interface for ExcDomnObjects
func (e *ExcDomnObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for ExcHostObjects
func (e *ExcHostObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for ExcRootObjects
func (e *ExcRootObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for FIODataObjects
func (f *FIODataObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for PreDomnObjects
func (p *PreDomnObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for PreHostObjects
func (p *PreHostObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for PreRootObjects
func (p *PreRootObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for URLDomnObjects
func (u *URLDomnObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// GetList implements the Contenter interface for URLHostObjects
func (u *URLHostObjects) GetList() *Objects { _ = "STUB: not implemented"; return nil }

// Len returns how many sources there are
func (e *ExcDomnObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (e *ExcHostObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (e *ExcRootObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (f *FIODataObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (p *PreDomnObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (p *PreHostObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (p *PreRootObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (u *URLDomnObjects) Len() int {
	_ = "STUB: not implemented"

	// Len returns how many sources there are
	return 0
}

func (u *URLHostObjects) Len() int {
	_ = "STUB: not implemented"

	// SetURL sets the Object's url field value
	return 0
}

func (e *ExcDomnObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (e *ExcHostObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (e *ExcRootObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (f *FIODataObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (p *PreDomnObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (p *PreHostObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (p *PreRootObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (u *URLDomnObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

// SetURL sets the Object's url field value
func (u *URLHostObjects) SetURL(name, url string) { _ = "STUB: not implemented"; return }

func (e *ExcDomnObjects) String() string { _ = "STUB: not implemented"; return "" }
func (e *ExcHostObjects) String() string { _ = "STUB: not implemented"; return "" }
func (e *ExcRootObjects) String() string { _ = "STUB: not implemented"; return "" }
func (f *FIODataObjects) String() string { _ = "STUB: not implemented"; return "" }
func (p *PreDomnObjects) String() string { _ = "STUB: not implemented"; return "" }
func (p *PreHostObjects) String() string { _ = "STUB: not implemented"; return "" }
func (p *PreRootObjects) String() string { _ = "STUB: not implemented"; return "" }
func (u *URLDomnObjects) String() string { _ = "STUB: not implemented"; return "" }
func (u *URLHostObjects) String() string { _ = "STUB: not implemented"; return "" }

func (i IFace) String() string { _ = "STUB: not implemented"; return "" }
