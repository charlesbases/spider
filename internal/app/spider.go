package app

import (
	"os"
	"path/filepath"

	"github.com/urfave/cli"

	"github.com/charlesbases/spider/internal/spider"
)

// Spider .
func Spider() *cli.App {
	return &cli.App{
		Name:      filepath.Base(os.Args[0]),
		Usage:     "video spider",
		ArgsUsage: "[link]",
		Action: func(ctx *cli.Context) error {
			return spider.New().Run(ctx)
		},
	}
}
