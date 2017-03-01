package flock

import (
	"os"
)

func LockFile(f *os.File) error {
	return Lock(f.Fd())
}

func UnlockFile(f *os.File) error {
	return Unlock(f.Fd())
}
