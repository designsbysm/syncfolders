package sync

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/designsbysm/syncfolders/progress"
	"github.com/designsbysm/timber/v2"
)

func gatherFiles(src string, dest string, exclude []string, include []string) (files []File, err error) {
	progress.Set("Found")
	err = filepath.Walk(src, func(srcPath string, srcInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := strings.Replace(srcPath, src, dest, 1)
		f := File{
			src:  srcPath,
			dest: destPath,
		}

		for _, ePattern := range exclude {
			eRE := regexp.MustCompile(ePattern)

			if eRE.MatchString(f.src) {
				for _, iPattern := range include {
					iRE := regexp.MustCompile(iPattern)

					if iRE.MatchString(f.src) {
						add, err := includeFile(f)
						if err != nil {
							return err
						} else if add {
							timber.Debug("include:", f.src)
							files = append(files, f)
							progress.Increment()
							return nil
						}
					}
				}

				timber.Debug("exclude:", f.src)
				return nil
			}
		}

		if srcInfo.IsDir() || !srcInfo.Mode().IsRegular() {
			timber.Debug("skip:", f.src)
			return nil
		}

		add, err := includeFile(f)
		if err != nil {
			return err
		} else if add {
			timber.Debug("include:", f.src)
			files = append(files, f)
			progress.Increment()
			return nil
		}

		return nil
	})

	progress.Finish()
	return
}

func includeFile(f File) (bool, error) {
	destInfo, err := os.Stat(f.dest)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}

		return true, nil
	}

	srcInfo, err := os.Stat(f.src)
	if err != nil {
		return false, err
	} else if srcInfo.ModTime().After(destInfo.ModTime()) || srcInfo.Size() != destInfo.Size() {
		return true, nil
	}

	return false, nil
}
