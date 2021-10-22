package sync

import "io/fs"

type File struct {
	src       string
	dest      string
	perm      fs.FileMode
	destExist bool
}
