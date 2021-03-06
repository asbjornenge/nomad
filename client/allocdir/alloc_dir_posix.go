// +build !windows

// Functions shared between linux/darwin.
package allocdir

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func (d *AllocDir) linkOrCopy(src, dst string, perm os.FileMode) error {
	// Attempt to hardlink.
	if err := os.Link(src, dst); err == nil {
		return nil
	}

	return fileCopy(src, dst, perm)
}

func (d *AllocDir) dropDirPermissions(path string) error {
	// Can't do anything if not root.
	if syscall.Geteuid() != 0 {
		return nil
	}

	u, err := user.Lookup("nobody")
	if err != nil {
		return err
	}

	uid, err := getUid(u)
	if err != nil {
		return err
	}

	gid, err := getGid(u)
	if err != nil {
		return err
	}

	if err := os.Chown(path, uid, gid); err != nil {
		return fmt.Errorf("Couldn't change owner/group of %v to (uid: %v, gid: %v): %v", path, uid, gid, err)
	}

	if err := os.Chmod(path, 0777); err != nil {
		return fmt.Errorf("Couldn't change owner/group of %v to (uid: %v, gid: %v): %v", path, uid, gid, err)
	}

	return nil
}

func getUid(u *user.User) (int, error) {
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return 0, fmt.Errorf("Unable to convert Uid to an int: %v", err)
	}

	return uid, nil
}

func getGid(u *user.User) (int, error) {
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return 0, fmt.Errorf("Unable to convert Gid to an int: %v", err)
	}

	return gid, nil
}
