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
	var include string
	var logLevel int
	var output string
	var prune bool
	var src string

	flag.StringVar(&dest, "dest", "", "destination folder")
	flag.StringVar(&exclude, "exclude", "", "comma seperated list of regex patterns")
	flag.StringVar(&include, "include", "", "comma seperated list of regex patterns")
	flag.IntVar(&logLevel, "loglevel", timber.LevelWarning, "log level")
	flag.StringVar(&output, "o", "", "output filename")
	flag.BoolVar(&prune, "prune", false, "remove orphaned dest files/folders")
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

	if include != "" {
		patterns := []string{}

		parts := strings.Split(include, ",")
		patterns = append(patterns, parts...)

		viper.Set("include", patterns)
	}

	viper.Set("src", src)
	viper.Set("dest", dest)
	viper.Set("prune", prune)

	// loggers
	if output == "" {
		if err := timber.New(
			os.Stdout,
			logLevel,
			"",
			timber.FlagColorful,
		); err != nil {
			return err
		}
	} else {
		f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		if err := timber.New(
			f,
			logLevel,
			"[15:04:05]",
			0,
		); err != nil {
			return err
		}
	}

	return nil
}
