package action

import (
	"io/ioutil"
	"path"

	"github.com/alex-my/ghelper/file"
)

// ListDirs 列出指定目录的文件夹，需要包含 contains，比如 .git
func ListDirs(dirname string, contains string, deep int) ([]string, error) {
	if deep <= 0 {
		return nil, nil
	}

	dirs, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(dirs))

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		newdirname := path.Join(dirname, dir.Name())

		if file.DirContains(newdirname, contains) {
			out = append(out, newdirname)
		} else {
			childout, err := ListDirs(newdirname, contains, deep-1)
			if err == nil && childout != nil {
				out = append(out, childout...)
			}
		}
	}

	return out, nil
}
