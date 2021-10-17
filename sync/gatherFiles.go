package sync

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/designsbysm/timber/v2"
)

func gatherFiles(src string, dest string, exclude []string) (files []File, err error) {
	err = filepath.Walk(src, func(srcPath string, srcInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, pattern := range exclude {
			re := regexp.MustCompile(pattern)

			if re.Match([]byte(srcPath)) {
				timber.Debug("exclude:", srcPath)
				return nil
			}
		}

		if srcInfo.IsDir() {
			timber.Debug("skip:", srcPath)
			return nil
		}

		destPath := strings.Replace(srcPath, src, dest, 1)
		f := File{
			src:  srcPath,
			dest: destPath,
		}

		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}

			// dest doesn't exist, always copy
			timber.Debug("include:", srcPath)
			files = append(files, f)

			return nil
		}

		destInfo, err := os.Stat(destPath)
		if srcInfo.ModTime().After(destInfo.ModTime()) || srcInfo.Size() != destInfo.Size() {
			timber.Debug("include:", srcPath)
			files = append(files, f)
		}

		return nil
	})

	return
}
