package openmirror

import (
	"errors"
)

var ErrCannotPaired = errors.New("cannot paired")

// Omirror defines a openmirror data structure
type Omirror struct {
	Left  Mirror
	Right Mirror
}

// Mirror defines a sub mirror
type Mirror struct {
	paired bool
	setup  bool
}

// Openmirror defines Omirror's interface
type Openmirror interface {
	Paired() error
	Setup() bool
}

// Init inits a new omirror node
func New() *Omirror {
	om := &Omirror{}
	om.Left.paired = false
	om.Right.paired = false
	om.Left.setup = true
	om.Right.setup = true

	return om
}

// Pairing pairs two omirror node
func (o *Omirror) Paired() error {
	if o.Left.paired && o.Right.paired {
		return nil
	}
	return ErrCannotPaired
}

// Setup sets omirror nodes' set function
func (o *Omirror) Setup() bool {
	if o.Left.setup || o.Right.setup {
		return true
	}
	return false
}
