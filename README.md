## helper-tools

安装

```text
git clone https://github.com/alex-my/helper-tools.git
cd helper-tools
make mod
make i
```

如果`h`或者`h.exe`没有出现在`go/bin`目录下，可以看`GOBIN`是否已有制定

```
go env
```

如果没有指定`GOBIN`，可以使用以下命令指定后再次执行`make i`

```
go env -w GOBIN="你的地址"

# 例如
go env -w GOBIN="D:\ProgramCode\go\bin"
```

如果提示`h`命令不存在，需要将`go/bin`目录添加到路径中

```
# 添加到 .bash_profile 或者 .zshrc 中
export PATH="~/go/bin:$PATH"
```

在`windows`下，以上都准备好后，执行`h pwd`会报错，可以将`h.exe`重命名为`helper`，注意重新开启终端(powershell 等)

```
helper pwd
```

## 命令列表

| 命令            | 缩写 | 作用                                                                     |
| --------------- | ---- | ------------------------------------------------------------------------ |
| PWD             | pwd  | 显示当前目录                                                             |
| GitPullAll      | gpa  | 遍历当前目录下的所有 git 项目，包括子文件夹(6 层)，分别执行 git pull     |
| GitPushAll      |      | 遍历当前目录下的所有 git 项目，包括子文件夹(6 层)，分别执行 git push     |
| GitSetRemoteUrl | gsru | 遍历当前目录下所有的 git 项目，将项目地址设置为指定地址，并执行 git push |
| SVNUpAll        | sua  | 遍历当前目录下的所有 svn 项目，包括子文件夹(6 层)，分别执行 svn up       |
| SVNCleanupAll   | sca  | 遍历当前目录下的所有 svn 项目，包括子文件夹(6 层)，分别执行 svn cleanup  |

## 使用

假设`code`文件夹下有许多的`git`项目

```text
cd code

# 在 code 下执行以下命令，会在当前目录及子目录中寻找 git 项目
h GitPullAll

# 也可以使用缩写
h gpa

# 带目标
h GitSetRemoteUrl git@xxx.git
h GitSetRemoteUrl -url=git@xxx.git
```

效果如图(里面有收集的各种语言的各种 git 项目)

![gitpullall](./images/gitpullall-process.png)
