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
var pkgCmd = &Cmd{}

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

func Run(name string, arg ...string) {
	pkgCmd = Command(name, arg...)
	pkgCmd.Run()
}

func (c *Cmd) Error() error {
	c.validate()
	return c.exitError
}

func Error() error {
	return pkgCmd.Error()
}

func (c *Cmd) Stdout() string {
	c.validate()
	return c.stdout
}

func Stdout() string {
	return pkgCmd.Stdout()
}

func (c *Cmd) Stderr() string {
	c.validate()
	return c.stderr
}

func Stderr() string {
	return pkgCmd.Stderr()
}

// StdoutContains determines if command's STDOUT contains `str`, this operation
// is case insensitive.
func (c *Cmd) StdoutContains(str string) bool {
	c.validate()
	str = strings.ToLower(str)
	return strings.Contains(strings.ToLower(c.stdout), str)
}

func StdoutContains(str string) bool {
	return pkgCmd.StdoutContains(str)
}

// StdoutContains determines if command's STDERR contains `str`, this operation
// is case insensitive.
func (c *Cmd) StderrContains(str string) bool {
	c.validate()
	str = strings.ToLower(str)
	return strings.Contains(strings.ToLower(c.stderr), str)
}

func StderrContains(str string) bool {
	return pkgCmd.StderrContains(str)
}

func (c *Cmd) Success() bool {
	c.validate()
	return c.exitError == nil
}

func Success() bool {
	return pkgCmd.Success()
}

func (c *Cmd) Failure() bool {
	c.validate()
	return c.exitError != nil
}

func Failure() bool {
	return pkgCmd.Failure()
}
