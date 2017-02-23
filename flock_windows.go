package flock

import (
	"os"
	"syscall"
	"unsafe"
)

var (
	modkernel32      = syscall.NewLazyDLL("kernel32.dll")
	procLockFileEx   = modkernel32.NewProc("LockFileEx")
	procUnlockFileEx = modkernel32.NewProc("UnlockFileEx")
)

const (
	LOCKFILE_FAIL_IMMEDIATELY = 1
	LOCKFILE_EXCLUSIVE_LOCK   = 2
	RESERVED                  = 0
	LOCKLOW                   = 0
	LOCKHIGH                  = 1
)

func Lock(f *os.File) error {
	var ol syscall.Overlapped
	r1, _, e1 := syscall.Syscall6(
		procLockFileEx.Addr(),
		6, // Number of arguments.
		uintptr(f.Fd()),
		uintptr(LOCKFILE_EXCLUSIVE_LOCK|LOCKFILE_FAIL_IMMEDIATELY),
		uintptr(RESERVED),
		uintptr(LOCKLOW),
		uintptr(LOCKHIGH),
		uintptr(unsafe.Pointer(&ol)),
	)
	if e1 != 0 {
		return e1
	} else if r1 == 0 {
		return syscall.EINVAL
	}
	return nil
}

func Unlock(f *os.File) error {
	var ol syscall.Overlapped
	r1, _, e1 := syscall.Syscall6(
		procUnlockFileEx.Addr(),
		5, // Number of arguments.
		uintptr(f.Fd()),
		uintptr(RESERVED),
		uintptr(LOCKLOW),
		uintptr(LOCKHIGH),
		uintptr(unsafe.Pointer(&ol)),
		0, // 6th argument is not used.
	)
	if e1 != 0 {
		return e1
	} else if r1 == 0 {
		return syscall.EINVAL
	}
	return nil
}
