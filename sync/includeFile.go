package sync

import (
	"os"
)

func includeFile(f *File) (bool, error) {
	srcInfo, err := os.Stat(f.src)
	if err != nil {
		return false, err
	}

	f.perm = srcInfo.Mode().Perm()

	destInfo, err := os.Stat(f.dest)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, err
		}

		// f.destPerm = destInfo.Mode().Perm()
		return true, nil
	}
	f.destExist = true

	if srcInfo.ModTime().After(destInfo.ModTime()) || srcInfo.Size() != destInfo.Size() {
		return true, nil
	}

	return false, nil
}
