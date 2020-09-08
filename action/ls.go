// Package git 处理 git 相关
package action

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// PWD 显示当前路径
func PWD(ctx *cli.Context) error {
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	return nil
}
