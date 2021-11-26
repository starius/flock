//go:build darwin || dragonfly || freebsd || linux || nacl || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package flock

import (
	"golang.org/x/sys/unix"
)

// LockFile places an exclusive lock on the file.
// If the file is already locked, exists with error.
func LockFd(fd uintptr) error {
	return unix.Flock(int(fd), unix.LOCK_EX|unix.LOCK_NB)
}

// UnlockFile removes an existing lock held by this process.
func UnlockFd(fd uintptr) error {
	return unix.Flock(int(fd), unix.LOCK_UN|unix.LOCK_NB)
}
