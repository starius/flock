package flock

import (
	"os"
)

func LockFile(f *os.File) error {
	return LockFd(f.Fd())
}

func UnlockFile(f *os.File) error {
	return UnlockFd(f.Fd())
}
