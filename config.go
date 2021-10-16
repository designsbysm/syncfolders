package main

import (
	"errors"
	"flag"
	"os"

	"github.com/designsbysm/timber/v2"
	"github.com/designsbysm/timberfile"
	"github.com/spf13/viper"
)

func config() error {
	viper.SetConfigName("syncfolders")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("./syncfolders.yaml not found")
		} else {
			return err
		}
	}

	// loggers
	if err := timber.New(
		os.Stdout,
		viper.GetInt("timber.cli.level"),
		viper.GetString("timber.cli.timestamp"),
		viper.GetInt("timber.cli.flags"),
	); err != nil {
		return err
	}

	if err := timber.New(
		timberfile.New(viper.GetString("timber.file.path")),
		viper.GetInt("timber.file.level"),
		viper.GetString("timber.file.timestamp"),
		viper.GetInt("timber.file.flags"),
	); err != nil {
		return err
	}

	// flags
	var dest string
	var profile string
	var src string

	flag.StringVar(&dest, "dest", "", "destination folder")
	flag.StringVar(&src, "src", "", "source folder")
	flag.StringVar(&profile, "profile", "", "sync profile")
	flag.Parse()

	if src == "" {
		return errors.New("--src missing")
	} else if dest == "" {
		return errors.New("--dest missing")
		// } else if _, err := os.Stat(src); err != nil {
		// 	return err
	}

	viper.Set("src", src)
	viper.Set("dest", dest)
	viper.Set("profile", profile)

	return nil
}
