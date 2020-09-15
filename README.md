## helper-tools

安装

```text
git clone https://github.com/alex-my/helper-tools.git
cd helper-tools
make mod
make i
```

## 命令列表

| 命令       | 缩写 | 作用                                                                 |
| ---------- | ---- | -------------------------------------------------------------------- |
| PWD        | pwd  | 显示当前目录                                                         |
| GitPullAll | gpa  | 遍历当前目录下的所有 git 项目，包括子文件夹(6 层)，分别执行 git pull |
| SVNUpAll   | sua  | 遍历当前目录下的所有 svn 项目，包括子文件夹(6 层)，分别执行 svn up   |

## 使用

假设`code`文件夹下有许多的`git`项目

```text
cd code

# 在 code 下执行以下命令，会在当前目录及子目录中寻找 git 项目
h GitPullAll

# 也可以使用缩写
h gpa
```

效果如图(里面有收集的各种语言的各种 git 项目)

![gitpullall](./images/gitpullall-process.png)
