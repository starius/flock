package flock

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLock(t *testing.T) {
	f, err := ioutil.TempFile("", "flock-test-")
	if err != nil {
		t.Fatalf("Unable to create temp file: %s.", err)
	}
	defer os.Remove(f.Name())
	if err := LockFile(f); err != nil {
		t.Fatalf("Unable to lock file: %s.", err)
	}
	if err := UnlockFile(f); err != nil {
		t.Fatalf("Unable to unlock file: %s.", err)
	}
}
