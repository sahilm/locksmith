package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Repo struct {
	Dir     string
	Timeout time.Duration
}

func (r Repo) Init() error {
	_, stderr, err := r.run(exec.Command("git", "-C", r.Dir, "init"))
	if err != nil {
		return errors.Wrapf(err, "failed to init git repo %v", chomp(stderr))
	}
	return nil
}

func (r Repo) Add(files ...string) error {
	_, stderr, err := r.run(exec.Command("git", "-C", r.Dir, "add", strings.Join(files, " ")))
	if err != nil {
		return errors.Wrapf(err, "failed to add files %v", chomp(stderr))
	}
	return nil
}

func (r Repo) Commit(message string) error {
	_, stderr, err := r.run(exec.Command("git", "-C", r.Dir, "commit", "-m", message))
	if err != nil {
		return errors.Wrapf(err, "failed to commit %v", chomp(stderr))
	}
	return nil
}

func (r Repo) ListFiles() ([]string, error) {
	stdout, stderr, err := r.run(exec.Command("git", "-C", r.Dir, "ls-files"))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to list files %v", chomp(stderr))
	}
	files := make([]string, 0)
	lines := strings.Split(stdout.String(), "\n")
	files = append(files, lines[:len(lines)-1]...)
	return files, nil
}

func (r Repo) run(cmd *exec.Cmd) (stdout bytes.Buffer, stderr bytes.Buffer, err error) {
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
	case <-time.After(r.Timeout):
		return stdout, stderr, fmt.Errorf("%v timed out after %v", strings.Join(cmd.Args, " "), r.Timeout)
	}
	return stdout, stderr, nil
}

func chomp(buf bytes.Buffer) string {
	return strings.TrimSuffix(buf.String(), "\n")
}
