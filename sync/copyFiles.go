package sync

import (
	"io"
	"os"

	"github.com/designsbysm/syncfolders/progress"
	"github.com/designsbysm/timber/v2"
	"golang.org/x/sync/errgroup"
)

func copyFiles(files []File) error {
	eg := new(errgroup.Group)
	wg := make(chan struct{}, 10)
	progress.Set("Copied")

	for _, file := range files {
		file := file
		wg <- struct{}{}
		progress.Increment()
		timber.Info("copy:", file.src)

		eg.Go(func() error {
			var fin *os.File
			var fout *os.File
			var err error

			if fin, err = os.Open(file.src); err != nil {
				return closeCopy(wg, err)
			}
			defer fin.Close()

			if fout, err = os.Create(file.dest); err != nil {
				return closeCopy(wg, err)
			}
			defer fout.Close()

			if _, err = io.Copy(fout, fin); err != nil {
				return closeCopy(wg, err)
			}

			return closeCopy(wg, nil)
		})
	}

	err := eg.Wait()
	progress.Finish()

	return err
}

func closeCopy(wg chan struct{}, err error) error {
	<-wg
	return err
}
