package sync

import (
	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func Go() error {
	dest := viper.GetString("dest")
	exclude := viper.GetStringSlice("exclude")
	src := viper.GetString("src")


	srcFiles, err := gatherFiles(src, dest, exclude)
	if err != nil {
		return err
	}

	if err := createFolders(srcFiles); err != nil {
		return err
	}

	if err := copyFiles(srcFiles); err != nil {
		return err
	}

	timber.Info("files copied:", len(srcFiles))

	return nil
}
