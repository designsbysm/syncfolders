package main

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/designsbysm/timber/v2"
	"github.com/spf13/viper"
)

func config() error {

	// flags
	var dest string
	var exclude string
	var logLevel int
	var src string

	flag.StringVar(&dest, "dest", "", "destination folder")
	flag.StringVar(&exclude, "exclude", "", "comma seperated list of regex patterns")
	flag.IntVar(&logLevel, "loglevel", timber.LevelInfo, "log level")
	flag.StringVar(&src, "src", "", "source folder")
	flag.Parse()

	if src == "" {
		return errors.New("--src missing")
	} else if dest == "" {
		return errors.New("--dest missing")
	}

	if exclude != "" {
		patterns := []string{}

		parts := strings.Split(exclude, ",")
		patterns = append(patterns, parts...)

		viper.Set("exclude", patterns)
	}

	viper.Set("dest", dest)
	viper.Set("src", src)

	// loggers
	if err := timber.New(
		os.Stdout,
		logLevel,
		"",
		timber.FlagColorful,
	); err != nil {
		return err
	}

	return nil
}
