package sync

import (
	"fmt"

	"github.com/spf13/viper"
)

func Go() error {
	src := viper.GetString("src")
	dest := viper.GetString("dest")

	name := viper.GetString("profile")
	key := fmt.Sprintf("profiles.%s", name)
	exclude := viper.GetStringSlice(fmt.Sprintf("%s.exclude", key))

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

	return nil
}
