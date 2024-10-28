package spider

import (
	"github.com/charlesbases/lifecycle"
	"github.com/urfave/cli"

	"github.com/charlesbases/spider/internal/context"
)

// Spider .
type Spider interface {
	Run(ctx *cli.Context) error
	stop() error
}

// spider .
type spider struct {
	ctx context.Context

	lc lifecycle.Lifecycle
}

func (s *spider) Run(ctx *cli.Context) error {
	if err := s.initContext(ctx); err != nil {
		return err
	}

	s.initPlugin()

	defer s.stop()
	return s.run()
}

// initContext .
func (s *spider) initContext(ctx *cli.Context) error {
	s.ctx = context.New(ctx)
	return s.ctx.Error()
}

// initPlugin .
func (s *spider) initPlugin() {
	s.initLifecycle()
}

// initLifecycle .
func (s *spider) initLifecycle() {
	s.lc = lifecycle.New()

	s.lc.Append(lifecycle.Hook{})
}

func (s *spider) run() error {
	// TODO
}

func (s *spider) stop() error {
	// TODO
}

// New .
func New() Spider {
	return new(spider)
}
