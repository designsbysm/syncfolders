package sync

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/designsbysm/syncfolders/progress"
	"github.com/designsbysm/timber/v2"
)

func removeOrphans(src string, dest string) error {
	progress.Set("Orphaned")

	err := removeFiles(src, dest)
	if err != nil {
		progress.Finish()
		return err
	}

	for {
		orphans, err := removeFolders(dest)
		timber.Debug(orphans, err)
		if err != nil {
			return err
		} else if !orphans {
			break
		}
	}

	progress.Finish()
	return err
}

func removeFiles(src string, dest string) error {
	return filepath.Walk(dest, func(destPath string, destInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if destInfo.IsDir() {
			return nil
		}

		srcPath := strings.Replace(destPath, dest, src, 1)

		_, err = os.Stat(srcPath)
		if os.IsNotExist(err) {
			timber.Info("orphan:", destPath)
			progress.Increment()

			fi, err := os.Lstat(destPath)
			if err != nil {
				return err
			}

			if fi.Mode().IsRegular() {
				if err := os.Chmod(destPath, 0777); err != nil {
					return err
				}
			}

			return os.Remove(destPath)
		}

		return nil
	})
}

func removeFolders(dest string) (orphans bool, err error) {
	err = filepath.Walk(dest, func(destPath string, destInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		} else if !destInfo.IsDir() {
			return nil
		}

		err = os.Remove(destPath)
		if err == nil {
			timber.Debug("orphan:", destPath)
			progress.Increment()
			orphans = true
		}

		return nil
	})

	return
}
