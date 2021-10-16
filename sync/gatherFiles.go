package sync

import (
	"os"
	"path/filepath"
	"strings"
)

func gatherFiles(src string, dest string, exclude []string) (files []File, err error) {
	// files := []File{}

	err = filepath.Walk(src, func(srcPath string, srcInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		//TODO: exclude paths

		if srcInfo.IsDir() {
			return nil
		}

		destPath := strings.Replace(srcPath, src, dest, 1)
		_, err = os.Stat(destPath)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
			f := File{
				src:  srcPath,
				dest: destPath,
			}

			// dest doesn't exist, always copy
			files = append(files, f)
			return nil
		}

		// TODO: compare src & dest
		// timber.Debug(destPath)

		// timber.Debug(path, info)
		// match := false
		// for _, white := range whitelist {
		// 	if match {
		// 		continue
		// 	}

		// 	whiteRE := regexp.MustCompile(white)
		// 	whiteLocation := whiteRE.FindIndex([]byte(path))

		// 	if len(whiteLocation) > 0 {
		// 		match = true

		// 		for _, black := range blacklist {
		// 			blackRE := regexp.MustCompile(black)
		// 			blackLocation := blackRE.FindIndex([]byte(path))

		// 			if len(blackLocation) > 0 {
		// 				match = false
		// 			}
		// 		}
		// 	}
		// }

		// if !match {
		// 	os.Remove(path)
		// 	_, err := os.Stat(path)
		// 	if err != nil {
		// 		complete = false
		// 	}
		// }

		return nil
	})

	return
}
