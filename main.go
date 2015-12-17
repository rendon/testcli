package testcli

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

type Cmd struct {
	cmd       *exec.Cmd
	exitError error
	executed  bool
	stdout    string
	stderr    string
}

var UninitializedCmd = errors.New("You need to run this command first")

func Command(name string, arg ...string) *Cmd {
	return &Cmd{
		cmd: exec.Command(name, arg...),
	}
}

func (c *Cmd) validate() {
	if !c.executed {
		log.Fatal(UninitializedCmd)
	}
}

func (c *Cmd) Run() {
	var outBuf bytes.Buffer
	c.cmd.Stdout = &outBuf

	var errBuf bytes.Buffer
	c.cmd.Stderr = &errBuf

	if err := c.cmd.Run(); err != nil {
		c.exitError = err
	}
	c.stdout = string(outBuf.Bytes())
	c.stderr = string(errBuf.Bytes())
	c.executed = true
}

func (c *Cmd) Error() error {
	c.validate()
	return c.exitError
}

func (c *Cmd) Stdout() string {
	c.validate()
	return c.stdout
}

func (c *Cmd) Stderr() string {
	c.validate()
	return c.stderr
}

// StdoutContains determines if command's STDOUT contains `str`, this operation
// is case insensitive.
func (c *Cmd) StdoutContains(str string) bool {
	c.validate()
	str = strings.ToLower(str)
	return strings.Contains(strings.ToLower(c.stdout), str)
}

// StdoutContains determines if command's STDERR contains `str`, this operation
// is case insensitive.
func (c *Cmd) StderrContains(str string) bool {
	c.validate()
	str = strings.ToLower(str)
	return strings.Contains(strings.ToLower(c.stderr), str)
}

func (c *Cmd) Success() bool {
	c.validate()
	return c.exitError == nil
}

func (c *Cmd) Failure() bool {
	c.validate()
	return c.exitError != nil
}
