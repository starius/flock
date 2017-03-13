package flock

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestLockFile(t *testing.T) {
	f, err := ioutil.TempFile("", "flock-test-")
	if err != nil {
		t.Fatalf("Unable to create temp file: %s.", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if err := LockFile(f); err != nil {
		t.Fatalf("Unable to lock file: %s.", err)
	}
	if err := UnlockFile(f); err != nil {
		t.Fatalf("Unable to unlock file: %s.", err)
	}
}

const lockFileCode = `package main

import (
	"fmt"
	"log"
	"os"

	"github.com/starius/flock"
)

func main() {
	f, err := os.Open("FILENAME")
	if err != nil {
		log.Fatalf("Unable to open file: %s.", err)
	}
	if err := flock.LockFile(f); err != nil {
		fmt.Printf("flock failed")
	}
}
`

func runGo(code string) ([]byte, error) {
	d, err := ioutil.TempDir("", "flock-go-")
	if err != nil {
		return nil, fmt.Errorf("ioutil.TempDir: %s", err)
	}
	defer os.Remove(d)
	f := d + "/test.go"
	ioutil.WriteFile(f, []byte(code), 0600)
	defer os.Remove(f)
	return exec.Command("go", "run", f).CombinedOutput()
}

func TestExclusiveness(t *testing.T) {
	f, err := ioutil.TempFile("", "flock-test-")
	if err != nil {
		t.Fatalf("Unable to create temp file: %s.", err)
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if err := LockFile(f); err != nil {
		t.Fatalf("Unable to lock file: %s.", err)
	}
	defer UnlockFile(f)
	code := strings.Replace(lockFileCode, "FILENAME", f.Name(), 1)
	result, err := runGo(code)
	if err != nil {
		t.Fatalf("Unable to run child process: %s. %s", err, result)
	}
	if string(result) != "flock failed" {
		t.Fatalf(
			"Expected to get %q from the child, got %q.",
			"flock failed",
			string(result),
		)
	}
}
