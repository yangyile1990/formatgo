package utils

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/erero"
)

func WriteBytes(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func IsRootExist(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && info != nil && info.IsDir() //这是简化版的就不要考虑其它错误啦
}

func IsFileExist(path string) bool {
	info, err := os.Stat(path)
	return !os.IsNotExist(err) && info != nil && !info.IsDir() //这是简化版的就不要考虑其它错误啦
}

func LsMapNamePath(root string) (map[string]string, error) {
	names, err := Ls(root)
	if err != nil {
		return nil, erero.WithMessage(err, "wrong")
	}
	var mp = make(map[string]string, len(names))
	for _, name := range names {
		mp[name] = filepath.Join(root, name)
	}
	return mp, nil
}

func Ls(root string) (names []string, err error) {
	infos, err := os.ReadDir(root)
	if err != nil {
		return nil, erero.WithMessage(err, "wrong")
	}
	names = make([]string, 0, len(infos))
	for _, info := range infos {
		names = append(names, info.Name())
	}
	return names, nil
}

func FilepathWalkOnFilesWithSuffixes(root string, suffixes []string, run func(path string, info os.FileInfo) error) (err error) {
	err = filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return erero.WithMessage(err, "wrong")
			}
			if info == nil {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if IsStringHasAnySuffix(path, suffixes) {
				return run(path, info)
			}
			return nil
		},
	)
	return err
}

func IsStringHasAnySuffix(s string, suffixes []string) bool {
	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}
