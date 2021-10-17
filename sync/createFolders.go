package sync

import (
	"os"
	"path/filepath"

	"github.com/designsbysm/mccoy"
	"github.com/designsbysm/timber/v2"
)

func createFolders(files []File) error {
	// create a list of folder from the found files
	folders := mccoy.Set{}
	for _, file := range files {
		path := filepath.Dir(file.dest)
		folders.Add(path)
	}

	for _, f := range folders.Items() {
		folder := mccoy.ItemString(f)
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return err
		}

		timber.Debug("create:", folder)
	}

	return nil
}
