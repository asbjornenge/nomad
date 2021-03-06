package allocdir

import (
	"errors"
	"os"
)

func (d *AllocDir) linkOrCopy(src, dst string, perm os.FileMode) error {
	return fileCopy(src, dst, perm)
}

// The windows version does nothing currently.
func (d *AllocDir) mountSharedDir(dir string) error {
	return errors.New("Mount on Windows not supported.")
}

// The windows version does nothing currently.
func (d *AllocDir) dropDirPermissions(path string) error {
	return nil
}

// The windows version does nothing currently.
func (d *AllocDir) unmountSharedDir(dir string) error {
	return nil
}
