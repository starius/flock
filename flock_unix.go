// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package flock

import (
	"syscall"
)

// LockFile places an exclusive lock on the file.
// If the file is already locked, exists with error.
func LockFd(fd uintptr) error {
	return syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB)
}

// UnlockFile removes an existing lock held by this process.
func UnlockFd(fd uintptr) error {
	return syscall.Flock(int(fd), syscall.LOCK_UN|syscall.LOCK_NB)
}
