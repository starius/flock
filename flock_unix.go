// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package flock

import (
	"syscall"
)

func LockFd(fd uintptr) error {
	return syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB)
}

func UnlockFd(fd uintptr) error {
	return syscall.Flock(int(fd), syscall.LOCK_UN|syscall.LOCK_NB)
}
