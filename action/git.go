package action

import (
	"os"
	"os/exec"

	"github.com/alex-my/helper-tools/utils"
	"github.com/urfave/cli/v2"
)

// GitPullAll 遍历目录下的所有 git 项目，并且更新
func GitPullAll(ctx *cli.Context) error {
	pwd, _ := os.Getwd()

	dirnames, err := utils.ListDirs(pwd, ".git", 6)
	if err != nil {
		log.Errorf("dir: %s list failed: %s", pwd, err.Error())
		return err
	}

	if len(dirnames) == 0 {
		log.Warnf("dir: %s no dirs", pwd)
	}

	bar := utils.NewBar(len(dirnames))
	bar.Run()

	gitPullAll(dirnames, bar)

	return nil
}

func gitPullAll(dirnames []string, bar *utils.Bar) {
	for _, dirname := range dirnames {
		cmd := exec.Command("git", "pull")
		cmd.Dir = dirname
		_, err := cmd.Output()

		bar.Add(dirname)

		if err != nil {
			log.Error(err.Error())
			continue
		}
	}

	bar.WaitClose()
}
