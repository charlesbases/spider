package context

import "github.com/urfave/cli"

// Context .
type Context interface {
	Error() error
}

// context .
type context struct {
	err error
}

func (c *context) Error() error {
	return c.err
}

// New .
func New(ctx *cli.Context) Context {
	// TODO
	// return new(context)
}
