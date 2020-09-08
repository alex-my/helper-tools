package action

import (
	"os"
	"os/exec"

	"github.com/alex-my/ghelper/file"
	"github.com/urfave/cli/v2"
)

// GitPullAll 遍历目录下的所有 git 项目，并且更新
func GitPullAll(ctx *cli.Context) error {
	pwd, _ := os.Getwd()

	dirnames, err := file.ListDirs(pwd)
	if err != nil {
		log.Errorf("dir: %s list failed: %s", pwd, err.Error())
		return err
	}

	if len(dirnames) == 0 {
		log.Warnf("dir: %s no dirs", pwd)
	}

	for _, dirname := range dirnames {
		if file.DirContains(dirname, ".git") {
			cmd := exec.Command("git", "pull")
			cmd.Dir = dirname
			_, err := cmd.Output()
			if err != nil {
				log.Error(err.Error())
				continue
			}
			log.Infof("dir: %s done", dirname)
		}
	}

	log.Info("GitPullAll done")

	return nil
}