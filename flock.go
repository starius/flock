package flock

import (
	"os"
)

// LockFile places an exclusive lock on the file.
// If the file is already locked, exists with error.
func LockFile(f *os.File) error {
	return LockFd(f.Fd())
}

// UnlockFile removes an existing lock held by this process.
func UnlockFile(f *os.File) error {
	return UnlockFd(f.Fd())
}
