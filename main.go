package clitesting

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

type Command struct {
	cmd       *exec.Cmd
	exitError error
	executed  bool
	stdout    string
	stderr    string
}

var UninitializedCommand = errors.New("You need to run this command first")

func NewCommand(name string, arg ...string) *Command {
	return &Command{
		cmd: exec.Command(name, arg...),
	}
}

func (c *Command) validate() {
	if !c.executed {
		log.Fatal(UninitializedCommand)
	}
}

func (c *Command) Run() {
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

func (c *Command) Error() error {
	c.validate()
	return c.exitError
}

func (c *Command) Stdout() string {
	c.validate()
	return c.stdout
}

func (c *Command) Stderr() string {
	c.validate()
	return c.stderr
}

// StdoutContains determines if command's STDOUT contains `str`, this operation
// is case insensitive.
func (c *Command) StdoutContains(str string) bool {
	c.validate()
	str = strings.ToLower(str)
	return strings.Contains(strings.ToLower(c.stdout), str)
}

// StdoutContains determines if command's STDERR contains `str`, this operation
// is case insensitive.
func (c *Command) StderrContains(str string) bool {
	c.validate()
	str = strings.ToLower(str)
	return strings.Contains(strings.ToLower(c.stderr), str)
}

func (c *Command) Success() bool {
	c.validate()
	return c.exitError == nil
}

func (c *Command) Failure() bool {
	c.validate()
	return c.exitError != nil
}
