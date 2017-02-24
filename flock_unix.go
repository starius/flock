// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package flock

import (
	"syscall"
)

func Lock(fd uintptr) error {
	return syscall.Flock(int(fd), syscall.LOCK_EX|syscall.LOCK_NB)
}

func Unlock(fd uintptr) error {
	return syscall.Flock(int(fd), syscall.LOCK_UN|syscall.LOCK_NB)
}
