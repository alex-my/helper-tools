package main

import (
	"os"

	"github.com/alex-my/ghelper/logger"
	"github.com/alex-my/helper-tools/action"
	"github.com/urfave/cli/v2"
)

var log = logger.NewLogger()

func main() {
	app := &cli.App{
		Name:        "h",
		HelpName:    "helper-tools",
		Version:     "0.1.0",
		HideVersion: true,
		Commands: []*cli.Command{
			{
				Name:    "PWD",
				Aliases: []string{"pwd"},
				Usage:   "显示当前路径",
				Action:  action.PWD,
			},
			{
				Name:    "GitPullAll",
				Aliases: []string{"gpa"},
				Usage:   "遍历并更新当前目录下所有的git项目，执行 git pull",
				Action:  action.GitPullAll,
			},
			{
				Name:    "SVNUpAll",
				Aliases: []string{"sua"},
				Usage:   "遍历并更新当前目录下所有的svn项目，执行 svn up",
				Action:  action.SVNUpAll,
			},
			{
				Name:    "SVNCleanupAll",
				Aliases: []string{"sca"},
				Usage:   "遍历并更新当前目录下所有的svn项目，执行 svn cleanup",
				Action:  action.SVNCleanupAll,
			},
		},
		ExitErrHandler: ExitErrHandlerFunc,
	}

	app.Run(os.Args)
}

// ExitErrHandlerFunc ..
func ExitErrHandlerFunc(ctx *cli.Context, err error) {
	log.Errorf("%s exit: %s", ctx.Command.Name, err.Error())
}
