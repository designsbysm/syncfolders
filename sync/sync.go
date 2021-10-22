package sync

import (
	"github.com/spf13/viper"
)

func Go() error {
	dest := viper.GetString("dest")
	exclude := viper.GetStringSlice("exclude")
	include := viper.GetStringSlice("include")
	prune := viper.GetBool("prune")
	src := viper.GetString("src")

	srcFiles, err := gatherFiles(src, dest, exclude, include)
	if err != nil {
		return err
	}

	if err := createFolders(srcFiles); err != nil {
		return err
	}

	if err := copyFiles(srcFiles); err != nil {
		return err
	}

	if prune {
		err := removeOrphans(src, dest)
		if err != nil {
			return err
		}
	}

	return nil
}
