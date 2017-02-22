// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package flock

import (
	"os"
	"syscall"
)

func Lock(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
}

func Unlock(f *os.File) error {
	return syscall.Flock(int(f.Fd()), syscall.LOCK_UN|syscall.LOCK_NB)
}
