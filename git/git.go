package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func ListFiles(dir string, timeout time.Duration) ([]string, error) {
	stdout, stderr, err := run(dir, "ls-files", timeout)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list files %v", strings.TrimSuffix(stderr.String(), "\n"))
	}
	files := make([]string, 0)
	lines := strings.Split(stdout.String(), "\n")
	files = append(files, lines[:len(lines)-1]...)
	return files, nil
}

func run(dir string, subcommand string, timeout time.Duration) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
	cmd := exec.Command("git", "--git-dir", dir, subcommand)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	if err != nil {
		return stdout, stderr, err
	}
	done := make(chan bool)
	go func() {
		err = cmd.Wait()
		done <- true
	}()
	select {
	case <-done:
		if err != nil {
			return stdout, stderr, err
		}
	case <-time.After(timeout):
		return stdout, stderr, fmt.Errorf("%v timed out after %v", strings.Join(cmd.Args, " "), timeout)
	}
	return stdout, stderr, nil
}
