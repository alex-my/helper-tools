package action

import (
	"os"
	"os/exec"

	"github.com/alex-my/helper-tools/utils"
	"github.com/urfave/cli/v2"
)

// SVNUpAll 遍历目录下的所有 svn 项目，执行 svn up
func SVNUpAll(ctx *cli.Context) error {
	return svnExecute("up")
}

// SVNCleanupAll 遍历目录下的所有 svn 项目，执行 svn cleanup
func SVNCleanupAll(ctx *cli.Context) error {
	return svnExecute("cleanup")
}

func svnExecute(command string) error {
	pwd, _ := os.Getwd()

	dirnames, err := utils.ListDirs(pwd, ".svn", 6)
	if err != nil {
		log.Errorf("dir: %s list failed: %s", pwd, err.Error())
		return err
	}

	if len(dirnames) == 0 {
		log.Warnf("dir: %s no dirs", pwd)
	}

	bar := utils.NewBar(len(dirnames))
	bar.Run()

	for _, dirname := range dirnames {
		cmd := exec.Command("svn", command)
		cmd.Dir = dirname
		_, err := cmd.Output()

		bar.Add(dirname)

		if err != nil {
			log.Error(err.Error())
			continue
		}
	}

	bar.WaitClose()

	return nil
}
