package sync

import (
	"io"
	"os"

	"github.com/designsbysm/timber/v2"
	"golang.org/x/sync/errgroup"
)

func copyFiles(files []File) error {
	wg := make(chan struct{}, 100)
	eg := new(errgroup.Group)

	for _, file := range files {
		file := file
		wg <- struct{}{}

		eg.Go(func() error {
			fin, err := os.Open(file.src)
			if err != nil {
				return closeCopy(wg, err)
			}
			defer fin.Close()

			fout, err := os.Create(file.dest)
			if err != nil {
				return closeCopy(wg, err)
			}
			defer fout.Close()

			_, err = io.Copy(fout, fin)
			if err != nil {
				return closeCopy(wg, err)
			}

			timber.Debug("copy:", file.src)

			return closeCopy(wg, nil)
		})
	}

	return eg.Wait()
}

func closeCopy(wg chan struct{}, err error) error {
	<-wg
	return err
}
