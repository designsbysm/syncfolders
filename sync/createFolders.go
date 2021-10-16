package sync

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/designsbysm/mccoy"
)

func createFolders(files []File) error {
	// create a list of folder from the found files
	folders := mccoy.Set{}
	for _, file := range files {
		folders.Add(filepath.Dir(file.dest))
	}

	// remove folders with children, os.MkdirAll() will create an entire path
	for _, i := range folders.Items() {
		needle := mccoy.ItemString(i)

		for _, ii := range folders.Items() {
			test := mccoy.ItemString(ii)

			if strings.Contains(test, needle) && test != needle {
				folders.Delete(i)
				break
			}
		}
	}

	// make the folder structure
	for _, i := range folders.Items() {
		path := mccoy.ItemString(i)

		err := os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
