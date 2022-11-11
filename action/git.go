package action

import (
	"errors"
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

// GitPushAll 遍历目录下的所有 git 项目，并且推送
func GitPushAll(ctx *cli.Context) error {
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

	gitPushAll(dirnames, bar)

	return nil
}

func gitPushAll(dirnames []string, bar *utils.Bar) {
	for _, dirname := range dirnames {
		cmd := exec.Command("git", "push")
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

// GitSetRemoteURL 遍历当前目录下所有的 git 项目，将项目地址设置为指定地址，并执行 git push
func GitSetRemoteURL(ctx *cli.Context) error {
	// 获取目标地址
	targetURL := ctx.String("url")
	if len(targetURL) == 0 {
		targetURL = ctx.Args().Get(0)
		if len(targetURL) == 0 {
			log.Error("target url is empty")
			return errors.New("target url is empty")
		}
	}

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

	setRemoteURLAll(dirnames, bar, targetURL)

	return nil
}

func setRemoteURLAll(dirnames []string, bar *utils.Bar, targetURL string) {
	for _, dirname := range dirnames {
		cmd := exec.Command("git", []string{"remote", "set-url", "origin", targetURL}...)
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
